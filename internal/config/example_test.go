package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleConfig(t *testing.T) {
	cfg := &Config{
		Version: ConfigVersion,
		Mode:    ModeFile,
		Plans: map[string][]string{
			"all":  nil,
			"mypc": {"bash"},
		},
		Groups: map[string]map[string]string{
			"bash": {
				"nux": "~",
				"win": "$USERPROFILE",
			},
		},
	}

	genCfg, err := NewFromString(GetExample())
	assert.NoError(t, err)
	assert.True(t, cfg.IsEqual(genCfg), "Generated config should be equal to the sample config")
}
