package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// parse config from string
	genCfg, err := parseConfigFromStr(t1s)
	assert.NoError(t, err)
	assert.True(t, genCfg.isEqual(t1c), "Generated config should be equal to the expected config")
}
