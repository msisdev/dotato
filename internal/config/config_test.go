package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// parse config from string
	genCfg, err := NewFromString(testcase1String)
	assert.NoError(t, err)
	assert.True(t, genCfg.IsEqual(testcase1Config), "Generated config should be equal to the expected config")
}
