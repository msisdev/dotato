package exportcmd

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

func ExportPlan(logger *log.Logger, args *args.ExportPlanArgs) {
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
		ps []dotato.Preview
		mods int
	)
	if s.GetMode() == config.ModeFile {
		for group := range groups {
			temp, m, err := s.PreviewExportGroupFile(group, args.Resolver)
			if err != nil {
				logger.Fatal(err)
				return
			}
			ps = append(ps, temp...)
			mods += m
		}
	} else {
		for group := range groups {
			temp, m, err := s.PreviewExportGroupLink(group, args.Resolver)
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
			println(shared.SprintPreviewExportFile(p))
		} else {
			println(shared.SprintPreviewExportLink(p))
		}
	}
	fmt.Println()

	if mods == 0 {
		fmt.Println("No files to export.")
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
		title = "Exporting files..."
	} else {
		title = "Exporting links..."
	}
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
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
