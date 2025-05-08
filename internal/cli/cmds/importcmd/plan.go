package importcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/shared"
	"github.com/msisdev/dotato/internal/cli/ui/inputconfirm"
	"github.com/msisdev/dotato/internal/cli/ui/mxspinner"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/dotato"
	"github.com/msisdev/dotato/internal/lib/store"
)

func ImportPlan(logger *log.Logger, args *args.ImportPlanArgs) {
	s, err := shared.New(logger)
	if err != nil {
		logger.Fatal(err)
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
	if s.GetMode() == config.ModeFile {
		for group := range groups {
			temp, m, err := s.PreviewImportGroupFile(group, args.Resolver)
			if err != nil {
				logger.Fatal(err)
				return
			}
			ps = append(ps, temp...)
			mods += m
		}
	} else {
		for group := range groups {
			temp, m, err := s.PreviewImportGroupLink(group, args.Resolver)
			if err != nil {
				logger.Fatal(err)
				return
			}
			ps = append(ps, temp...)
			mods += m
		}
	}

	// Print preview list
	fmt.Print("\n🔎 Preview\n\n")
	for _, p := range ps {
		if s.GetMode() == config.ModeFile {
			fmt.Println(shared.SprintPreviewImportFile(p))
		} else {
			fmt.Println(shared.SprintPreviewImportLink(p))
		}
	}
	fmt.Println()

	if mods == 0 {
		fmt.Println("No files to import.")
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

	// Import
	var title string
	if s.GetMode() == config.ModeFile {
		title = "Importing files..."
	} else {
		title = "Importing links..."
	}
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		for _, pre := range ps {
			// Check quit
			select {
			case <-quit:
				return errQuit
			default:
			}

			// import
			if s.GetMode() == config.ModeFile {
				err := s.ImportFile(pre)
				if err != nil {
					return err
				}
			} else {
				err := s.ImportLink(pre)
				if err != nil {
					return err
				}
			}

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
