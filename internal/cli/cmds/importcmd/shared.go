package importcmd

// import (
// 	"fmt"

// 	"github.com/charmbracelet/log"
// 	"github.com/go-git/go-billy/v5"
// 	"github.com/go-git/go-billy/v5/osfs"
// 	"github.com/msisdev/dotato/internal/ui/chspinner"
// 	"github.com/msisdev/dotato/pkg/dotato"
// 	gp "github.com/msisdev/dotato/pkg/gardenpath"
// )

// type shared struct {
// 	logger	*log.Logger
// 	fs 			billy.Filesystem
// 	d 			*dotato.Dotato
// 	mode 		string
// }

// func newShared(logger *log.Logger) (*shared, error) {
// 	s := &shared{
// 		logger: logger,
// 		fs: osfs.New("/"),
// 	}
// 	s.d = dotato.NewWithFS(s.fs, false)

// 	// Get mode
// 	text := "Loading config mode..."
// 	err := chspinner.Run(text, func(up chan<- string, quit <-chan bool) error {
// 		var err error
// 		s.mode, err = s.d.GetConfigMode()
// 		if err != nil {
// 			up <- "Error loading config mode"
// 			return err
// 		}
// 		up <- "Config mode: " + s.mode
// 		return nil
// 	})
// 	if err != nil {
// 		logger.Fatal(err)
// 		return nil, err
// 	}

// 	return s, nil
// }

// func (s shared) getGroupBase(group, resolver string) (base gp.GardenPath, err error) {
// 	text := "Loading config group base..."
// 	err = chspinner.Run(text, func(up chan<- string, quit <-chan bool) error {
// 		var notFound []string
// 		base, notFound, err = s.d.GetConfigGroupBase(group, resolver)
// 		if err != nil {
// 			if notFound != nil {
// 				up <- "Env vars not set: " + group
// 				return nil
// 			}

// 			up <- "Error loading config group base"
// 			return err
// 		}

// 		up <- "Config group base: " + base.Abs()
// 		return nil
// 	})
// 	if err != nil {
// 		s.logger.Fatal(err)
// 		return
// 	}

// 	return
// }

// func (s shared) previewImportGroupFile(
// 	group, resolver string,
// ) ([]dotato.Preview, error) {
// 	base, err := s.getGroupBase(group, resolver)
// 	if err != nil {
// 		s.logger.Fatal(err)
// 		return nil, err
// 	}

// 	var (
// 		ps []dotato.Preview
// 		create int
// 		overwrite int
// 		title = fmt.Sprintf("Scanning files of group %s...", group)
// 	)
// 	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
// 		return s.d.WalkImportFile(group, base, func(p dotato.Preview) error {
// 			// Check quit
// 			select {
// 			case <-quit:
// 				return errQuit
// 			default:
// 			}

// 			// Add preview
// 			ps = append(ps, p)
// 			switch p.DttOp {
// 			case dotato.FileOpNone:
// 				// do nothing
// 			case dotato.FileOpCreate:
// 				create++
// 			case dotato.FileOpOverwrite:
// 				overwrite++
// 			}

// 			// Update spinner
// 			up <- fmt.Sprintf(
// 				"group %s: create %d, overwrite %d, total %d",
// 				group, create, overwrite, len(ps),
// 			)

// 			return nil
// 		})
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ps, nil
// }

// func (s shared) previewImportGroupLink(
// 	group, resolver string,
// ) ([]dotato.Preview, error) {
// 	base, err := s.getGroupBase(group, resolver)
// 	if err != nil {
// 		s.logger.Fatal(err)
// 		return nil, err
// 	}

// 	var (
// 		ps []dotato.Preview
// 		dotOW int
// 		dttCR int
// 		dttOW int
// 		title = fmt.Sprintf("Scanning files of group %s...", group)
// 	)
// 	err = chspinner.Run(title, func(up chan<- string, quit <-chan bool) error {
// 		return s.d.WalkImportLink(group, base, func(p dotato.Preview) error {
// 			// Check quit
// 			select {
// 			case <-quit:
// 				return errQuit
// 			default:
// 			}

// 			// Add preview
// 			ps = append(ps, p)

// 			// Count operations
// 			switch p.DotOp {
// 			case dotato.FileOpNone:
// 				// do nothing
// 			case dotato.FileOpCreate:
// 				return fmt.Errorf("dot file %s doesn't exist", p.Dot.Path.Abs())
// 			case dotato.FileOpOverwrite:
// 				dotOW++
// 			}

// 			switch p.DttOp {
// 			case dotato.FileOpNone:
// 				// do nothing
// 			case dotato.FileOpCreate:
// 				dttCR++
// 			case dotato.FileOpOverwrite:
// 				dttOW++
// 			}

// 			// Update spinner
// 			up <- fmt.Sprintf(
// 				"group %s: dot overwrite %d, dtt create %d, dtt overwrite %d, total %d",
// 				group, dotOW, dttCR, dttOW, len(ps),
// 			)

// 			return nil
// 		})
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ps, nil
// }

