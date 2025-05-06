package importcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/lib/shared"
	"github.com/msisdev/dotato/internal/ui/chspinner"
	"github.com/msisdev/dotato/internal/ui/inputconfirm"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
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
	var ps []dotato.Preview
	if s.GetMode() == config.ModeFile {
		for group := range groups {
			temp, err := s.PreviewImportGroupFile(group, args.Resolver)
			if err != nil {
				logger.Fatal(err)
				return
			}
			ps = append(ps, temp...)
		}
	} else {
		for group := range groups {
			temp, err := s.PreviewImportGroupLink(group, args.Resolver)
			if err != nil {
				logger.Fatal(err)
				return
			}
			ps = append(ps, temp...)
		}
	}

	// Print preview list
	fmt.Print("\n🔎 Preview\n\n")
	for _, p := range ps {
		var symbol string
		switch p.DttOp {
		case dotato.FileOpNone:
			symbol = "✔"
		case dotato.FileOpCreate:
			symbol = "+"
		case dotato.FileOpOverwrite:
			symbol = "!"
		default:
			symbol = "?"
		}

		fmt.Printf("%s %s -> %s\n", symbol, p.Dot.Path.Abs(), p.Dtt.Path.Abs())
	}
	fmt.Println()

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
