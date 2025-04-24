package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	genCfg, err := NewConfigFromStr(t1s)
	assert.NoError(t, err)
	assert.True(t, genCfg.IsEqual(t1c), "Generated config should be equal to the expected config")
}
