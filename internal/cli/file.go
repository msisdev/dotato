package cli

import (
	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/arg"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/factory"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/state"
)

func fileMove(logger *log.Logger, args *arg.FileMoveArgs) {
	var (
		cfg *config.Config
		base gp.GardenPath
	)
	{
		var err error
		cfg, base, err = factory.ReadConfig()	
		if err != nil {
			logger.Fatal(err)
			return
		}

		// Find group of the argument paths
		
	}

	{

	}


	panic("unimplemented")
}
