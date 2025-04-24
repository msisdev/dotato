package state

import (
	"testing"
	"time"

	"github.com/msisdev/dotato/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestV1(t *testing.T) {
	d, err := NewDB(PathInMemory)
	assert.NoError(t, err)

	now := time.Now()

	hs := []History{
		{ "t1", "s1", config.ModeFile, now, now, "h1"},
		{ "t2", "s2", config.ModeLink, now, now, "h2"},
		{ "t3", "s3", config.ModeFile, now, now, "h3"},
		{ "t4", "s4", config.ModeLink, now, now, "h4"},
	}

	// Upsert
	for _, h := range hs {
		err := d.v1_upsertOne(h)
		assert.NoError(t, err)
	}

	// Get all
	{
		all, err := d.v1_getAll()
		assert.NoError(t, err)
		assert.Equal(t, len(all), len(hs))
		assert.Equal(t, all[0].TargetPath, hs[0].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[1].TargetPath)
		assert.Equal(t, all[2].TargetPath, hs[2].TargetPath)
		assert.Equal(t, all[3].TargetPath, hs[3].TargetPath)
	}

	// Get all by mode
	{
		all, err := d.v1_getAllByMode(config.ModeFile)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].TargetPath, hs[0].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[2].TargetPath)

		all, err = d.v1_getAllByMode(config.ModeLink)
		assert.NoError(t, err)
		assert.Equal(t, len(all), 2)
		assert.Equal(t, all[0].TargetPath, hs[1].TargetPath)
		assert.Equal(t, all[1].TargetPath, hs[3].TargetPath)
	}

	// Get one
	{
		h, err := d.v1_getOne(hs[0].TargetPath)
		assert.NoError(t, err)
		assert.Equal(t, h.TargetPath, hs[0].TargetPath)
	}

	// Delete many
	{
		err := d.v1_deleteMany([]string{hs[2].TargetPath, hs[3].TargetPath})
		assert.NoError(t, err)

		hs, err := d.v1_getAll()
		assert.NoError(t, err)
		assert.Equal(t, len(hs), 2)
		assert.Equal(t, hs[0].TargetPath, "t1")
	}
}