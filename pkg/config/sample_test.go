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
		Groups: map[string]string{
			"home": "~",
		},
	}

	genCfg, err := NewFromString(GetSampleStr())
	assert.NoError(t, err)
	assert.True(t, cfg.isEqual(genCfg), "Generated config should be equal to the sample config")
}
