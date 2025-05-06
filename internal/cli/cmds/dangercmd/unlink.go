package dangercmd

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

func Unlink(logger *log.Logger, args *args.DangerUnlinkArgs) {
	s, err := shared.New(logger)
	if err != nil {
		logger.Fatal(err)
		return
	}
	if s.GetMode() == config.ModeFile {
		logger.Fatal("unlink group not supported in file mode")
		return	
	}

	// Preview
	ps, err := s.PreviewDangerUnlink()
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Print preview list
	fmt.Print("\n🔎 Preview\n\n")
	for _, p := range ps {
		var symbol string
		switch p.DttOp {
		case dotato.FileOpNone:
			symbol = "✔"
		case dotato.FileOpCreate:
			symbol = "?"
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

	// Unlink
	title := "Unlinking all ..."
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