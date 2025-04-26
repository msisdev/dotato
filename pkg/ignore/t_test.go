package ignore

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

type FileType bool
const (
	IsDir				FileType = true
	IsFile			FileType = false
)

type IgnoreType bool
const (
	Ignored 		IgnoreType = true
	NotIgnored 	IgnoreType = false
)

type FileEntry struct {
	path 			string
	isDir 		FileType
	isIgnored	IgnoreType
}

func makeMemFS(entries []FileEntry) billy.Filesystem {
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
