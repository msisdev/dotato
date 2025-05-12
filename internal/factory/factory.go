package factory

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/ignore"
	"github.com/msisdev/dotato/internal/state"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func ReadConfig(
	fs billy.Filesystem,
) (
	cfg *config.Config, cdir gp.GardenPath, err error,
) {
	// directory
	dir, err := gp.New(".")
	if err != nil {
		panic(err)
	}

	cfg, cdir, err = config.ReadRecur(fs, dir, DotatoFileNameConfig)
	return
}

// Returns true if config file was created
//
// Returns false if config file already exists
func WriteExampleConfig(fs billy.Filesystem, perm os.FileMode) (bool, error) {
	// Check if config file exists
	var path string
	{
		wd, err := os.Getwd()
		if err != nil {
			return false, err
		}

		path = filepath.Join(wd, DotatoFileNameConfig)

		_, err = fs.Stat(path)
		if err == nil && !os.IsNotExist(err) {
			return false, nil
		}
	}

	// Create config file
	f, err := fs.OpenFile(path, os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Write example config
	_, err = f.Write([]byte(config.GetExample()))
	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadIgnore(
	fs billy.Filesystem, dir gp.GardenPath,
) (
	ig *ignore.Ignore, err error,
) {
	// Init ignore
	ig = ignore.NewWithFS(fs, dir, DotatoFileNameIgnore)

	// Read ignore file in dir
	_, err = ig.Read(dir)
	if err != nil {
		return nil, err
	}

	return
}

func ReadIgnoreRecur(
	fs billy.Filesystem, dir gp.GardenPath,
) (
	ig *ignore.Ignore, err error,
) {
	// Init ignore
	ig = ignore.NewWithFS(fs, dir, DotatoFileNameIgnore)

	// Read ignore file in dir
	err = ig.ReadRecur(dir)
	if err != nil {
		return nil, err
	}

	return
}

func ReadState(fs billy.Filesystem, isMem bool) (*state.State, error) {
	if isMem {
		return state.New(fs, state.PathInMemory)
	}

	return state.New(fs, DotatoFilePathState)
}
