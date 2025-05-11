package app

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	"github.com/msisdev/dotato/pkg/engine"
	"github.com/msisdev/dotato/pkg/state"
)

type App struct {
	logger	*log.Logger
	fs     	billy.Filesystem
	E      	*engine.Engine
	State		*state.State
}

func New(logger *log.Logger) App {
	return NewWithFS(logger, filesystem.NewOSFS(), false)
}

func NewWithFS(logger *log.Logger, fs billy.Filesystem, isMem bool) App {
	state, err := factory.ReadState(fs, isMem)
	if err != nil {
		panic(err)
	}

	return App{
		logger: logger,
		fs:     fs,
		E:      engine.NewWithFS(fs, isMem),
		State: 	state,
	}
}
