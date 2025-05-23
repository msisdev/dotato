package importcmd

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

func ImportGroup(logger *log.Logger, args *args.ImportGroupArgs) {
	a := app.New(logger)

	// Get mode
	mode, err := modespinner.Run(a)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	if mode != config.ModeFile && mode != config.ModeLink {
		logger.Fatal("Invalid mode")
		os.Exit(1)
	}

	// Get base
	base, err := basespinner.Run(a, args.Group, args.Resolver)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	// Preview
	var ps []app.Preview
	if mode == config.ModeFile {
		ps, err = previewspinner.RunPreviewImportGroupFile(a, args.Group, base)
	} else {
		ps, err = previewspinner.RunPreviewImportGroupLink(a, args.Group, base)
	}
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	// Print preview
	var count int
	if mode == config.ModeFile {
		count = previewprinter.RunPreviewImportFile(ps, args.ViewAll)
	} else {
		count = previewprinter.RunPreviewImportLink(ps, args.ViewAll)
	}
	if count == 0 {
		return
	}

	// Confirm
	if args.No {
		return
	}
	if !args.Yes {
		yes, err := confirm.Run("Do you want to proceed?")
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}
		if !yes {
			return
		}
	}

	// Execute
	title := "Importing ..."
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.E.TxSafe(func(tx *gorm.DB) error {
			for _, pre := range ps {
				// Check quit
				select {
				case <-quit:
					return ui.ErrQuit
				default:
				}

				// Import
				var err error
				if mode == config.ModeFile {
					err = a.ImportFile(pre, os.FileMode(args.DirPerm), os.FileMode(args.FilePerm))
				} else {
					err = a.ImportLink(pre, tx, os.FileMode(args.DirPerm), os.FileMode(args.FilePerm))
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
		if err == ui.ErrQuit {
			return
		}

		logger.Fatal(err)
		os.Exit(1)
	}
}
