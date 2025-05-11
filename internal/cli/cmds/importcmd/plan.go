package importcmd

import (
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
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"gorm.io/gorm"
)

func ImportPlan(logger *log.Logger, args *args.ImportPlanArgs) {
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

	// Get groups
	groups, ok, err := a.E.GetConfigGroups(args.Plan)
	if err != nil {
		logger.Fatal(err)
		return
	}
	if !ok {
		// Plan not found
		logger.Fatal("No such plan")
		return
	}
	if len(groups) == 0 {
		// Empty group list means all groups
		groups, err = a.E.GetConfigGroupAll()
		if err != nil {
			logger.Fatal(err)
			return
		}
	}

	// Get base
	bases := make(map[string]gp.GardenPath)
	for group := range groups {
		base, err := basespinner.Run(a, group, args.Resolver)
		if err != nil {
			logger.Fatal(err)
			return
		}

		bases[group] = base
	}

	// Preview
	var ps []app.Preview
	for group, base := range bases {
		var list []app.Preview
		
		if mode == config.ModeFile {
			list, err = previewspinner.RunPreviewImportGroupFile(a, group, base)
			if err != nil {
				logger.Fatal(err)
				return
			}
		} else {
			list, err = previewspinner.RunPreviewImportGroupLink(a, group, base)
			if err != nil {
				logger.Fatal(err)
				return
			}
		}

		ps = append(ps, list...)
	}

	// Print preview
	var count int
	if mode == config.ModeFile {
		count = previewprinter.RunPreviewImportFile(ps)
	} else {
		count = previewprinter.RunPreviewImportLink(ps)
	}
	if count == 0 {
		return
	}

	// Confirm
	if !args.Yes {
		ok, err := confirm.Run("Do you want to proceed?")
		if err != nil {
			logger.Fatal(err)
			return
		}
		if !ok {
			return
		}
	}

	// Execute
	title := "Importing..."
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.State.TxSafe(func(tx *gorm.DB) error {
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
					err = a.ImportFile(pre, tx)
				} else {
					err = a.ImportLink(pre, tx)
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
