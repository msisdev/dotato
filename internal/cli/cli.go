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
			return
		}
	}

	if args.Import != nil {
		if args.Import.Plan != nil {
			importPlan(logger, args.Import.Plan)
			return
		}
		if args.Import.Group != nil {
			importGroup(logger, args.Import.Group)
			return
		}
	}

	if args.Export != nil {
		if args.Export.Plan != nil {
			exportPlan(logger, args.Export.Plan)
			return
		}
		if args.Export.Group != nil {
			exportGroup(logger, args.Export.Group)
			return
		}
	}

	if args.Unlink != nil {
		if args.Unlink.Plan != nil {
			unlinkPlan(logger, args.Unlink.Plan)
			return
		}
		if args.Unlink.Group != nil {
			unlinkGroup(logger, args.Unlink.Group)
			return
		}
		return
	}

	if args.Version != nil {
		printVersion()
		return
	}
}

func printVersion() {
	panic("unimplemented")
}

func unlinkGroup(logger *log.Logger, unlinkGroupArgs *arg.UnlinkGroupArgs) {
	panic("unimplemented")
}

func unlinkPlan(logger *log.Logger, unlinkPlanArgs *arg.UnlinkPlanArgs) {
	panic("unimplemented")
}

func exportGroup(logger *log.Logger, exportGroupArgs *arg.ExportGroupArgs) {
	panic("unimplemented")
}

func exportPlan(logger *log.Logger, exportPlanArgs *arg.ExportPlanArgs) {
	panic("unimplemented")
}


