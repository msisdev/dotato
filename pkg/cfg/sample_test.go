package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleConfig(t *testing.T) {
	cfg := &Config{
		Version: getDotatoVersion(),
		Plans:   map[string]GroupList{
			"arch": {"alacritty"},
		},
		Groups: map[string]string{
			"alacritty": "~/.config/alacritty",
		},
	}

	genCfg, err := NewConfigFromStr(GetSampleConfigStr())
	assert.NoError(t, err)

	assert.True(t, cfg.IsEqual(genCfg), "Generated config should be equal to the sample config")
}
