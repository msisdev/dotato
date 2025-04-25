package ignore

import (
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

type FileHelper struct {
	fs billy.Filesystem
}

func NewFileHelper(es []Entry, ies []IgnoreEntry) *FileHelper {
	// Create files
	h := &FileHelper{
		fs: NewMemFS(es),
	}

	// Create ignore files
	for _, ie := range ies {
		h.createIgnoreFile(ie)
	}

	return h
}

func (h *FileHelper) createIgnoreFile(ie IgnoreEntry) {
	// Get path object
	path, err := gp.New(ie.path)
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
	h := NewFileHelper(t1e, t1i)

	// Build rule tree from root dir
	root, err := gp.New("/")
	if err != nil {
		panic(err)
	}
	tree := NewRuleTree(0)
	err = ReadIgnoreFileRecur(h.fs, tree, root, DefaultIgnoreFileName)
	assert.NoError(t, err)

	// Test rules on each file
	for _, e := range t1e {
		path, err := gp.New(e.path)
		if err != nil {
			panic(err)
		}

		assert.Equal(
			t, e.isIgnored, tree.Ignore(path),
			"Path %s should be ignored: %v", e.path, e.isIgnored,
		)
	}
}
