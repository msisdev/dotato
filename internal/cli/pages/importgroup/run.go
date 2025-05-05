package importgroup

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/lib/store"
	"github.com/msisdev/dotato/internal/ui/mxspinner"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func Run(logger *log.Logger, args *args.ImportGroupArgs) {
	var (
		fs 	= osfs.New("/")
		dtt =	dotato.NewWithFS(fs, false)
		base 	gp.GardenPath
		mode	string
	)
	{
		var err error

		// Get mode
		modeTask := func(store *store.Store[string], _ <-chan bool) error {
			// Get mode
			mode, err = dtt.GetConfigMode()
			if err != nil {
				store.Set("Error loading config mode")
				return err
			}

			store.Set("Config mode: " + mode)
			return nil
		}
		err = mxspinner.Run("Loading config mode...", modeTask)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Get base
		baseTask := func(store *store.Store[string], quit <-chan bool) error {
			// Get base
			var notFound []string
			base, notFound, err = dtt.GetConfigGroupBase(args.Group, args.Resolver)
			if err != nil {
				// Env vars not set
				if notFound != nil {
					store.Set("Env vars not set: " + args.Group)
					return nil
				}

				// general error
				store.Set("Error loading config group base")
				return err
			}

			store.Set("Config group base: " + base.Abs())

			return nil
		}
		err = mxspinner.Run("Loading config group base...", baseTask)
		if err != nil {
			logger.Fatal(err)
			return
		}
	}

	// Import
	switch mode {
	case config.ModeFile:
		importGroupFile(logger, args, fs, dtt, base)
	case config.ModeLink:
		importGroupLink(logger, args, fs, dtt, base)
	}
}
