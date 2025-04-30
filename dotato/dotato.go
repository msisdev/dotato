package dotato

import (
	"io/fs"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/state"
)

// Dotato is a kind of engine.
// It handles building blocks and exposes high level functions.
type Dotato struct {
	fs 		billy.Filesystem
	isMem	bool

	cdir	gp.GardenPath			// config directory
	cfg		*config.Config
	cig		*ignore.Ignore		// ignore file in config directory
	state	*state.State
}

// New Dotato instance with filesystem
func New() *Dotato {
	return &Dotato{
		fs: osfs.New("/"),
		isMem: false,
	}
}

func NewMemfs() *Dotato {
	return &Dotato{
		fs: memfs.New(),
		isMem: true,
	}
}

func (d *Dotato) setConfig() (err error) {
	if d.cfg != nil {
		return
	}

	// Config
	d.cfg, d.cdir, err = readConfig(d.fs)
	return
}

func (d *Dotato) setState() (err error) {
	if d.state != nil {
		return
	}

	d.state, err = readStateUnsafe(d.fs, d.isMem)
	return
}

func (d *Dotato) setConfigIgnore() (err error) {
	if d.cig != nil {
		return
	}

	// base is required
	err = d.setConfig()
	if err != nil {
		return
	}

	d.cig, err = readIgnore(d.fs, d.cdir)
	return
}

// func (d *Dotato) addIgnore(group string) (err error) {
// 	// ignore is required
// 	err = d.setConfigIgnore()
// 	if err != nil {
// 		return
// 	}

// 	// use group as directory name
// 	dir, err := gp.New(group)
// 	if err != nil {
// 		return
// 	}

// 	// append group to ignore
// 	err = d.cig.ReadRecur(dir)
// 	return
// }

func (d Dotato) GetGroups(plan string) (groups map[string]bool, err error) {
	err = d.setConfig()
	if err != nil {
		return
	}

	// Get groups
	groups = d.cfg.GetGroups(plan)
	return
}

func (d Dotato) GetGroupBase(group, resolver string) (base gp.GardenPath, notFound []string, err error) {
	err = d.setConfig()
	if err != nil {
		return
	}

	// Get base of group
	base, notFound, err = d.cfg.GetGroupBase(group, resolver)
	return
}

// Scan which files have to be imported
func (d Dotato) GetImportPaths(group string, base gp.GardenPath) (err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setState(); err != nil { return }
	if err = d.setConfigIgnore(); err != nil { return }
	// if err = d.addIgnore(group); err != nil { return }

	ig, err := readIgnoreRecur(d.fs, append(d.cdir, group))
	if err != nil {
		return
	}

	paths := make([]string, 0)
	filepath.Walk(base.Abs(), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// garden path
		var gpath gp.GardenPath
		gpath, err = gp.New(path)
		if err != nil {
			return err
		}

		// test config ignore
		if d.cig.IsIgnoredWithBaseDir(base, gpath) {
			return nil
		}

		// test group ignore
		if ig.IsIgnoredWithBaseDir(base, gpath) {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	println("Paths:", len(paths))

	return
}
