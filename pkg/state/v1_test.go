package state

import (
	"testing"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestV1_GetAllByMode(t *testing.T) {
	fs := memfs.New()
	state, err := New(fs, PathInMemory)
	assert.NoError(t, err)

	now := time.Now()
	hs := []History{
		{"t1", "s1", config.ModeFile, now, now},
		{"t2", "s2", config.ModeLink, now, now},
		{"t3", "s3", config.ModeFile, now, now},
		{"t4", "s4", config.ModeLink, now, now},
	}

	// Insert
	{
		for _, h := range hs {
			err := state.v1_upsertOne(h)
			assert.NoError(t, err)
		}
	}

	// Get all by mode
	{
		all, err := state.v1_getAllByMode(config.ModeFile)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].DotPath, hs[0].DotPath)
		assert.Equal(t, all[1].DotPath, hs[2].DotPath)

		all, err = state.v1_getAllByMode(config.ModeLink)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].DotPath, hs[1].DotPath)
		assert.Equal(t, all[1].DotPath, hs[3].DotPath)
	}
}
