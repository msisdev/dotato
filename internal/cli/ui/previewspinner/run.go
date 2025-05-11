package previewspinner

import (
	"fmt"

	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/ui"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/lib/store"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/state"
)

func RunPreviewDangerUnlink(a app.App, hs []state.History) ([]app.Preview, error) {
	var ps []app.Preview

	title := "Preview ..."
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

			store.TrySet(fmt.Sprintf("Preview %s ...", p.Dot.Path.Abs()))
		}
		store.Set("Preview done")

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewImportGroupFile(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := fmt.Sprintf("Preview %s ...", group)
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		err := a.WalkImportFile(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet(fmt.Sprintf("Preview %s %s ...", group, p.Dot.Path.Abs()))
			return nil
		})
		if err != nil {
			store.Set(fmt.Sprintf("Preview %s: %s", group, err.Error()))
		} else {
			store.Set(fmt.Sprintf("Preview %s done", group))
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewImportGroupLink(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := fmt.Sprintf("Preview %s ...", group)
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		err := a.WalkImportLink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet(fmt.Sprintf("Preview %s ...", p.Dot.Path.Abs()))

			return nil
		})
		if err != nil {
			store.Set(fmt.Sprintf("Preview %s: %s", group, err.Error()))
		} else {
			store.Set(fmt.Sprintf("Preview %s done", group))
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewExportGroupFile(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := fmt.Sprintf("Preview %s ...", group)
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		err := a.WalkExportFile(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet(fmt.Sprintf("Preview %s %s ...", group, p.Dot.Path.Abs()))

			return nil
		})
		if err != nil {
			store.Set(fmt.Sprintf("Preview %s: %s", group, err.Error()))
		} else {
			store.Set(fmt.Sprintf("Preview %s done", group))
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewExportGroupLink(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := fmt.Sprintf("Preview %s ...", group)
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		err := a.WalkExportLink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet(fmt.Sprintf("Preview %s %s ...", group, p.Dot.Path.Abs()))

			return nil
		})
		if err != nil {
			store.Set(fmt.Sprintf("Preview %s: %s", group, err.Error()))
		} else {
			store.Set(fmt.Sprintf("Preview %s done", group))
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func RunPreviewUnlinkGroup(a app.App, group string, base gp.GardenPath) ([]app.Preview, error) {
	var ps []app.Preview

	title := fmt.Sprintf("Preview %s ...", group)
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		err := a.WalkUnlink(group, base, func(p app.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ui.ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			store.TrySet(fmt.Sprintf("Preview %s %s ...", group, p.Dot.Path.Abs()))

			return nil
		})
		if err != nil {
			store.Set(fmt.Sprintf("Preview %s: %s", group, err.Error()))
		} else {
			store.Set(fmt.Sprintf("Preview %s done", group))
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}
