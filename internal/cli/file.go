package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var fileCli = &cli.Command{
	Name: "file",
	Aliases: []string{"g"},
	Flags: []cli.Flag{},
	Commands: []*cli.Command{
		{
			Name: "in",
			Usage: "import files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("file in")
				return nil
			},
		},
		{
			Name: "out",
			Usage: "export files",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("file out")
				return nil
			},
		},
		{
			Name: "move",
			Usage: "modify file position",
			Aliases: []string{"mv"},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				println("file move")
				return nil
			},
		},
	},
}
