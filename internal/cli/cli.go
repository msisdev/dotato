package cli

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/arg"
)

func setLogLevel(logger *log.Logger, level log.Level) {
	// Set level
	logger.SetLevel(level)

	// Set options
	switch level {
	default:
		fallthrough
	case log.DebugLevel:
		logger.SetReportCaller(true)

	case log.InfoLevel:
	case log.WarnLevel:
	case log.ErrorLevel:
	case log.FatalLevel:
	}
}

func Run() {
	logger := log.New(os.Stderr)

	args, err := arg.Parse()
	if err != nil {
		logger.Fatal(err)
		return
	}

	if args.Danger != nil {
		if args.Danger.Unlink != nil {
			dangerUnlink(logger, args.Danger.Unlink)
		}
		return
	}

	if args.Plan != nil {
		if args.Plan.In != nil {
			planIn(logger, args.Plan.In)
		} else if args.Plan.Out != nil {
			planOut(logger, args.Plan.Out)
		} else if args.Plan.Tidy != nil {
			planTidy(logger, args.Plan.Tidy)
		}
		return
	}

	if args.Group != nil {
		if args.Group.In != nil {
			groupIn(logger, args.Group.In)
		} else if args.Group.Out != nil {
			groupOut(logger, args.Group.Out)
		} else if args.Group.Tidy != nil {
			groupTidy(logger, args.Group.Tidy)
		}
		return
	}

	if args.File != nil {
		if args.File.Move != nil {
			fileMove(logger, args.File.Move)
		}
		return
	}
}
