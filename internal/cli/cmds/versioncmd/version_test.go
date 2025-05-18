package versioncmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestVersion(t *testing.T) {
	version, err := getLatestVersion(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, version)
}

func TestGetCurrentVersion(t *testing.T) {
	version := getCurrentVersion()
	assert.NotEmpty(t, version)
	assert.Equal(t, version, "v0.0.0")
}
