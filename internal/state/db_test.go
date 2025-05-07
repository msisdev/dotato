package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// Test the NewDB function
	db, ver, err := NewDB(PathInMemory)
	assert.NoError(t, err)
	assert.Equal(t, ver, Version1)

	// Test the version
	{
		getVer, ok, err := GetVersion(db)
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.Equal(t, ver, getVer)
	}
}
