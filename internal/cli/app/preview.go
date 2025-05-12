package app

import (
	"fmt"

	"github.com/msisdev/dotato/internal/lib/filesystem"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

type FileOp int

const (
	FileOpNone FileOp = iota
	FileOpSkip
	FileOpCreate
	FileOpOverwrite
)

type Preview struct {
	Dot   *filesystem.PathStat
	DotOp FileOp
	Dtt   *filesystem.PathStat
	DttOp FileOp
}

func (a App) newPreview(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p := Preview{
		DotOp: FileOpNone,
		DttOp: FileOpNone,
	}

	var err error

	p.Dot, err = filesystem.NewPathStat(a.fs, dot)
	if err != nil {
		return nil, err
	}

	p.Dtt, err = filesystem.NewPathStat(a.fs, dtt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (a App) PreviewImportFile(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := a.newPreview(dot, dtt)
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
			equal, err := filesystem.IsFileContentEqual(
				a.fs,
				p.Dot.Path.Abs(),
				p.Dtt.Path.Abs(),
			)
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
func (a App) PreviewImportLink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := a.newPreview(dot, dtt)
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
			equal, err := filesystem.IsFileContentEqual(
				a.fs,
				p.Dot.Path.Abs(),
				p.Dtt.Path.Abs(),
			)
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
			equal, err := filesystem.IsFileContentEqual(
				a.fs,
				p.Dot.Path.Abs(),
				p.Dtt.Path.Abs(),
			)
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

func (a App) PreviewExportFile(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := a.newPreview(dot, dtt)
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
			equal, err := filesystem.IsFileContentEqual(
				a.fs,
				p.Dot.Path.Abs(),
				p.Dtt.Path.Abs(),
			)
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

func (a App) PreviewExportLink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := a.newPreview(dot, dtt)
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
			link, err := a.fs.Readlink(p.Dot.Path.Abs())
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

func (a App) PreviewUnlink(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*Preview, error) {
	p, err := a.newPreview(dot, dtt)
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
		target, err := a.fs.Readlink(p.Dot.Path.Abs())
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
