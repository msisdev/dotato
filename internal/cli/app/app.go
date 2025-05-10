package app

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/engine"
	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/state"
)

type App struct {
	logger	*log.Logger
	fs     	billy.Filesystem
	E      	*engine.Engine
	state		*state.State
}

func NewWithFS(logger *log.Logger, fs billy.Filesystem, isMem bool) *App {
	state, err := factory.ReadState(fs, isMem)
	if err != nil {
		panic(err)
	}

	return &App{
		logger: logger,
		fs:     fs,
		E:      engine.NewWithFS(fs, isMem),
		state: 	state,
	}
}
