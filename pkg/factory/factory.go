package factory

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/state"
)

const (
	// Location of user-wide config directory
	DotatoDirPathEnv 			= "DOTATO_DIRECTORY"
	DotatoDirName					= "dotato"

	// State file name
	StateFileNameEnv 			= "DOTATO_STATE"
	StateFileNameDefault 	= "dotatostate.sqlite"

	// Config file name
	ConfigFileNameEnv 		= "DOTATO_CONFIG"
	ConfigFileNameDefault = "dotato.yaml"
	
	// Ignore file name
	IgnoreFileNameEnv 		= "DOTATO_IGNORE"
	IgnoreFileNameDefault = "dotato.ignore"
)

var (
	ErrConfigNotFound = fmt.Errorf("config file not found")
)

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

	panic("Oops, dotato doesn't know your OS. Please provide a user-wide directory with DOTATO_DIR env var to let dotato save some files")
}

func getStatePathUnsafe() string {
	return filepath.Join(
		getDotatoDirUnsafe(),
		useEnvOrDefault(StateFileNameEnv, StateFileNameDefault),
	)
}

func getConfigFileName() string {
	return useEnvOrDefault(ConfigFileNameEnv, ConfigFileNameDefault)
}

func getIgnoreFileName() string {
	return useEnvOrDefault(IgnoreFileNameEnv, IgnoreFileNameDefault)
}

/////////////////////////////////////////////////

func ReadStateUnsafe() (*state.State, error) {
	return state.New(getStatePathUnsafe())
}

func ReadConfig() (cfg *config.Config, base gp.GardenPath, err error) {
	dir, err := gp.New(".")
	if err != nil {
		panic(err)
	}

	cfg, base, err = config.ReadRecur(osfs.New("."), dir, getConfigFileName())
	return
}

func ReadIgnore(base gp.GardenPath) *ignore.Ignore {
	return ignore.New(base, getIgnoreFileName())
}
