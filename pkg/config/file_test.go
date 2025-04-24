package config

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	const filename = ConfigFileName
	fs := memfs.New()

	// Try to read a non-existing file
	_, ok, err := ReadConfigFile(fs, filename)
	assert.NoError(t, err, "ReadConfigFile should not return an error")
	assert.False(t, ok, "File should not exist")

	// Write
	err = WriteConfigFile(fs, filename, t1c)
	assert.NoError(t, err, "WriteConfigFile should not return an error")

	// Read
	genCfg, ok, err := ReadConfigFile(fs, filename)
	assert.NoError(t, err, "ReadConfigFile should not return an error")
	assert.True(t, ok, "File should exist")
	assert.Equal(t, t1c.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(t1c), "Generated config should be equal to the expected config")

	// Read recursively
	root, err := gp.NewGardenPath("/")
	if err != nil {
		panic(err)
	}
	home, err := gp.NewGardenPath("~")
	if err != nil {
		panic(err)
	}
	if home.IsEqual(root) {
		panic("HOME is root")
	}
	genCfg, dir, err := ReadConfigFileRecur(fs, home, filename)
	assert.NoError(t, err, "ReadConfigFileRecur should not return an error")
	assert.Equal(t, gp.GardenPath{""}, dir, "Directory should be equal")
	assert.Equal(t, t1c.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(t1c), "Generated config should be equal to the expected config")
}
