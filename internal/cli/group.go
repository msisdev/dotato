package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var groupCli = &cli.Command{
	Name: "group",
	Aliases: []string{"g"},
	Flags: []cli.Flag{},
	Commands: []*cli.Command{
		{
			Name: "in",
			Usage: "import files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("group in")
				return nil
			},
		},
		{
			Name: "out",
			Usage: "export files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("group out")
				return nil
			},
		},
		{
			Name: "tidy",
			Usage: "remove unexpected files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("group tidy")
				return nil
			},
		},
	},
}
