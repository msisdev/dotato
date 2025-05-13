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

	if !p.Dot.Exists {
		return nil, fmt.Errorf("dotfile %s does not exist", p.Dot.Path)
	}
	// dot exists
	if p.Dot.IsFile && !p.Dtt.Exists {
		p.DotOp = FileOpNone
		p.DttOp = FileOpCreate
		return p, nil
	}
	if p.Dot.IsFile && p.Dtt.Exists && p.Dtt.IsFile {
		p.DotOp = FileOpNone

		eq, err := filesystem.IsFileContentEqual(
			a.fs,
			p.Dot.Path.Abs(),
			p.Dtt.Path.Abs(),
		)
		if err != nil {
			return nil, err
		}
		if eq {
			p.DttOp = FileOpNone
		} else {
			p.DttOp = FileOpOverwrite
		}

		return p, nil
	}
	if p.Dot.IsFile && p.Dtt.Exists && !p.Dtt.IsFile {
		p.DotOp = FileOpNone
		p.DttOp = FileOpOverwrite
		return p, nil
	}
	// dot is symlink
	targEq := p.Dot.Target.IsEqual(p.Dtt.Path)
	if targEq {
		p.DotOp = FileOpNone
		p.DttOp = FileOpSkip
		return p, nil
	}
	realEq := p.Dot.Real.IsEqual(p.Dtt.Path)
	if realEq {
		p.DotOp = FileOpNone
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot is symlink to another file
	if !p.Dtt.Exists {
		p.DotOp = FileOpNone
		p.DttOp = FileOpCreate
		return p, nil
	}
	// dtt exists
	if p.Dtt.IsFile {
		p.DotOp = FileOpNone
		
		eq, err := filesystem.IsFileContentEqual(
			a.fs,
			p.Dot.Path.Abs(),
			p.Dtt.Path.Abs(),
		)
		if err != nil {
			return nil, err
		}
		if eq {
			p.DttOp = FileOpNone
		} else {
			p.DttOp = FileOpOverwrite
		}

		return p, nil
	}
	// dtt is symlink
	p.DotOp = FileOpNone
	p.DttOp = FileOpOverwrite
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
	// dot exists
	if p.Dot.IsFile && !p.Dtt.Exists {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpCreate
		return p, nil
	}
	if p.Dot.IsFile && p.Dtt.Exists && p.Dtt.IsFile {
		p.DotOp = FileOpOverwrite

		eq, err := filesystem.IsFileContentEqual(
			a.fs,
			p.Dot.Path.Abs(),
			p.Dtt.Path.Abs(),
		)
		if err != nil {
			return nil, err
		}
		if eq {
			p.DttOp = FileOpNone
		} else {
			p.DttOp = FileOpOverwrite
		}
		return p, nil
	}
	if p.Dot.IsFile && p.Dtt.Exists && !p.Dtt.IsFile {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpOverwrite
		return p, nil
	}
	// dot is symlink
	targEq := p.Dot.Target.IsEqual(p.Dtt.Path)
	realEq := p.Dot.Real.IsEqual(p.Dtt.Path)
	if targEq && realEq && p.Dot.Exists {
		p.DotOp = FileOpNone
		p.DttOp = FileOpNone
		return p, nil
	}
	if targEq && realEq && !p.Dot.Exists {
		p.DotOp = FileOpNone
		p.DttOp = FileOpSkip
		return p, nil
	}
	if targEq && !realEq {
		p.DotOp = FileOpNone
		p.DttOp = FileOpSkip
		return p, nil
	}
	if !targEq && realEq && p.Dot.Exists {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpNone
		return p, nil
	}
	if !targEq && realEq && !p.Dot.Exists {
		p.DotOp = FileOpSkip
		p.DttOp = FileOpSkip
		return p, nil
	}
	// dot is symlink to another file
	if p.Dtt.Exists && p.Dtt.IsFile {
		p.DotOp = FileOpOverwrite

		eq, err := filesystem.IsFileContentEqual(
			a.fs,
			p.Dot.Path.Abs(),
			p.Dtt.Path.Abs(),
		)
		if err != nil {
			return nil, err
		}
		if eq {
			p.DttOp = FileOpNone
		} else {
			p.DttOp = FileOpOverwrite
		}
		return p, nil
	}
	if p.Dtt.Exists && !p.Dtt.IsFile {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpOverwrite
		return p, nil
	}
	// dtt does not exist
	p.DotOp = FileOpOverwrite
	p.DttOp = FileOpCreate
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

	if !p.Dtt.Exists {
		return nil, fmt.Errorf("dotato file %s does not exist", p.Dtt.Path)
	}
	// dtt exists
	if !p.Dtt.IsFile && (
		p.Dtt.Target.IsEqual(p.Dot.Path) ||
		p.Dtt.Real.IsEqual(p.Dot.Path)) {
		p.DotOp = FileOpSkip
		p.DttOp = FileOpNone
		return p, nil
	}
	// dtt is file or symlink to another file
	if !p.Dot.Exists {
		p.DotOp = FileOpCreate
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot exists
	if !p.Dot.IsFile {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot is file
	eq, err := filesystem.IsFileContentEqual(
		a.fs,
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
	if err != nil {
		return nil, err
	}
	if eq {
		p.DotOp = FileOpNone
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot is different
	p.DotOp = FileOpOverwrite
	p.DttOp = FileOpNone
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
	if !p.Dtt.IsFile && (
		p.Dtt.Target.IsEqual(p.Dot.Path) ||
		p.Dtt.Real.IsEqual(p.Dot.Path)) {
		p.DotOp = FileOpSkip
		p.DttOp = FileOpNone
		return p, nil
	}
	// dtt is file
	if !p.Dot.Exists {
		p.DotOp = FileOpCreate
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot exists
	if p.Dot.IsFile {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot is symlink
	targEq := p.Dot.Target.IsEqual(p.Dtt.Path)
	if targEq {
		p.DotOp = FileOpNone
		p.DttOp = FileOpNone
		return p, nil
	}
	// dot is symlink to another file
	p.DotOp = FileOpOverwrite
	p.DttOp = FileOpNone
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

	// target very specific case
	if p.Dot.Exists &&
		!p.Dot.IsFile &&
		p.Dot.Target.IsEqual(p.Dtt.Path) &&
		p.Dtt.Exists &&
		p.Dtt.IsFile {
		p.DotOp = FileOpOverwrite
		p.DttOp = FileOpNone
		return p, nil
	}

	if !p.Dot.Exists {
		p.DotOp = FileOpSkip
	} else {
		p.DotOp = FileOpNone
	}
	if !p.Dtt.Exists {
		return nil, fmt.Errorf("dotato file %s does not exist", p.Dtt.Path)
	} else {
		p.DttOp = FileOpNone
	}
	return p, nil
}
