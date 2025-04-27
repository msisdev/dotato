package ignore

import (
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func TestIgnore(t *testing.T) {
	fs := memfs.New()

	// Create ignore files
	for _, ie := range testcase1Ignore {
		// Get path object
		path, err := gp.New(ie.path)
		if err != nil {
			panic(err)
		}

		// Create directories
		if err = fs.MkdirAll(path[:len(path)-1].Abs(), 0755); err != nil {
			panic(err)
		}

		// Create file
		file, err := fs.Create(path.Abs())
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Write lines
		content := strings.Join(ie.lines, "\n")
		if _, err = file.Write([]byte(content)); err != nil {
			panic(err)
		}
	}

	// Create a new Ignore instance
	base := gp.Root
	ig := NewWithFS(fs, base, IgnoreFileName)
	err := ig.ReadRecur(base)
	assert.NoError(t, err)

	// Test files
	for _, fe := range testcase1Files {
		path, err := gp.New(fe.path)
		if err != nil {
			panic(err)
		}

		isIgnored := ig.IsIgnored(path)

		assert.Equal(t, fe.isIgnored, isIgnored, "Expected %v for %s", fe.isIgnored, fe.path)
	}
}
