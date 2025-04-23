package state

import (
	"testing"
	"time"

	"github.com/msisdev/dotato/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestV1(t *testing.T) {
	d, err := NewDB(DBPathInMemory)
	assert.NoError(t, err)

	now := time.Now()

	hs := []V1History{
		{ "t1", "s1", config.ModeFile, now, now, "h1"},
		{ "t2", "s2", config.ModeLink, now, now, "h2"},
		{ "t3", "s3", config.ModeFile, now, now, "h3"},
		{ "t4", "s4", config.ModeLink, now, now, "h4"},
	}

	// Upsert
	for _, h := range hs {
		err := d.V1_UpsertOne(h)
		assert.NoError(t, err)
	}

	// Get all
	{
		all, err := d.V1_GetAll()
		assert.NoError(t, err)
		assert.Equal(t, len(all), len(hs))
		assert.Equal(t, all[0].TargetPath, hs[0].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[1].TargetPath)
		assert.Equal(t, all[2].TargetPath, hs[2].TargetPath)
		assert.Equal(t, all[3].TargetPath, hs[3].TargetPath)
	}

	// Get one
	{
		h, err := d.V1_GetOne(hs[0].TargetPath)
		assert.NoError(t, err)
		assert.Equal(t, h.TargetPath, hs[0].TargetPath)
	}

	// Delete many
	{
		err := d.V1_DeleteMany([]string{hs[2].TargetPath, hs[3].TargetPath})
		assert.NoError(t, err)

		hs, err := d.V1_GetAll()
		assert.NoError(t, err)
		assert.Equal(t, len(hs), 2)
		assert.Equal(t, hs[0].TargetPath, "t1")
	}
}