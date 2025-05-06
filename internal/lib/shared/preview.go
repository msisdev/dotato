package shared

import (
	"fmt"

	"github.com/msisdev/dotato/internal/ui/chspinner"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func (s Shared) PreviewDangerUnlink() ([]dotato.Preview, error) {
	hs, err := s.d.GetAllHistoryByMode(config.ModeLink)
	if err != nil {
		s.logger.Fatal(err)
		return nil, err
	}

	var (
		ps []dotato.Preview
		overwrite int
		title = "Scanning files ..."
	)
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
		for _, h := range hs {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			dot, err := gp.New(h.DotPath)
			if err != nil {
				return err
			}

			dtt, err := gp.New(h.DttPath)
			if err != nil {
				return err
			}

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
			up <- fmt.Sprintf(
				"overwrite %d, total %d",
				overwrite, len(ps),
			)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s Shared) PreviewImportGroupFile(
	group, resolver string,
) ([]dotato.Preview, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, err
	}

	var (
		ps []dotato.Preview
		create int
		overwrite int
		title = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
		return s.d.WalkImportFile(group, base, func(p dotato.Preview) error {
			// Check quit
			select {
			case <-quit:
				return ErrQuit
			default:
			}

			// Add preview
			ps = append(ps, p)
			switch p.DttOp {
			case dotato.FileOpNone:
				// do nothing
			case dotato.FileOpCreate:
				create++
			case dotato.FileOpOverwrite:
				overwrite++
			}

			// Update spinner
			up <- fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s Shared) PreviewImportGroupLink(
	group, resolver string,
) ([]dotato.Preview, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, err
	}

	var (
		ps []dotato.Preview
		dotOW int
		dttCR int
		dttOW int
		title = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
		return s.d.WalkImportLink(group, base, func(p dotato.Preview) error {
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
				// do nothing
			case dotato.FileOpCreate:
				return fmt.Errorf("dot file %s doesn't exist", p.Dot.Path.Abs())
			case dotato.FileOpOverwrite:
				dotOW++
			}

			switch p.DttOp {
			case dotato.FileOpNone:
				// do nothing
			case dotato.FileOpCreate:
				dttCR++
			case dotato.FileOpOverwrite:
				dttOW++
			}

			// Update spinner
			up <- fmt.Sprintf(
				"group %s: dot overwrite %d, dtt create %d, dtt overwrite %d, total %d",
				group, dotOW, dttCR, dttOW, len(ps),
			)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s Shared) PreviewExportGroupFile(
	group, resolver string,
) ([]dotato.Preview, error) {
	var (
		ps []dotato.Preview
		create int
		overwrite int
		title = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err := chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
		return s.d.WalkExportFile(group, func(p dotato.Preview) error {
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
			up <- fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s Shared) PreviewExportGroupLink(
	group, resolver string,
) ([]dotato.Preview, error) {
	var (
		ps []dotato.Preview
		create int
		overwrite int
		title = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err := chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
		return s.d.WalkExportLink(group, func(p dotato.Preview) error {
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
			up <- fmt.Sprintf(
				"group %s: create %d, overwrite %d, total %d",
				group, create, overwrite, len(ps),
			)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s Shared) PreviewUnlinkGroup(
	group, resolver string,
) ([]dotato.Preview, error) {
	base, err := s.GetGroupBase(group, resolver)
	if err != nil {
		s.logger.Fatal(err)
		return nil, err
	}

	var (
		ps []dotato.Preview
		overwrite int
		title = fmt.Sprintf("Scanning files of group %s...", group)
	)
	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
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
			up <- fmt.Sprintf(
				"group %s: overwrite %d, total %d",
				group, overwrite, len(ps),
			)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}
