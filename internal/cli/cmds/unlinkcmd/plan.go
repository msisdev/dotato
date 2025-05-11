package unlinkcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/component/inputconfirm"
	"github.com/msisdev/dotato/internal/cli/component/mxspinner"
	"github.com/msisdev/dotato/internal/cli/shared"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/dotato"
	"github.com/msisdev/dotato/internal/lib/store"
)

func UnlinkPlan(logger *log.Logger, args *args.UnlinkPlanArgs) {
	s, err := shared.New(logger)
	if err != nil {
		logger.Fatal(err)
		return
	}
	if s.GetMode() == config.ModeFile {
		logger.Fatal("unlink group not supported in file mode")
		return
	}

	// Get groups
	groups, err := s.GetGroups(args.Plan)
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Preview
	var (
		ps   []dotato.Preview
		mods int
	)
	for group := range groups {
		temp, m, err := s.PreviewUnlinkGroup(group, args.Resolver)
		if err != nil {
			logger.Fatal(err)
			return
		}
		ps = append(ps, temp...)
		mods += m
	}

	// Print preview list
	shared.PrintPreviewUnlink(ps)

	if mods == 0 {
		fmt.Println("No changes to be made.")
		return
	}

	// Confirm
	if !args.Yes {
		yes, err := inputconfirm.Run("Do you want to proceed?")
		if err != nil {
			logger.Fatal(err)
			return
		}
		if !yes {
			return
		}
	} else {
		fmt.Println("Proceeding...")
	}

	// Unlink
	title := fmt.Sprintf("Unlinking plan %s...", args.Plan)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		for _, pre := range ps {
			// Check quit
			select {
			case <-quit:
				return errQuit
			default:
			}

			// Unlink
			err := s.Unlink(pre)
			if err != nil {
				return err
			}

			// Update spinner
			store.TrySet(pre.Dot.Path.Abs())
		}

		store.Set("Done")

		return nil
	})
	if err != nil {
		if err != errQuit {
			logger.Fatal(err)
		}
		return
	}
}
