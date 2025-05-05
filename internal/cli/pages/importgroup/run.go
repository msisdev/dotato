package importgroup

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/internal/cli"
	"github.com/msisdev/dotato/internal/ui/chanspinner"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func Run(logger *log.Logger, args cli.ImportGroupArgs) {
	var (
		fs 	= osfs.New("/")
		dtt =	dotato.NewWithFS(fs, false)
		base 	gp.GardenPath
		mode	string
	)
	{
		var err error

		// Get mode
		modeTask := func(up chan<- string, _ <-chan bool) error {
			// Get mode
			mode, err = dtt.GetConfigMode()
			if err != nil {
				up <- "Error loading config mode"
				return err
			}

			up <- "Config mode: " + mode
			return nil
		}
		err = chanspinner.Run("Loading config mode...", modeTask)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Get base
		baseTask := func(up chan<- string, quit <-chan bool) error {
			// Get base
			var notFound []string
			base, notFound, err = dtt.GetConfigGroupBase(args.Group, args.Resolver)
			if err != nil {
				// Env vars not set
				if notFound != nil {
					up <- "Env vars not set: " + args.Group
					return nil
				}

				// general error
				up <- "Error loading config group base"
				return err
			}

			return nil
		}
		err = chanspinner.Run("Loading config group base...", baseTask)
		if err != nil {
			logger.Fatal(err)
			return
		}
	}

	
}
