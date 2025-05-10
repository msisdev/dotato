package app

import (
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/engine"
)

type App struct {
	logger	*log.Logger
	fs			billy.Filesystem
	e 			*engine.Engine
}

func NewWithFS(logger *log.Logger, fs billy.Filesystem, isMem bool) *App {
	return &App{
		logger: logger,
		fs: fs,
		e: engine.NewWithFS(fs, isMem),
	}
}
