package shared

import (
	"fmt"

	"github.com/msisdev/dotato/internal/cli/ui/mxspinner"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/dotato"
	"github.com/msisdev/dotato/internal/lib/store"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func (s Shared) PreviewDangerUnlink() ([]dotato.Preview, int, error) {
	hs, err := s.d.GetAllHistoryByMode(config.ModeLink)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps        []dotato.Preview
		overwrite int
		title     = "Scanning files ..."
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		for _, h := range hs {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
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
			p, err := s.d.PreviewUnlink(dot, dtt)
			if err != nil {
				return err
			}

			// Add preview
			ps = append(ps, *p)

			// Count operations
			switch p.DotOp {
			case dotato.FileOpNone:
				// Do nothing
			case dotato.FileOpCreate:
				return fmt.Errorf("dot file %s doesn't exist", p.Dot.Path.Abs())
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"overwrite %d, total %d",
				overwrite, len(ps),
			))
		}

		return nil		
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, overwrite, nil
}

func (s Shared) PreviewImportGroupFile(
	group, resolver string,
) ([]dotato.Preview, int, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps        []dotato.Preview
		create    = 0
		overwrite = 0
		title     = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return s.d.WalkImportFile(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			// Count operations
			switch p.DttOp {
			case dotato.FileOpNone:
				// Do nothing
			case dotato.FileOpCreate:
				create++
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			))

			return nil
		})
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, create + overwrite, nil
}

func (s Shared) PreviewImportGroupLink(
	group, resolver string,
) ([]dotato.Preview, int, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps       []dotato.Preview
		dotOW    int
		dttCR    int
		dttOW    int
		totalMod int
		title    = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return s.d.WalkImportLink(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			isMod := false

			// Count operations
			switch p.DotOp {
			case dotato.FileOpNone:
				// do nothing
			case dotato.FileOpCreate:
				return fmt.Errorf("dot file %s doesn't exist", p.Dot.Path.Abs())
			case dotato.FileOpOverwrite:
				dotOW++
				isMod = true
			}

			switch p.DttOp {
			case dotato.FileOpNone:
				// do nothing
			case dotato.FileOpCreate:
				dttCR++
				isMod = true
			case dotato.FileOpOverwrite:
				dttOW++
				isMod = true
			}

			if isMod {
				totalMod++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"group %s: dot overwrite %d, dtt create %d, dtt overwrite %d, total %d",
				group, dotOW, dttCR, dttOW, len(ps),
			))

			return nil
		})
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, totalMod, nil
}

func (s Shared) PreviewExportGroupFile(
	group, resolver string,
) ([]dotato.Preview, int, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps        []dotato.Preview
		create    int
		overwrite int
		title     = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return s.d.WalkExportFile(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			// Count operations
			switch p.DotOp {
			case dotato.FileOpNone:
				// Do nothing
			case dotato.FileOpCreate:
				create++
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			))

			return nil
		})
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, create + overwrite, nil
}

func (s Shared) PreviewExportGroupLink(
	group, resolver string,
) ([]dotato.Preview, int, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps        []dotato.Preview
		create    int
		overwrite int
		title     = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return s.d.WalkExportLink(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			// Count operations
			switch p.DotOp {
			case dotato.FileOpNone:
				// Do nothing
			case dotato.FileOpCreate:
				create++
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			))

			return nil
		})
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, create + overwrite, nil
}

func (s Shared) PreviewUnlinkGroup(
	group, resolver string,
) ([]dotato.Preview, int, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, 0, err
	}

	var (
		ps        []dotato.Preview
		overwrite int
		title     = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		return s.d.WalkUnlink(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)

			// Count operations
			switch p.DotOp {
			case dotato.FileOpNone:
				// Do nothing
			case dotato.FileOpCreate:
				return fmt.Errorf("dot file %s doesn't exist", p.Dot.Path.Abs())
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			store.Set(fmt.Sprintf(
				"group %s: overwrite %d, total %d",
				group, overwrite, len(ps),
			))

			return nil
		})
	})
	if err != nil {
		return nil, 0, err
	}

	return ps, overwrite, nil
}
