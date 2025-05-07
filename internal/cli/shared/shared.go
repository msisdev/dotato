package shared

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/internal/cli/ui/chspinner"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrQuit = fmt.Errorf("quit")
)

type Shared struct {
	logger *log.Logger
	fs     billy.Filesystem
	d      *dotato.Dotato
	mode   string
}

func New(logger *log.Logger) (*Shared, error) {
	s := &Shared{
		logger: logger,
		fs:     osfs.New("/"),
	}
	s.d = dotato.NewWithFS(s.fs, false)

	// Get mode
	text := "Loading config mode..."
	err := chspinner.Run(text, func(up chan<- string, quit <-chan bool) error {
		var err error
		s.mode, err = s.d.GetConfigMode()
		if err != nil {
			up <- "Error loading config mode"
			return err
		}
		up <- "Config mode: " + s.mode
		return nil
	})
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return s, nil
}

func (s Shared) GetMode() string {
	return s.mode
}

func (s Shared) GetGroupBase(group, resolver string) (base gp.GardenPath, err error) {
	text := "Loading config group base..."
	err = chspinner.Run(text, func(up chan<- string, quit <-chan bool) error {
		var notFound []string
		base, notFound, err = s.d.GetConfigGroupBase(group, resolver)
		if err != nil {
			if notFound != nil {
				up <- "Env vars not set: " + group
				return nil
			}

			up <- "Error loading config group base"
			return err
		}

		up <- "Config group base: " + base.Abs()
		return nil
	})
	if err != nil {
		s.logger.Fatal(err)
		return
	}

	return
}

func (s Shared) GetGroups(plan string) (groups map[string]bool, err error) {
	return s.d.GetConfigGroups(plan)
}
