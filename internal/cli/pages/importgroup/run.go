package importgroup

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/internal/cli"
	"github.com/msisdev/dotato/internal/ui/spinner"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func spin(init spinner.State, f spinner.Task) error {
	m := spinner.New(init, f)

	p := tea.NewProgram(m)

	if model, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	} else {
		// Change type of the final model to spinner
		m = model.(spinner.Spinner)
	}

	return m.State.Result
}

func Run(logger *log.Logger, args cli.ImportGroupArgs) {
	var (
		fs 	= osfs.New("/")
		dtt =	dotato.NewWithFS(fs, false)
		base 	gp.GardenPath
		mode	string
	)
	{
		var err error

		// Load mode
		spin(
			spinner.State{Text: "Loading config mode..."},
			func(up chan<- spinner.State, _ <-chan bool) error {
				mode, err = dtt.GetConfigMode()
				if err != nil {
					up <- spinner.State{
						Text: "Error loading config mode",
					}
					return err
				}
				up <- spinner.State{
					Frame: "✔",
					Text: "Config mode: " + mode,
					End: true,
				}
				return nil
			},
		)

		// Get base
		
		spin(
			spinner.State{Text: "Loading group base..."},
			func(up chan<- spinner.State, quit <-chan bool) error {
				var notFound []string
				base, notFound, err = dtt.GetConfigGroupBase(args.Group, args.Resolver)

			},
		)
	}
}
