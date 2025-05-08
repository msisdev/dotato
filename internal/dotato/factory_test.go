package dotato

import (
	"testing"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	fs := osfs.New("")
	_, _, err := readConfig(fs)
	assert.NoError(t, err)
}