package cli

import (
	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/arg"
	"github.com/msisdev/dotato/pkg/dotato"
)

type Cli struct {
	d	*dotato.Dotato
	l	*log.Logger
}

func NewCli(d *dotato.Dotato, logger *log.Logger) *Cli {
	return &Cli{
		d: d,
		l: logger,
	}
}

func (c *Cli) setLogLevel(level log.Level) {
	// Set level
	c.l.SetLevel(level)

	// Set options
	switch level {
	default: fallthrough
	case log.DebugLevel:
		c.l.SetReportCaller(true)

	case log.InfoLevel:
	case log.WarnLevel:
	case log.ErrorLevel:
	case log.FatalLevel:
	}
}

func (c *Cli) Run() {
	args, err := arg.Parse()
	if err != nil {
		c.l.Fatal(err)
		return
	}

	if args.Danger != nil {
		if args.Danger.Unlink != nil {
			defer c.DangerUnlink(args.Danger.Unlink)
			return
		}
	}
}
