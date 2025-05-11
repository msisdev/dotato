package basespinner

import (
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/component/mxspinner"
	"github.com/msisdev/dotato/internal/lib/store"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Returns the base of a config group
func Run(a app.App, group, resolver string) (base gp.GardenPath, err error) {
	text := "Loading config group base..."
	err = mxspinner.Run(text, func(store *store.Store[string], quit <-chan bool) error {
		var notFound []string
		base, notFound, err = a.E.GetConfigGroupBase(group, resolver)
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

	return
}
