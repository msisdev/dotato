package log

import (
	"os"

	"github.com/charmbracelet/log"
)

// Port types
type Logger = log.Logger
type Level = log.Level
const (
	DebugLevel Level = log.DebugLevel
	InfoLevel  Level = log.InfoLevel
	WarnLevel  Level = log.WarnLevel
	ErrorLevel Level = log.ErrorLevel
	FatalLevel Level = log.FatalLevel
)

func NewLogger(level log.Level) *log.Logger {
	opt := log.Options{ Level: level }

	switch level {
	default: fallthrough
	case log.DebugLevel:
		opt.ReportCaller = true

	case log.InfoLevel:
	case log.WarnLevel:
	case log.ErrorLevel:
	case log.FatalLevel:
	}

	logger := log.NewWithOptions(os.Stderr, opt)

	return logger
}
