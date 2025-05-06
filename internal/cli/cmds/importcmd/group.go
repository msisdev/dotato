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

func ImportGroup(logger *log.Logger, args *args.ImportGroupArgs) {
	s, err := shared.New(logger)
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Preview
	var ps []dotato.Preview
	if s.GetMode() == config.ModeFile {
		ps, err = s.PreviewImportGroupFile(args.Group, args.Resolver)
		if err != nil {
			logger.Fatal(err)
			return
		}
	} else {
		ps, err = s.PreviewImportGroupLink(args.Group, args.Resolver)
		if err != nil {
			logger.Fatal(err)
			return
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
