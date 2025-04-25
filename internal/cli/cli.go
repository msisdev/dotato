package cli

import (
	"github.com/msisdev/dotato/internal/arg"
	"github.com/msisdev/dotato/pkg/dotato"
	"github.com/msisdev/dotato/pkg/log"
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
	log.SetLevel(c.l, level)
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
