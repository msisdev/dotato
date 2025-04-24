package dotato

import (
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/log"
)

type Dotato struct {
	fs 		billy.Filesystem
	log		*log.Logger

	farm	*gardenpath.GardenPath

}
