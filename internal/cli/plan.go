package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var planCli = &cli.Command{
	Name: "plan",
	Aliases: []string{"p"},
	Flags: []cli.Flag{},
	Commands: []*cli.Command{
		{
			Name: "in",
			Usage: "import files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("plan in")
				return nil
			},
		},
		{
			Name: "out",
			Usage: "export files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("plan out")
				return nil
			},
		},
		{
			Name: "tidy",
			Usage: "remove unexpected files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("plan tidy")
				return nil
			},
		},
	},
}
