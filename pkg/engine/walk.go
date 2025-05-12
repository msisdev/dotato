package engine

import (
	"fmt"
	"os"

	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/ignore"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrMaxIterExceeded = fmt.Errorf("max iteration exceeded. please check your ignore file")
)

// Run `onSelect` selectively and recursively on files
//
//   - If selectIgnored = true, call onSelect on ignored files.
//   - If selectIgnored = false, call onSelect on non-ignored files.
func (e Engine) Walk(
	root gp.GardenPath,
	ig *ignore.Ignore,
	selectIgnored bool,
	onSelect func(gp.GardenPath, os.FileInfo) error,
) error {
	if err := e.readIgnore(); err != nil {
		return err
	}

	iter := 0

	var dfs func(dir gp.GardenPath) (err error)
	dfs = func(dir gp.GardenPath) (err error) {
		iter++

		// max iter exceeded ?
		if iter > e.maxIter {
			return ErrMaxIterExceeded
		}

		// Get file infos
		fis, err := e.fs.ReadDir(dir.Abs())
		if err != nil {
			return err
		}

		// Iterate over file infos
		for _, fi := range fis {
			path := dir.Copy()
			path = append(path, fi.Name())

			isIgnored :=
				e.ig.IsIgnoredWithBaseDir(root, path) ||
					ig.IsIgnoredWithBaseDir(root, path)

			if selectIgnored != isIgnored {
				continue
			}
			// dotato files are always excluded
			if ok := factory.DotatoFileNames[path.Last()]; ok {
				continue
			}

			if fi.IsDir() {
				// Directory: recurse
				err = dfs(path)
				if err != nil {
					return err
				}
			} else {
				// File: call function
				err = onSelect(path, fi)
				if err != nil {
					return err
				}
			}
		}

		return
	}

	return dfs(root)
}

func (e Engine) WalkIgnored(
	root gp.GardenPath,
	ig *ignore.Ignore,
	onIgnored func(gp.GardenPath, os.FileInfo) error,
) error {
	return e.Walk(root, ig, true, onIgnored)
}

func (e Engine) WalkNonIgnored(
	root gp.GardenPath,
	ig *ignore.Ignore,
	onNonIgnored func(gp.GardenPath, os.FileInfo) error,
) error {
	return e.Walk(root, ig, false, onNonIgnored)
}

// Scan dotfiles that is not ignored
func (e *Engine) WalkDotDir(
	group string,
	base gp.GardenPath,
	onDot func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = e.readConfig(); err != nil {
		return
	}
	if err = e.readIgnore(); err != nil {
		return
	}

	// Get group ignore rules
	ig, err := e.ReadGroupIgnore(group)
	if err != nil {
		return
	}

	return e.WalkNonIgnored(base, ig, onDot)
}

// Scan dotato files in group that is not ignored
func (e *Engine) WalkDttDir(
	group string,
	onDtt func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = e.readConfig(); err != nil {
		return
	}
	if err = e.readIgnore(); err != nil {
		return
	}

	// Get group ignore rules
	ig, err := e.ReadGroupIgnore(group)
	if err != nil {
		return
	}

	// Get group dir path
	base := e.cdir.Copy()
	base = append(base, group)

	return e.WalkNonIgnored(base, ig, onDtt)
}
