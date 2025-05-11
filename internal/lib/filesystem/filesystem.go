package filesystem

import (
	"runtime"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

// What is this function doing?
//
// dotato depends on "github.com/go-git/go-billy/v5".
// It provides a abstract interface for "github.com/go-git/go-billy/v5/osfs"
// and "github.com/go-git/go-billy/v5/memfs".
//
// osfs.New("") occurs errors in linux filesystem when it is used with
// Readlink(absPath).
func NewOSFS() billy.Filesystem {
	if runtime.GOOS == "windows" {
		return osfs.New("")
	}
	return osfs.New("/")
}
