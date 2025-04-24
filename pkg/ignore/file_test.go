package ignore

import (
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

type FileTestHelper struct {
	fs billy.Filesystem
}

func NewFileTestHelper(es []Entry, ies []IgnoreEntry) *FileTestHelper {
	// Create files
	h := &FileTestHelper{
		fs: NewMemFS(es),
	}

	// Create ignore files
	for _, ie := range ies {
		h.createIgnoreFile(ie)
	}

	return h
}

func (h *FileTestHelper) createIgnoreFile(ie IgnoreEntry) {
	// Get path object
	path, err := gardenpath.NewGardenPath(ie.path)
	if err != nil {
		panic(err)
	}

	// Create directories
	if err = h.fs.MkdirAll(path[:len(path)-1].String(), 0755); err != nil {
		panic(err)
	}

	// Create file
	file, err := h.fs.Create(path.String())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write lines
	content := strings.Join(ie.lines, "\n")
	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

func TestCompileIgnoreFile1(t *testing.T) {
	h := NewFileTestHelper(t1e, t1i)

	gp, err := gardenpath.NewGardenPath("/")
	if err != nil {
		panic(err)
	}

	tree := NewRuleTree(0)

	err = CompileIgnoreFileRecur(h.fs, tree, gp, DefaultIgnoreFileName)
	assert.NoError(t, err)

	for _, e := range t1e {
		path, err := gardenpath.NewGardenPath(e.path)
		if err != nil {
			panic(err)
		}

		assert.Equal(
			t, e.isIgnored, tree.Ignore(path),
			"Path %s should be ignored: %v", e.path, e.isIgnored,
		)
	}
}
