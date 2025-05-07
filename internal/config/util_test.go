package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandEnv(t *testing.T) {
	envName := "TEST"
	envValue := "good"

	strs := []string{}
	exps := []string{}

	strs = append(strs, "$"+envName)
	exps = append(exps, envValue)

	strs = append(strs, "$"+envName+"/")
	exps = append(exps, envValue+"/")

	strs = append(strs, "${"+envName+"}")
	exps = append(exps, envValue)

	strs = append(strs, "${"+envName+"}/")
	exps = append(exps, envValue+"/")

	// Test with no env var
	for i := 0; i < len(strs); i++ {
		_, notFound := expandEnv(strs[i])
		assert.Equal(t, envName, notFound[0], "Expected %s, got %s", envName, notFound[0])
	}

	os.Setenv(envName, envValue)
	for i := 0; i < len(strs); i++ {
		exp, notFound := expandEnv(strs[i])
		assert.Equal(t, exps[i], exp, "Expected %s, got %s", exps[i], exp)
		assert.Empty(t, notFound, "Expected empty, got %s", notFound)
	}
}
