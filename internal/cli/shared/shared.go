package shared

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/cli/ui/component/mxspinner"
	"github.com/msisdev/dotato/internal/dotato"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	"github.com/msisdev/dotato/internal/lib/store"
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
		fs:     filesystem.NewOSFS(),
	}
	s.d = dotato.NewWithFS(s.fs, false)

	// Get mode
	text := "Loading config mode..."
	err := mxspinner.Run(text, func(store *store.Store[string], quit <-chan bool) error {
		var err error
		s.mode, err = s.d.GetConfigMode()
		if err != nil {
			store.Set("Error loading config mode")
			return err
		}
		store.Set("Config mode: " + s.mode)
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
	err = mxspinner.Run(text, func(store *store.Store[string], quit <-chan bool) error {
		var notFound []string
		base, notFound, err = s.d.GetConfigGroupBase(group, resolver)
		if err != nil {
			if notFound != nil {
				store.Set("Env vars not set: " + group)
				return nil
			}

			store.Set("Error loading config group base")
			return err
		}

		store.Set("Config group base: " + base.Abs())
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
