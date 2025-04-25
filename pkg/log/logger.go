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

func New(level log.Level) *log.Logger {
	logger := log.New(os.Stderr)
	SetLevel(logger, level)

	return logger
}

func SetLevel(logger	*log.Logger, level log.Level) {
	// Set level
	logger.SetLevel(level)

	// Set options
	switch level {
	default: fallthrough
	case log.DebugLevel:
		logger.SetReportCaller(true)

	case log.InfoLevel:
	case log.WarnLevel:
	case log.ErrorLevel:
	case log.FatalLevel:
	}
}
