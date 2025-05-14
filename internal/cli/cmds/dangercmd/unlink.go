package dangercmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/ui"
	"github.com/msisdev/dotato/internal/cli/ui/confirm"
	"github.com/msisdev/dotato/internal/cli/ui/modespinner"
	"github.com/msisdev/dotato/internal/cli/ui/previewprinter"
	"github.com/msisdev/dotato/internal/cli/ui/previewspinner"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/lib/store"
	"gorm.io/gorm"
)

func Unlink(logger *log.Logger, args *args.DangerUnlinkArgs) {
	a := app.New(logger)

	// Check mode
	mode, err := modespinner.Run(a)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	if mode == config.ModeFile {
		logger.Fatal("unlink group not supported in file mode")
		os.Exit(1)
	}

	// Get histories
	hs, err := a.E.GetHistoryByMode(mode)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	// Preview
	ps, err := previewspinner.RunPreviewDangerUnlink(a, hs)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	// Print preview
	update := previewprinter.RunPreviewUnlink(ps, args.ViewAll)
	if update == 0 {
		return
	}

	// Confirm
	if args.No {
		return
	}
	if !args.Yes {
		ok, err := confirm.Run("Do you want to proceed?")
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}
		if !ok {
			return
		}
	}

	// Execute
	title := "Unlinking all ..."
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.E.TxSafe(func(tx *gorm.DB) error {
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
		if err == ui.ErrQuit {
			return
		}
		logger.Fatal(err)
		os.Exit(1)
	}
}
