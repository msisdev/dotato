package config

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	const filename = "dotato.yaml"
	fs := memfs.New()

	// Write
	err := WriteFile(fs, filename, t1c)
	assert.NoError(t, err, "WriteConfigFile should not return an error")

	// Read
	genCfg, err := ReadFile(fs, filename)
	assert.NoError(t, err, "ReadConfigFile should not return an error")
	assert.Equal(t, t1c.Version, genCfg.Version, "Version should be equal")
	assert.True(t, genCfg.IsEqual(t1c), "Generated config should be equal to the expected config")
}
