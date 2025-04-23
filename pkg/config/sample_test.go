package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleConfig(t *testing.T) {
	cfg := &Config{
		Version: GetDotatoVersion(),
		Mode:	ModeFile,
		Plans: map[string]GroupList{
			"desktop": {"bash"},
		},
		Groups: map[string]string{
			"bash": "~",
		},
	}

	genCfg, err := NewFromStr(GetSampleConfigStr())
	assert.NoError(t, err)
	assert.True(t, cfg.IsEqual(genCfg), "Generated config should be equal to the sample config")
}
