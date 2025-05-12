package exportcmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/ui"
	"github.com/msisdev/dotato/internal/cli/ui/basespinner"
	"github.com/msisdev/dotato/internal/cli/ui/confirm"
	"github.com/msisdev/dotato/internal/cli/ui/modespinner"
	"github.com/msisdev/dotato/internal/cli/ui/previewprinter"
	"github.com/msisdev/dotato/internal/cli/ui/previewspinner"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/lib/store"
	"gorm.io/gorm"
)

func ExportGroup(logger *log.Logger, args *args.ExportGroupArgs) {
	a := app.New(logger)

	// Get mode
	mode, err := modespinner.Run(a)
	if err != nil {
		logger.Fatal(err)
		return
	}
	if mode != config.ModeFile && mode != config.ModeLink {
		logger.Fatal("Invalid mode")
		return
	}

	// Get base
	base, err := basespinner.Run(a, args.Group, args.Resolver)
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Preview
	var ps []app.Preview
	if mode == config.ModeFile {
		ps, err = previewspinner.RunPreviewExportGroupFile(a, args.Group, base)
		if err != nil {
			logger.Fatal(err)
			return
		}
	} else {
		ps, err = previewspinner.RunPreviewExportGroupLink(a, args.Group, base)
		if err != nil {
			logger.Fatal(err)
			return
		}
	}

	// Print preview
	var count int
	if mode == config.ModeFile {
		count = previewprinter.RunPreviewExportFile(ps)
	} else {
		count = previewprinter.RunPreviewExportLink(ps)
	}
	if count == 0 {
		return
	}

	// Confirm
	if !args.Yes {
		yes, err := confirm.Run("Do you want to proceed?")
		if err != nil {
			logger.Fatal(err)
			return
		}
		if !yes {
			return
		}
	}

	// Execute
	title := "Exporting..."
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.E.StateTxSafe(func(tx *gorm.DB) error {
			for _, pre := range ps {
				// Check quit
				select {
				case <-quit:
					return ui.ErrQuit
				default:
				}

				// Export
				var err error
				if mode == config.ModeFile {
					err = a.ExportFile(pre, tx, os.FileMode(args.DirPerm), os.FileMode(args.FilePerm))
				} else {
					err = a.ExportLink(pre, tx, os.FileMode(args.DirPerm), os.FileMode(args.FilePerm))
				}
				if err != nil {
					return err
				}

				store.TrySet(pre.Dot.Path.Abs())
			}

			store.Set("Done")
			return nil
		})
	})
	if err != nil {
		if err != ui.ErrQuit {
			logger.Fatal(err)
		}
		return
	}
}
