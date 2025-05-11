package modespinner

import (
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/lib/store"
)

func Run(a app.App) (mode string, err error) {
	text := "Loading config mode..."
	err = mxspinner.Run(text, func(store *store.Store[string], quit <-chan bool) error {
		mode, err = a.E.GetConfigMode()
		if err != nil {
			store.Set("Error loading config mode")
			return err
		}
		store.Set("Config mode: " + mode)
		return nil
	})

	return
}
