package previewspinner

import (
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/component/mxspinner"
	"github.com/msisdev/dotato/internal/cli/ui"
	"github.com/msisdev/dotato/internal/lib/store"
	"github.com/msisdev/dotato/internal/state"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func RunPreviewDangerUnlink(a app.App, hs []state.History) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning histories ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		for _, h := range hs {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Make path from history
			dot, err := gp.New(h.DotPath)
			if err != nil {
				return err
			}
			dtt, err := gp.New(h.DttPath)
			if err != nil {
				return err
			}

			// Get preview
			p, err := a.PreviewUnlink(dot, dtt)
			if err != nil {
				return err
			}

			// Add preview
			ps = append(ps, *p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())
		}
		store.Set("Done")

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewImportGroupFile(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning files ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.WalkImportFile(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewImportGroupLink(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning links ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.WalkImportLink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewExportGroupFile(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning files ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.WalkExportFile(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewExportGroupLink(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning links ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.WalkExportLink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewUnlinkGroup(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Scanning links ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return a.WalkUnlink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet("Previewing " + p.Dot.Path.Abs())

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}
