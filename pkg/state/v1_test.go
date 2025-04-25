package state

import (
	"testing"
	"time"

	"github.com/msisdev/dotato/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestV1_GetAllByMode(t *testing.T) {
	s, err := New(StatePathInMemory)
	assert.NoError(t, err)

	now := time.Now()

	hs := []History{
		{ "t1", "s1", config.ModeFile, now, now, "h1"},
		{ "t2", "s2", config.ModeLink, now, now, "h2"},
		{ "t3", "s3", config.ModeFile, now, now, "h3"},
		{ "t4", "s4", config.ModeLink, now, now, "h4"},
	}

	// Insert
	{
		for _, h := range hs {
			err := s.v1_upsertOne(h)
			assert.NoError(t, err)
		}
	}

	// Get all by mode
	{
		all, err := s.v1_getAllByMode(config.ModeFile)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].TargetPath, hs[0].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[2].TargetPath)

		all, err = s.v1_getAllByMode(config.ModeLink)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].TargetPath, hs[1].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[3].TargetPath)
	}
}
