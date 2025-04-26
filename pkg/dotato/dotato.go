package dotato

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
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

	// Separator
	SeparatorEnv 					= "DOTATO_SEPARATOR"
	SeparatorDefault 			= string(gp.DefaultSeparator)
	
	// Ignore file name
	IgnoreFileNameEnv 		= "DOTATO_IGNORE"
	IgnoreFileNameDefault = "dotato.ignore"
)

var (
	ErrConfigNotFound = fmt.Errorf("config file not found")
)

type Dotato struct {
	fs 			billy.Filesystem

	base	*gp.GardenPath
	cfg		*config.Config
	rt		*ignore.RuleTree
	state	*state.State
}

// New Dotato instance with filesystem
func NewDotato() *Dotato {
	return &Dotato{
		fs: osfs.New("."),
	}
}

// New Dotato instance with memory filesystem
func NewDotatoMemfs() *Dotato {
	return &Dotato{
		fs: memfs.New(),
	}
}

func getDotatoDirUnsafe() (string) {
	// Look up env var
	if val, ok := os.LookupEnv(DotatoDirPathEnv); ok {
		return val
	}

	// Try to get user config dir
	dir, err := os.UserConfigDir()
	if err == nil {
		return filepath.Join(dir, DotatoDirName)
	}

	// Determine OS
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

func getConfigPathSeparator() rune {
	return rune(useEnvOrDefault(SeparatorEnv, SeparatorDefault)[0])
}

func getIgnoreFileName() string {
	return useEnvOrDefault(IgnoreFileNameEnv, IgnoreFileNameDefault)
}
