package unlinkcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/shared"
	"github.com/msisdev/dotato/internal/cli/ui/chspinner"
	"github.com/msisdev/dotato/internal/cli/ui/inputconfirm"
	"github.com/msisdev/dotato/pkg/config"
)

func UnlinkGroup(logger *log.Logger, args *args.UnlinkGroupArgs) {
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
	ps, mods, err := s.PreviewUnlinkGroup(args.Group, args.Resolver)
	if err != nil {
		logger.Fatal(err)
		return
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
	title := fmt.Sprintf("Unlinking group %s...", args.Group)
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
