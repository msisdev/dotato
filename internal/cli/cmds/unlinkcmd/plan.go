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
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"gorm.io/gorm"
)

func UnlinkPlan(logger *log.Logger, args *args.UnlinkPlanArgs) {
	a := app.New(logger)

	{
		plans, err := a.E.GetConfigPlans()
		if err != nil {
			logger.Fatal(err)
			return
		}

		if groupList, ok := plans[args.Plan]; !ok {
			logger.Fatal("Group not found")
			return
		} else {
			if len(groupList) == 0 {
				
			}
		}
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
		list, err := previewspinner.RunPreviewUnlinkGroup(a, group, base)
		if err != nil {
			logger.Fatal(err)
			return
		}

		ps = append(ps, list...)
	}

	// Print preview
	count := previewprinter.RunPreviewUnlink(ps)
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
