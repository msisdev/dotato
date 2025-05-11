package app

import (
	"os"

	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	"github.com/msisdev/dotato/pkg/state"
	"gorm.io/gorm"
)

func (a App) ImportFile(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if pre.DttOp == FileOpNone || pre.DttOp == FileOpSkip {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	if pre.DttOp == FileOpOverwrite {
		// Remove dtt
		err := a.fs.Remove(dttabs)
		if err != nil {
			return err
		}
	} else {
		// Make directory
		err := a.fs.MkdirAll(pre.Dtt.Path.Parent().Abs(), dirPerm)
		if err != nil {
			return err
		}
	}

	// Copy file
	err := filesystem.CreateAndCopyFile(a.fs, dotabs, dttabs, filePerm)
	if err != nil {
		return err
	}

	// Write history
	err = a.State.TxUpsertOne(tx, state.History{
		DotPath: dotabs,
		DttPath: dttabs,
		Mode:    config.ModeFile,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a App) ImportLink(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if (pre.DotOp == FileOpNone || pre.DotOp == FileOpSkip) &&
		(pre.DttOp == FileOpNone || pre.DttOp == FileOpSkip) {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	// Create dtt first
	if pre.DttOp != FileOpNone {
		if pre.DttOp == FileOpOverwrite {
			// Remove dtt
			err := a.fs.Remove(dttabs)
			if err != nil {
				return err
			}
		} else {
			// Create directory
			err := a.fs.MkdirAll(pre.Dtt.Path.Parent().Abs(), dirPerm)
			if err != nil {
				return err
			}
		}

		// Create file
		err := filesystem.CreateAndCopyFile(a.fs, dotabs, dttabs, filePerm)
		if err != nil {
			return err
		}
	}

	// Create dot after
	if pre.DotOp != FileOpNone {
		if pre.DotOp == FileOpOverwrite {
			// Remove dot
			err := a.fs.Remove(dotabs)
			if err != nil {
				return err
			}
		} else {
			// Create directory
			err := a.fs.MkdirAll(pre.Dot.Path.Parent().Abs(), dirPerm)
			if err != nil {
				return err
			}
		}

		// Create link
		err := a.fs.Symlink(dttabs, dotabs)
		if err != nil {
			return err
		}
	}

	// Write history
	err := a.State.TxUpsertOne(tx, state.History{
		DotPath: dotabs,
		DttPath: dttabs,
		Mode:    config.ModeLink,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a App) ExportFile(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if pre.DotOp == FileOpNone {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	if pre.DotOp == FileOpOverwrite {
		// Do nothing
	} else {
		// Make directory
		err := a.fs.MkdirAll(pre.Dot.Path.Parent().Abs(), dirPerm)
		if err != nil {
			return err
		}
	}

	// Copy file
	err := filesystem.CreateAndCopyFile(a.fs, dttabs, dotabs, filePerm)
	if err != nil {
		return err
	}

	// Write history
	err = a.State.TxUpsertOne(tx, state.History{
		DotPath: dotabs,
		DttPath: dttabs,
		Mode:    config.ModeFile,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a App) ExportLink(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if pre.DotOp == FileOpNone {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	// Handle dot
	if pre.DotOp == FileOpOverwrite {
		// Remove dot
		err := a.fs.Remove(dotabs)
		if err != nil {
			return err
		}
	} else {
		// Make directory
		err := a.fs.MkdirAll(pre.Dot.Path.Parent().Abs(), dirPerm)
		if err != nil {
			return err
		}
	}

	// Create link
	err := a.fs.Symlink(dttabs, dotabs)
	if err != nil {
		return err
	}

	// Write history
	err = a.State.TxUpsertOne(tx, state.History{
		DotPath: dotabs,
		DttPath: dttabs,
		Mode:    config.ModeLink,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a App) Unlink(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if pre.DotOp == FileOpNone ||
		pre.DotOp == FileOpCreate {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	// Remove dot
	err := a.fs.Remove(dotabs)
	if err != nil {
		return err
	}

	// Copy file
	err = filesystem.CreateAndCopyFile(a.fs, dttabs, dotabs, filePerm)
	if err != nil {
		return err
	}

	// Delete history
	err = a.State.TxDeleteOne(tx, state.History{
		DotPath: dotabs,
	})
	if err != nil {
		return err
	}

	return nil
}
