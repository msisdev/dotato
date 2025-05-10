package dotato

import (
	"fmt"
	"os"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

type FileOp int
const (
	FileOpNone FileOp = iota
	FileOpSkip
	FileOpCreate
	FileOpOverwrite
)

type PathStat struct {
	Path gp.GardenPath
	Target gp.GardenPath	// direct target of the symlink
	Real gp.GardenPath		// final target of the chained symlinks
	IsFile bool
	Exists bool
}

func (d Dotato) newPathStat(path gp.GardenPath) (*PathStat, error) {
	s := PathStat{
		Path: path,
		Target: path,
		Real: path,
	}

	abs := path.Abs()

	// Check if file exists
	info, err := d.fs.Lstat(abs)
	if err != nil {
		if os.IsNotExist(err) {
			s.Exists = false
			return &s, nil
		}
		return nil, err
	}
	s.Exists = true

	// Check if file is a symlink
	if info.Mode().Type()&os.ModeSymlink == 0 {
		// Not a symlink
		s.IsFile = true
		s.Real = make(gp.GardenPath, len(s.Path))
		copy(s.Real, s.Path)
	} else {
		// Symlink
		s.IsFile = false

		// Get target
		target, err := d.fs.Readlink(abs)
		if err != nil {
			return nil, err
		}
		s.Target, err = gp.New(target)
		if err != nil {
			return nil, err
		}

		// Get real
		realPath, err := d.evalSymlinks(abs)
		if err != nil {
			return nil, err
		}
		s.Real, err = gp.New(realPath)
		if err != nil {
			return nil, err
		}
	}
	
	return &s, nil
}

type Preview struct {
	Dot 	*PathStat
	DotOp	FileOp
	Dtt 	*PathStat
	DttOp FileOp
}

func (d Dotato) newPreview(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p := Preview{
		DotOp: FileOpNone,
		DttOp: FileOpNone,
	}

	var err error

	p.Dot, err = d.newPathStat(dot)
	if err != nil {
		return nil, err
	}

	p.Dtt, err = d.newPathStat(dtt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Preview Import File ////////////////////////////////////////////////////////

func (d Dotato) PreviewImportFile(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := d.newPreview(dot, dtt)
	if err != nil {
		return nil, err
	}

	// Dot file operation
	if !p.Dot.Exists {
		return nil, fmt.Errorf("dotfile %s does not exist", p.Dot.Path)
	} else if !p.Dot.IsFile {
		if p.Dot.Target.IsEqual(p.Dtt.Path) {
			// Dot file is a symlink to dtt
			p.DttOp = FileOpSkip
			return p, nil
		}
	}
	p.DotOp = FileOpNone

	// Dtt file operation
	if p.Dtt.Exists {
		if p.Dtt.IsFile {
			// Compare files
			equal, err := d.compareFile(p.Dot.Path.Abs(), p.Dtt.Path.Abs())
			if err != nil {
				return nil, err
			}

			if equal {
				// Files are equal
				p.DttOp = FileOpNone
			} else {
				// Files are not equal
				p.DttOp = FileOpOverwrite
			}
		} else {
			// Overwrite symlink
			p.DttOp = FileOpOverwrite
		}
	} else {
		// Create file
		p.DttOp = FileOpCreate
	}

	return p, nil
}

// Preview Import Link ////////////////////////////////////////////////////////

// Dot should be link and dtt should be file.
func (d Dotato) PreviewImportLink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := d.newPreview(dot, dtt)
	if err != nil {
		return nil, err
	}

	if !p.Dot.Exists {
		return nil, fmt.Errorf("dotfile %s does not exist", p.Dot.Path)
	}

	// Case 1: dotfile is a file
	if p.Dot.IsFile {
		// dot op
		p.DotOp = FileOpOverwrite

		// dtt op
		if !p.Dtt.Exists {
			p.DttOp = FileOpCreate
		} else if p.Dtt.IsFile {
			// Compare files
			equal, err := d.compareFile(p.Dot.Path.Abs(), p.Dtt.Path.Abs())
			if err != nil {
				return nil, err
			}
			if equal {
				p.DttOp = FileOpNone
			} else {
				p.DttOp = FileOpOverwrite
			}
		} else {
			p.DttOp = FileOpOverwrite
		}
		
		return p, nil
	}

	// Case 2: dotfile is a symlink

	// Dotfile operation
	if p.Dot.Real.IsEqual(p.Dtt.Path) {
		p.DotOp = FileOpNone
		p.DttOp = FileOpNone
		return p, nil
	}
	p.DotOp = FileOpOverwrite

	// Dotato operation
	if !p.Dtt.Exists {
		p.DttOp = FileOpCreate
	} else {
		if p.Dtt.IsFile {
		 	equal, err := d.compareFile(p.Dot.Path.Abs(), p.Dtt.Path.Abs())
			if err != nil {
				return nil, err
			}
			if equal {
				p.DttOp = FileOpNone
			} else {
				p.DttOp = FileOpOverwrite
			}
		} else {
			p.DttOp = FileOpOverwrite
		}
	}

	return p, nil
}

// Preview Export File ////////////////////////////////////////////////////////

func (d Dotato) PreviewExportFile(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := d.newPreview(dot, dtt)
	if err != nil {
		return nil, err
	}

	// Dtt file operation
	if !p.Dtt.Exists {
		return nil, fmt.Errorf("dotato file %s does not exist", p.Dtt.Path)
	}
	p.DttOp = FileOpNone

	// Dot file operation
	if p.Dot.Exists {
		if p.Dot.IsFile {
			// Compare files
			equal, err := d.compareFile(p.Dot.Path.Abs(), p.Dtt.Real.Abs())
			if err != nil {
				return nil, err
			}

			if equal {
				p.DotOp = FileOpNone
			} else {
				p.DotOp = FileOpOverwrite
			}
		} else {
			p.DotOp = FileOpOverwrite
		}
	} else {
		p.DotOp = FileOpCreate
	}

	return p, nil
}

// Preview Export Link ////////////////////////////////////////////////////////

func (d Dotato) PreviewExportLink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := d.newPreview(dot, dtt)
	if err != nil {
		return nil, err
	}

	// Dtt operation
	if !p.Dtt.Exists {
		return nil, fmt.Errorf("dotato file %s does not exist", p.Dtt.Path)
	}
	p.DttOp = FileOpNone

	// Dot operation
	if p.Dot.Exists {
		if p.Dot.IsFile {
			p.DotOp = FileOpOverwrite
		} else {
			link, err := d.fs.Readlink(p.Dot.Path.Abs())
			if err != nil {
				return nil, err
			}
			if link == p.Dtt.Path.Abs() {
				p.DotOp = FileOpNone
			} else {
				p.DotOp = FileOpOverwrite
			}
		}
	} else {
		p.DotOp = FileOpCreate
	}

	return p, nil
}

// Preview Unlink /////////////////////////////////////////////////////////////

func (d Dotato) PreviewUnlink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := d.newPreview(dot, dtt)
	if err != nil {
		return nil, err
	}

	// Dot file operation
	if !p.Dot.Exists {
		// Dot file does not exist.
		p.DotOp = FileOpNone
	} else if p.Dot.IsFile {
		// Dot file is a file.
		p.DotOp = FileOpNone
	} else {
		// Dot file is a symlink.
		target, err := d.fs.Readlink(p.Dot.Path.Abs())
		if err != nil {
			return nil, err
		}

		if target == p.Dtt.Path.Abs() {
			// Dot file is a symlink to dtt.
			p.DotOp = FileOpOverwrite
		} else {
			// Dot file is a symlink to another file.
			p.DotOp = FileOpNone
		}
	}

	// Dtt file operation
	if !p.Dtt.Exists {
		return nil, fmt.Errorf("dotato file %s does not exist", p.Dtt.Path)
	}
	p.DttOp = FileOpNone

	return p, nil
}
