package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// Test the NewDB function
	db, ver, err := newDB(PathInMemory)
	assert.NoError(t, err)
	assert.Equal(t, ver, Version1)

	// Test the version
	{
		getVer, ok, err := getVersion(db)
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.Equal(t, ver, getVer)
	}
}
