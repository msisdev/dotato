package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var dangerCli = &cli.Command{
	Name: "danger",
	Aliases: []string{"!"},
	Flags: []cli.Flag{},
	Commands: []*cli.Command{
		{
			Name: "unlink",
			Usage: "revert all symlinks based on current state",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("danger unlink")
				return nil
			},
		},
	},
}
