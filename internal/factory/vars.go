package factory

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	// Location of user-wide config directory
	dotatoDirPathEnv = "DOTATO_DIRECTORY_PATH"
	dotatoDirName    = "dotato"

	// State file name
	stateFileNameEnv     = "DOTATO_STATE_FILENAME"
	stateFileNameDefault = "dotatostate.sqlite"

	// Config file name
	configFileNameEnv     = "DOTATO_CONFIG_FILENAME"
	configFileNameDefault = "dotato.yaml"

	// Ignore file name
	ignoreFileNameEnv     = "DOTATO_IGNORE_FILENAME"
	ignoreFileNameDefault = ".dotatoignore"

	// Max iteration for file system IO
	maxFSIterEnv     = "DOTATO_MAX_FS_ITER"
	maxFSIterDefault = 10000
)

var (
	DotatoFileNameConfig = useEnvOrDefault(configFileNameEnv, configFileNameDefault)
	DotatoFileNameIgnore = useEnvOrDefault(ignoreFileNameEnv, ignoreFileNameDefault)
	DotatoDirPath        = getDotatoDirUnsafe()
	DotatoFileNameState  = useEnvOrDefault(stateFileNameEnv, stateFileNameDefault)
	DotatoFilePathState  = filepath.Join(DotatoDirPath, DotatoFileNameState)
	DotatoMaxFSIter      = useEnvOrDefaultInt(maxFSIterEnv, maxFSIterDefault)
	DotatoFileNames      = map[string]bool{
		DotatoFileNameConfig: true,
		DotatoFileNameIgnore: true,
		DotatoFileNameState:  true,
	}
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
func getDotatoDirUnsafe() string {
	// Check env var
	if val, ok := os.LookupEnv(dotatoDirPathEnv); ok {
		return val
	}

	// Check user config dir
	dir, err := os.UserConfigDir()
	if err == nil {
		return filepath.Join(dir, dotatoDirName)
	}

	// Check OS
	switch runtime.GOOS {
	}

	panic(`Oops! Dotato doesn't know your OS.
Please provide a user-wide directory with DOTATO_DIR env var
to let dotato save some files.`)
}
