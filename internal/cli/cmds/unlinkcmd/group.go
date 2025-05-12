package unlinkcmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/ui"
	"github.com/msisdev/dotato/internal/cli/ui/basespinner"
	"github.com/msisdev/dotato/internal/cli/ui/confirm"
	"github.com/msisdev/dotato/internal/cli/ui/previewprinter"
	"github.com/msisdev/dotato/internal/cli/ui/previewspinner"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/lib/store"
	"gorm.io/gorm"
)

func UnlinkGroup(logger *log.Logger, args *args.UnlinkGroupArgs) {
	a := app.New(logger)

	// Get base
	base, err := basespinner.Run(a, args.Group, args.Resolver)
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Preview
	ps, err := previewspinner.RunPreviewUnlinkGroup(a, args.Group, base)
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Print preview
	count := previewprinter.RunPreviewUnlink(ps)
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
	title := "Unlinking ..."
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.E.StateTxSafe(func(tx *gorm.DB) error {
			for _, pre := range ps {
				// Check quit
				select {
				case <-quit:
					return ui.ErrQuit
				default:
				}

				// Unlink
				err := a.Unlink(pre, tx, os.FileMode(args.DirPerm), os.FileMode(args.FilePerm))
				if err != nil {
					return err
				}

				// Update spinner
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
