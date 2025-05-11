package cli

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/cmds/dangercmd"
	"github.com/msisdev/dotato/internal/cli/cmds/exportcmd"
	"github.com/msisdev/dotato/internal/cli/cmds/importcmd"
	"github.com/msisdev/dotato/internal/cli/cmds/initcmd"
	"github.com/msisdev/dotato/internal/cli/cmds/unlinkcmd"
	"github.com/msisdev/dotato/internal/cli/cmds/wherecmd"
)

func parse() (args.Args, error) {
	var args args.Args
	err := arg.Parse(&args)
	return args, err
}

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

	args, err := parse()
	if err != nil {
		logger.Fatal(err)
		return
	}

	if args.Danger != nil {
		if args.Danger.Unlink != nil {
			dangercmd.Unlink(logger, args.Danger.Unlink)
			return
		}
	}

	if args.Import != nil {
		if args.Import.Plan != nil {
			importcmd.ImportPlan(logger, args.Import.Plan)
			return
		}
		if args.Import.Group != nil {
			importcmd.ImportGroup(logger, args.Import.Group)
			return
		}
	}

	if args.Init != nil {
		initcmd.Init(logger, args.Init)
		return
	}

	if args.Export != nil {
		if args.Export.Plan != nil {
			exportcmd.ExportPlan(logger, args.Export.Plan)
			return
		}
		if args.Export.Group != nil {
			exportcmd.ExportGroup(logger, args.Export.Group)
			return
		}
	}

	if args.Unlink != nil {
		if args.Unlink.Plan != nil {
			unlinkcmd.UnlinkPlan(logger, args.Unlink.Plan)
			return
		}
		if args.Unlink.Group != nil {
			unlinkcmd.UnlinkGroup(logger, args.Unlink.Group)
			return
		}
		return
	}

	if args.Version != nil {
		printVersion(logger)
		return
	}

	if args.Where != nil {
		if args.Where.State != nil {
			wherecmd.WhereState(logger, args.Where.State)
			return
		}
		return
	}
}

func printVersion(logger *log.Logger) {
	logger.Info("Dotato version: " + dotatoVersion())
}
