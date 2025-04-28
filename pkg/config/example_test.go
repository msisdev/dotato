package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleConfig(t *testing.T) {
	cfg := &Config{
		Version: DotatoVersion(),
		Mode:	ModeFile,
		Plans: map[string][]string{
			"all": nil,
		},
		Groups: map[string]map[string]string{
			"home": {
				"nux": "~",
			},
		},
	}

	genCfg, err := NewFromString(GetExample())
	assert.NoError(t, err)
	assert.True(t, cfg.IsEqual(genCfg), "Generated config should be equal to the sample config")
}
