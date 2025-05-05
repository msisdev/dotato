package cli

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/state"
)

func importGroup(logger *log.Logger, args *ImportGroupArgs) {
	var (
		fs = osfs.New("/")
		dtt = dotato.NewWithFS(fs, false)
		base gp.GardenPath
		mode string
	)
	{
		var err error
		mode, err = dtt.GetConfigMode()
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Base
		var notFound []string
		base, notFound, err = dtt.GetConfigGroupBase(args.Group, args.Resolver)
		if err != nil {
			if notFound != nil {
				logger.Fatal("Group not found: " + args.Group)
				return
			}

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

func importGroupFile(
	logger *log.Logger,
	args *ImportGroupArgs,
	fs billy.Filesystem,
	dtt *dotato.Dotato,
	base gp.GardenPath,
) {
	// Preview
	var pres []dotato.PreviewImportFile
	err := dtt.WalkAndPreviewImportFile(args.Group, base, func(pre dotato.PreviewImportFile) error {
		pres = append(pres, pre)
		logger.Infof(`Preview: %s -> %s`, pre.Dot.Abs(), pre.Dtt.Abs())

		return nil
	})
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Confirm
	if !args.Yes {
		logger.Info("Preview complete. Use --yes to confirm import.")
		return
	}

	// Import
	for _, pre := range pres {
		if pre.Equal {
			continue
		}

		// Paths
		var (
			dotPath string
			dttPath = pre.Dtt.Abs()
		)
		if pre.DotReal != nil {
			dotPath = pre.DotReal.Abs()
		} else {
			dotPath = pre.Dot.Abs()
		}

		// Make dotato directory
		err = fs.MkdirAll(pre.Dtt.Parent().Abs(), 0755)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Copy file
		err = dtt.CopyFile(dotPath, dttPath)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Write history
		err = dtt.PutHistory(state.History{
			DotPath: dotPath,
			DttPath: dttPath,
			Mode:    config.ModeFile,
		})
		if err != nil {	
			logger.Fatal(err)
			return
		}

		logger.Info("Imported: " + dotPath)
	}

	logger.Info("Import complete.")
}

func importGroupLink(
	logger *log.Logger,
	args *ImportGroupArgs,
	fs billy.Filesystem,
	dtt *dotato.Dotato,
	base gp.GardenPath,
) {
	// Preview
	var pres []dotato.PreviewImportFile
	err := dtt.WalkAndPreviewImportFile(args.Group, base, func(pre dotato.PreviewImportFile) error {
		pres = append(pres, pre)
		logger.Infof(`Preview: %s -> %s`, pre.Dot.Abs(), pre.Dtt.Abs())

		return nil
	})
	if err != nil {
		logger.Fatal(err)
		return
	}

	// Confirm
	if !args.Yes {
		logger.Info("Preview complete. Use --yes to confirm import.")
		return
	}

	// Import
	for _, pre := range pres {
		if pre.Equal {
			continue
		}

		// Paths
		var (
			dotPath string
			dttPath = pre.Dtt.Abs()
		)
		if pre.DotReal != nil {
			dotPath = pre.DotReal.Abs()
		} else {
			dotPath = pre.Dot.Abs()
		}

		// Make dotato directory
		err = fs.MkdirAll(pre.Dtt.Parent().Abs(), 0755)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Copy file
		err = dtt.CopyFile(dotPath, dttPath)
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Write history
		err = dtt.PutHistory(state.History{
			DotPath: dotPath,
			DttPath: dttPath,
			Mode:    config.ModeFile,
		})
		if err != nil {	
			logger.Fatal(err)
			return
		}

		logger.Info("Imported: " + dotPath)
	}

	logger.Info("Import complete.")
}