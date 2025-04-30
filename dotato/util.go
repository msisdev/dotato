package dotato

import (
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
)

func (d Dotato) walk(root gp.GardenPath, ig *ignore.Ignore) (es []Entity, err error) {
	iter := 0

	var dfs func(dir gp.GardenPath) (err error)
	dfs = func(dir gp.GardenPath) (err error) {
		iter++
		if iter > d.maxIter {
			return ErrMaxIterExceeded
		}

		// Get entries
		fis, err := d.fs.ReadDir(dir.Abs())
		if err != nil {
			return err
		}

		// Iterate over entries
		for _, fi := range fis {
			path := append(dir, fi.Name())
			if d.ig.IsIgnoredWithBaseDir(root, path) {
				continue
			}
			if ig.IsIgnoredWithBaseDir(root, path) {
				continue
			}

			// Handle file
			if !fi.IsDir() {
				es = append(es, Entity{path, fi})
				continue
			}

			// Handle directory
			if fi.IsDir() {
				err = dfs(path)
				if err != nil {
					return err
				}
			}
		}

		return
	}

	err = dfs(root)
	if err != nil {
		return nil, err
	}

	return
}
