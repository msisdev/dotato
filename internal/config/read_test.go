package config

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	const filename = ConfigFileName
	fs := memfs.New()

	// Try to read a non-existing file
	_, ok, err := Read(fs, filename)
	assert.NoError(t, err, "Read should not return an error")
	assert.False(t, ok, "File should not exist")

	// Write
	err = Write(fs, filename, testcase1Config)
	assert.NoError(t, err, "Write should not return an error")

	// Read
	genCfg, ok, err := Read(fs, filename)
	assert.NoError(t, err, "Read should not return an error")
	assert.True(t, ok, "File should exist")
	assert.Equal(t, testcase1Config.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(testcase1Config), "Generated config should be equal to the expected config")

	// Read recursively
	root, err := gp.New("/")
	if err != nil {
		panic(err)
	}
	home, err := gp.New("~")
	if err != nil {
		panic(err)
	}
	if home.IsEqual(root) {
		panic("HOME is root")
	}
	genCfg, dir, err := ReadRecur(fs, home, filename)
	assert.NoError(t, err, "ReadRecur should not return an error")
	assert.Equal(t, gp.GardenPath{""}, dir, "Directory should be equal")
	assert.Equal(t, testcase1Config.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(testcase1Config), "Generated config should be equal to the expected config")
}

func TestReadRecur(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

	const filename = "dotato.yaml"
	fs := memfs.New()

	// Write config file in the root directory
	err := Write(fs, filename, testcase1Config)
	assert.NoError(t, err, "Write should not return an error")

	// Read config file recursively
	root, err := gp.New("/")
	if err != nil {
		panic(err)
	}
	genCfg, dir, err := ReadRecur(fs, root, filename)
	assert.NoError(t, err, "ReadRecur should not return an error")
	assert.Equal(t, root, dir, "Directory should be equal")
	assert.Equal(t, testcase1Config.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(testcase1Config), "Generated config should be equal to the expected config")
}

func TestReadWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows path test on non-Windows OS")
		return
	}

	filepath := "C:\\Users\\mslit\\Work\\GitHub\\dotato\\dotato.yaml"
	fs := filesystem.NewOSFS()

	cfg, ok, err := Read(fs, filepath)
	assert.NoError(t, err, "Read should not return an error")
	assert.True(t, ok, "File should exist")
	assert.NotNil(t, cfg, "Config should not be nil")
}

func TestReadRecurWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows path test on non-Windows OS")
		return
	}

	const filename = "dotato.yaml"
	fs := memfs.New()

	// Read config file recursively
	wd, err := gp.New(".")
	if err != nil {
		panic(err)
	}

	// Write config file in the root directory
	root := wd[0] + "\\"
	configPath := filepath.Join(root, filename)
	err = Write(fs, configPath, testcase1Config)
	assert.NoError(t, err, "Write should not return an error")

	// ReadRecur
	genCfg, dir, err := ReadRecur(fs, wd, filename)
	assert.NoError(t, err, "ReadRecur should not return an error")
	assert.Equal(t, wd[:1], dir, "Directory should be equal")
	assert.Equal(t, testcase1Config.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(testcase1Config), "Generated config should be equal to the expected config")
}
