package dotato

import (
	"os"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
)

// Run `onSelect` selectively and recursively on files
//
//  - If selectIgnored = true, you are calling onSelect on ignored files.
//  - If selectIgnored = false, you are calling onSelect on non-ignored files.
func (d Dotato) Walk(
	root gp.GardenPath,
	ig *ignore.Ignore,
	selectIgnored bool,
	onSelect func(gp.GardenPath, os.FileInfo) error,
) error {
	if err := d.setIgnore(); err != nil { return err }

	iter := 0

	var dfs func(dir gp.GardenPath) (err error)
	dfs = func(dir gp.GardenPath) (err error) {
		// max iter exceeded ?
		iter++
		if iter > d.maxIter {
			return ErrMaxIterExceeded
		}

		// Get file infos
		fis, err := d.fs.ReadDir(dir.Abs())
		if err != nil {
			return err
		}

		// Iterate over file infos
		for _, fi := range fis {
			path := append(dir, fi.Name())

			isIgnored :=
				d.ig.IsIgnoredWithBaseDir(root, path) || 
				ig.IsIgnoredWithBaseDir(root, path)

			if selectIgnored != isIgnored {
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

func (d Dotato) WalkIgnored(
	root gp.GardenPath,
	ig *ignore.Ignore,
	onIgnored func(gp.GardenPath, os.FileInfo) error,
) error {
	return d.Walk(root, ig, true, onIgnored)
}

func (d Dotato) WalkNonIgnored(
	root gp.GardenPath,
	ig *ignore.Ignore,
	onNonIgnored func(gp.GardenPath, os.FileInfo) error,
) error {
	return d.Walk(root, ig, false, onNonIgnored)
}

// Scan dotfiles that is not ignored
func (d Dotato) WalkDotfile(
	group string,
	base gp.GardenPath,
	onDotfile func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setIgnore(); err != nil { return }

	// Get group ignore rules
	ig, err := d.GetGroupIgnore(group)
	if err != nil {
		return
	}

	return d.WalkNonIgnored(base, ig, onDotfile)
}

// Scan group dotato files that is not ignored
func (d Dotato) WalkDotato(
	group string,
	onDttfile func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setIgnore(); err != nil { return }

	// Get group ignore rules
	ig, err := d.GetGroupIgnore(group)
	if err != nil {
		return
	}

	return d.WalkNonIgnored(append(d.cdir, group), ig, onDttfile)
}

func (d Dotato) WalkAndPreviewImportFile(
	group string,
	base gp.GardenPath,
	onPreview func(PreviewImportFile) error,
) (err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setIgnore(); err != nil { return }

	onDotfile := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := d.DotToDtt(base, dot, group)

		// Get preview
		pre, err := d.PreviewImportFile(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return d.WalkDotfile(group, base, onDotfile)
}
