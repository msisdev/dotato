package engine

import (
	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/memfs"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/ignore"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

type Engine struct {
	fs      billy.Filesystem
	isMem   bool
	maxIter int

	cdir gp.GardenPath
	cfg  *config.Config
	ig   *ignore.Ignore
}

func New() *Engine {
	return NewWithFS(filesystem.NewOSFS(), false)
}

func NewMemfs() *Engine {
	return NewWithFS(memfs.New(), true)
}

func NewWithFS(fs billy.Filesystem, isMem bool) *Engine {
	return &Engine{
		fs:      fs,
		isMem:   isMem,
		maxIter: factory.DotatoMaxFSIter,
	}
}
