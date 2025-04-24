package cli

import (
	"context"
	"os"

	"github.com/msisdev/dotato/pkg/config"
	"github.com/urfave/cli/v3"
)

var spec = &cli.Command{
	Usage: "A dotfile manager",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "verbose",
			Aliases: []string{"v"},
			Value: "false",
			Usage: "enable verbose output",
		},
	},
	Commands: []*cli.Command{
		dangerCli,
		planCli,
		groupCli,
		fileCli,
		{
			Name: "version",
			Aliases: []string{"v"},
			Usage: "show version",
			Action: cli.ActionFunc(func(ctx context.Context, cmd *cli.Command) error {
				println(config.GetDotatoVersion())
				return nil
			}),
		},
	},
}

func Run(ctx context.Context) error {
	return spec.Run(ctx, os.Args)
}
