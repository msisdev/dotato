package importgroup

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/lib/io"
	"github.com/msisdev/dotato/internal/lib/store"
	"github.com/msisdev/dotato/internal/ui/inputconfirm"
	"github.com/msisdev/dotato/internal/ui/mxspinner"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	errQuit = fmt.Errorf("quit")
)

func importGroupFile(
	logger *log.Logger,
	args *args.ImportGroupArgs,
	fs billy.Filesystem,
	dtt *dotato.Dotato,
	base gp.GardenPath,
) {
	// Preview
	var (
		pres []dotato.Preview
		createCount = 0
		overwriteCount = 0
	)
	{
		task := func(s *store.Store[string], quit <-chan bool) error {
			return dtt.WalkAndPreviewImportFile(args.Group, base, func(pre dotato.Preview) error {
				select {
				case <-quit:
					return errQuit
				default:
				}

				pres = append(pres, pre)
				switch pre.DttOp {
				case dotato.FileOpNone:
					// do nothing
				case dotato.FileOpCreate:
					createCount++
				case dotato.FileOpOverwrite:
					overwriteCount++
				}
					
				s.TrySet(fmt.Sprintf("create %d, overwrite %d, total %d", createCount, overwriteCount, len(pres)))
				return nil
			})
		}

		err := mxspinner.Run("Scanning files...", task)
		if err != nil {
			if err != errQuit {
				logger.Fatal(err)
			}
	
			return
		}

		if createCount == 0 && overwriteCount == 0 {
			logger.Info("No files to import")
			return
		}
	}

	// Print preview list
	fmt.Print("\n🔎 Preview\n\n")
	for _, pre := range pres {
		var symbol string
		switch pre.DttOp {
		case dotato.FileOpNone:
			symbol = "✔"
		case dotato.FileOpCreate:
			symbol = "+"
		case dotato.FileOpOverwrite:
			symbol = "!"
		default:
			symbol = "?"
		}

		fmt.Printf("%s %s -> %s\n", symbol, pre.Dot.Path.Abs(), pre.Dtt.Path.Abs())
	}
	fmt.Println()

	// Confirm
	yes, err := inputconfirm.Run("Do you want to proceed?")
	if err != nil {
		logger.Fatal(err)
		return
	}
	if !yes {
		return
	}

	// Import
	{
		task := func(s *store.Store[string], quit <-chan bool) error {
			for _, pre := range pres {
				// Check quit
				select {
				case <-quit:
					return errQuit
				default:
				}

				err := io.ImportFile(fs, dtt, pre)
				if err != nil {
					return err
				}

				s.TrySet(pre.Dot.Path.Abs())
			}

			s.TrySet("Done")

			return nil
		}

		err := mxspinner.Run("Importing files...", task)
		if err != nil {
			if err != errQuit {
				logger.Fatal(err)
			}
	
			return
		}
	}
}
