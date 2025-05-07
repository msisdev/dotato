package dotato

import (
	"os"

	"github.com/msisdev/dotato/internal/ignore"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Run `onSelect` selectively and recursively on files
//
//   - If selectIgnored = true, you are calling onSelect on ignored files.
//   - If selectIgnored = false, you are calling onSelect on non-ignored files.
func (d Dotato) Walk(
	root gp.GardenPath,
	ig *ignore.Ignore,
	selectIgnored bool,
	onSelect func(gp.GardenPath, os.FileInfo) error,
) error {
	if err := d.setIgnore(); err != nil {
		return err
	}

	iter := 0

	var dfs func(dir gp.GardenPath) (err error)
	dfs = func(dir gp.GardenPath) (err error) {
		iter++
		
		// max iter exceeded ?
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
			path := dir.Copy()
			path = append(path, fi.Name())

			isIgnored :=
				d.ig.IsIgnoredWithBaseDir(root, path) ||
					ig.IsIgnoredWithBaseDir(root, path)

			if selectIgnored != isIgnored {
				continue
			}
			// dotato files are always excluded
			if ok := dotatoFileNames[path.Last()]; ok {
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

// Walk dot or dtt ////////////////////////////////////////////////////////////

// Scan dotfiles that is not ignored
func (d Dotato) WalkDotfile(
	group string,
	base gp.GardenPath,
	onDot func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	// Get group ignore rules
	ig, err := d.GetGroupIgnore(group)
	if err != nil {
		return
	}

	return d.WalkNonIgnored(base, ig, onDot)
}

// Scan group dotato files that is not ignored
func (d Dotato) WalkDotato(
	group string,
	onDtt func(gp.GardenPath, os.FileInfo) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	// Get group ignore rules
	ig, err := d.GetGroupIgnore(group)
	if err != nil {
		return
	}

	base := append(d.cdir, group)
	return d.WalkNonIgnored(base, ig, onDtt)
}

// Walk and Preview ///////////////////////////////////////////////////////////

func (d Dotato) WalkImportFile(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := d.DotToDtt(base, dot, group)

		// Get preview
		pre, err := d.PreviewImportFile(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return d.WalkDotfile(group, base, onDot)
}

func (d Dotato) WalkImportLink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := d.DotToDtt(base, dot, group)

		// Get preview
		pre, err := d.PreviewImportLink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return d.WalkDotfile(group, base, onDot)
}

func (d Dotato) WalkExportFile(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	onDot := func(dtt gp.GardenPath, fi os.FileInfo) error {
		// Get dot path
		dot := d.DttToDot(dtt, base)

		// Get preview
		p, err := d.PreviewExportFile(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*p)
	}

	return d.WalkDotato(group, onDot)
}

func (d Dotato) WalkExportLink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	onDot := func(dtt gp.GardenPath, fi os.FileInfo) error {
		// Get dot path
		dot := d.DttToDot(dtt, base)

		// Get preview
		p, err := d.PreviewExportLink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*p)
	}

	return d.WalkDotato(group, onDot)
}

func (d Dotato) WalkUnlink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	if err = d.setConfig(); err != nil {
		return
	}
	if err = d.setIgnore(); err != nil {
		return
	}

	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := d.DotToDtt(base, dot, group)

		// Get preview
		pre, err := d.PreviewUnlink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return d.WalkDotfile(group, base, onDot)
}
