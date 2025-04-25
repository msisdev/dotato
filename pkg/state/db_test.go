package state

import (
	"testing"

	"github.com/msisdev/dotato/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	// Test the NewDB function
	d, err := NewDB(PathInMemory)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	// Test the AutoMigrate function
	ver, ok, err := d.GetVersion()
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, config.DotatoVersion(), ver)	
}
