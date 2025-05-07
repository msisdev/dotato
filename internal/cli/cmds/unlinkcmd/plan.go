package unlinkcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/shared"
	"github.com/msisdev/dotato/internal/cli/ui/chspinner"
	"github.com/msisdev/dotato/internal/cli/ui/inputconfirm"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
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
		ps []dotato.Preview
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
	fmt.Print("\n🔎 Preview\n\n")
	for _, p := range ps {
		fmt.Println(shared.SprintPreviewUnlink(p))
	}
	fmt.Println()

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
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
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
			up <- pre.Dot.Path.Abs()
		}

		up <- "Done"

		return nil
	})
	if err != nil {
		if err != errQuit {
			logger.Fatal(err)
		}
		return
	}
}
