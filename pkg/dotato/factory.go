package dotato

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/state"
)

const (
	// Location of user-wide config directory
	DotatoDirPathEnv 					= "DOTATO_DIRECTORY"
	DotatoDirName							= "dotato"

	// State file name
	StateFileNameEnv 					= "DOTATO_STATE"
	StateFileNameDefault 			= "dotatostate.sqlite"

	// Config file name
	ConfigFileNameEnv 				= "DOTATO_CONFIG"
	ConfigFileNameDefault 		= "dotato.yaml"
	
	// Ignore file name
	IgnoreFileNameEnv 				= "DOTATO_IGNORE"
	IgnoreFileNameDefault 		= ".dotatoignore"

	MaxFileSystemIterEnv 			= "DOTATO_MAX_FS_ITER"
	MaxFileSystemIterDefault	= 10000
)

var (
	ErrConfigNotFound = fmt.Errorf("config file not found")
	ErrMaxIterExceeded = fmt.Errorf("max iteration exceeded")
)

// Loop up in the env var or use default value
func useEnvOrDefault(envVar, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}

func useEnvOrDefaultInt(envVar string, defaultValue int) int {
	if val, ok := os.LookupEnv(envVar); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultValue
}

// Get state file directory
func getDotatoDirUnsafe() (string) {
	// Check env var
	if val, ok := os.LookupEnv(DotatoDirPathEnv); ok {
		return val
	}

	// Check user config dir
	dir, err := os.UserConfigDir()
	if err == nil {
		return filepath.Join(dir, DotatoDirName)
	}

	// Check OS
	switch runtime.GOOS {
	}

	panic(`Oops! Dotato doesn't know your OS.
Please provide a user-wide directory with DOTATO_DIR env var
to let dotato save some files.`)
}

func readStateUnsafe(fs billy.Filesystem, isMem bool) (*state.State, error) {
	if isMem {
		return state.New(fs, state.PathInMemory)
	}

	path := filepath.Join(
		getDotatoDirUnsafe(),
		useEnvOrDefault(StateFileNameEnv, StateFileNameDefault),
	)
	
	return state.New(fs, path)
}

func readConfig(
	fs billy.Filesystem,
) (
	cfg *config.Config, cdir gp.GardenPath, err error,
) {
	// directory
	dir, err := gp.New(".")
	if err != nil {
		panic(err)
	}

	// config file name
	filename := useEnvOrDefault(ConfigFileNameEnv, ConfigFileNameDefault)

	cfg, cdir, err = config.ReadRecur(fs, dir, filename)
	return
}

func readIgnore(
	fs billy.Filesystem, dir gp.GardenPath,
) (
	ig *ignore.Ignore, err error,
) {
	// ignore file name
	filename := useEnvOrDefault(IgnoreFileNameEnv, IgnoreFileNameDefault)

	// Init ignore
	ig = ignore.NewWithFS(fs, dir, filename)

	// Read ignore file in dir
	_, err = ig.Read(dir)
	if err != nil {
		return nil, err
	}

	return
}

func readIgnoreRecur(
	fs billy.Filesystem, dir gp.GardenPath,
) (
	ig *ignore.Ignore, err error,
) {
	// ignore file name
	filename := useEnvOrDefault(IgnoreFileNameEnv, IgnoreFileNameDefault)

	// Init ignore
	ig = ignore.NewWithFS(fs, dir, filename)

	// Read ignore file in dir
	err = ig.ReadRecur(dir)
	if err != nil {
		return nil, err
	}

	return
}
