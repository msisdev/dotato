package filesystem

import (
	"os"

	"github.com/go-git/go-billy/v6"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

type PathStat struct {
	Path   gp.GardenPath
	Target gp.GardenPath // direct target of the symlink
	Real   gp.GardenPath // final target of the chained symlinks
	IsFile bool
	Exists bool
}

func NewPathStat(fs billy.Filesystem, path gp.GardenPath) (*PathStat, error) {
	s := PathStat{
		Path:   path,
		Target: path,
		Real:   path,
	}

	abs := path.Abs()

	// Check if file exists
	info, err := fs.Lstat(abs)
	if err != nil {
		if os.IsNotExist(err) {
			s.Exists = false
			return &s, nil
		}
		return nil, err
	}
	s.Exists = true

	// Check if file is a symlink
	if info.Mode().Type()&os.ModeSymlink == 0 {
		// Not a symlink
		s.IsFile = true
		s.Real = make(gp.GardenPath, len(s.Path))
		copy(s.Real, s.Path)
	} else {
		// Symlink
		s.IsFile = false

		// Get target
		target, err := fs.Readlink(abs)
		if err != nil {
			return nil, err
		}
		s.Target, err = gp.New(target)
		if err != nil {
			return nil, err
		}

		// Get real
		realPath, err := EvalSymlinks(fs, abs)
		if err != nil {
			return nil, err
		}
		s.Real, err = gp.New(realPath)
		if err != nil {
			return nil, err
		}
	}

	return &s, nil
}
