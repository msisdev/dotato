package app

import (
	"os"

	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	"github.com/msisdev/dotato/pkg/engine"
	"gorm.io/gorm"
)

func (a App) ImportFile(
	pre Preview,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	if pre.DttOp == FileOpNone || pre.DttOp == FileOpSkip {
		return nil
	}

	var (
		dotabs = pre.Dot.Real.Abs() // use real path to get actual file
		dttabs = pre.Dtt.Path.Abs()
	)

	if pre.DttOp == FileOpOverwrite {
		if !pre.Dtt.IsFile {
			// Remove symlink
			err := a.fs.Remove(dotabs)
			if err != nil {
				return err
			}
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

	// Do not write history because dotfile is not changed.
	return nil
}

func (a App) ImportLink(
	pre Preview,
	tx *gorm.DB,
	dirPerm os.FileMode,
	filePerm os.FileMode,
) error {
	// If both dot and dtt are none or skip, do nothing
	if (pre.DotOp == FileOpNone || pre.DotOp == FileOpSkip) &&
		(pre.DttOp == FileOpNone || pre.DttOp == FileOpSkip) {
		return nil
	}

	var (
		dotabs = pre.Dot.Path.Abs()
		dttabs = pre.Dtt.Path.Abs()
	)

	// Handle dtt first.
	//
	// Do rename if it's possible.
	// If not, do create.
	if pre.DttOp != FileOpNone && pre.DttOp != FileOpSkip {
		if pre.DttOp == FileOpCreate {
			// Create directory
			err := a.fs.MkdirAll(pre.Dtt.Path.Parent().Abs(), dirPerm)
			if err != nil {
				return err
			}

			// Rename
			err = a.fs.Rename(dotabs, dttabs)
			if err != nil {
				return err
			}
		} else if pre.Dtt.IsFile && pre.Dot.IsFile {
			// Rename
			err := a.fs.Rename(dotabs, dttabs)
			if err != nil {
				return err
			}
		} else if pre.Dtt.IsFile && !pre.Dot.IsFile {
			// Copy
			err := filesystem.CreateAndCopyFile(a.fs, dotabs, dttabs, filePerm)
			if err != nil {
				return err
			}
		} else {
			// Remove symlink
			err := a.fs.Remove(dotabs)
			if err != nil {
				return err
			}

			// Create file
			err = filesystem.CreateAndCopyFile(a.fs, dotabs, dttabs, filePerm)
			if err != nil {
				return err
			}
		}
	}

	// Now create dot
	if pre.DotOp != FileOpNone {
		if pre.DotOp == FileOpOverwrite {
			// Remove dot (don't care if it is a file or symlink)
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
	err := a.E.TxUpsertHistory(tx, engine.History{
		DotPath: pre.Dot.Path.Abs(),
		DttPath: pre.Dtt.Real.Abs(),
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
		dttabs = pre.Dtt.Real.Abs() // use real path to get actual file
	)

	// Check dot stat
	if pre.DotOp == FileOpOverwrite {
		if !pre.Dot.IsFile {
			// Remove symlink
			err := a.fs.Remove(dotabs)
			if err != nil {
				return err
			}
		}
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
	err = a.E.TxUpsertHistory(tx, engine.History{
		DotPath: pre.Dot.Path.Abs(),
		DttPath: pre.Dtt.Real.Abs(),
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
		// Remove dot (don't care if it is a file or symlink)
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
	err = a.E.TxUpsertHistory(tx, engine.History{
		DotPath: pre.Dot.Path.Abs(),
		DttPath: pre.Dtt.Real.Abs(),
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
	if pre.DotOp != FileOpOverwrite {
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
	err = a.E.TxDeleteHistory(tx, engine.History{
		DotPath: pre.Dot.Path.Abs(),
	})
	if err != nil {
		return err
	}

	return nil
}
