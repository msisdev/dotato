package ignore

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

const (
	IsDir = true
	IsFile = false
	Ignored = true
	NotIgnored = false
)

// Entry represents a file or directory
type Entry struct {
	path 			string
	isDir 		bool
	isIgnored	bool
}

func NewMemFS(entries []Entry) billy.Filesystem {
	fs := memfs.New()

	// Add entries
	for _, e := range entries {
		switch (e.isDir) {
		case IsDir:
			err := fs.MkdirAll(e.path, 0755)
			if err != nil {
				panic(err)
			}
	
		case IsFile:
			_, err := fs.Create(e.path)
			if err != nil {
				panic(err)
			}
	
		default:
			panic("unknown object type")
		}
	}

	return fs
}

type IgnoreEntry struct {
	path string
	lines []string
}
