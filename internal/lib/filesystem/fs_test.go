// This file tests behavior of the external library go-billy

package filesystem

import (
	"os"
	"runtime"
	"testing"

	"github.com/go-git/go-billy/v6/memfs"
	"github.com/stretchr/testify/assert"
)

func TestFS_Stat(t *testing.T) {
	var (
		fs       = memfs.New()
		filepath = "/file"
		// filename = "file"
		linkpath = "/link"
		linkname = "link"
		deeppath = "/deep"
		deepname = "deep"
		content  = []byte("Hello, World!")
	)

	// Prepare
	{
		// Create a file
		file, err := fs.Create(filepath)
		assert.NoError(t, err)
		_, err = file.Write(content)
		assert.NoError(t, err)
		err = file.Close()
		assert.NoError(t, err)

		// Create a symlink
		err = fs.Symlink(filepath, linkpath)
		assert.NoError(t, err)

		// Create a deep symlink
		err = fs.Symlink(linkpath, deeppath)
		assert.NoError(t, err)
	}

	// Stat(linkpath): file
	{
		info, err := fs.Stat(linkpath)
		assert.NoError(t, err)
		assert.Equal(t, info.Name(), linkname)
		assert.Equal(t, info.Size(), int64(len(content)))
	}

	// Lstat(linkpath): link
	{
		info, err := fs.Lstat(linkpath)
		assert.NoError(t, err)
		assert.Equal(t, info.Name(), linkname)
		assert.Equal(t, info.Size(), int64(len(filepath)))
		assert.Equal(t, info.Mode().Type()&os.ModeType, os.ModeSymlink)
	}

	// Stat(deeppath): file
	{
		info, err := fs.Stat(deeppath)
		assert.NoError(t, err)
		assert.Equal(t, info.Name(), deepname)
		assert.Equal(t, info.Size(), int64(len(content)))
	}

	// Lstat(deeppath): deep
	{
		info, err := fs.Lstat(deeppath)
		assert.NoError(t, err)
		assert.Equal(t, info.Name(), deepname)
		assert.Equal(t, info.Size(), int64(len(linkpath)))
		assert.Equal(t, info.Mode().Type()&os.ModeType, os.ModeSymlink)
	}

	// Conlusion:
	//  - fs.Stat() name: itself, content: final target
	//  - fs.Lstat() name: itself, content: itself
}

func TestFS_Open(t *testing.T) {
	var (
		fs       = memfs.New()
		filepath = "/file"
		linkpath = "/link"
		deeppath = "/deep"
		content  = []byte("Hello, World!")
		bufsiz   = len(content)
	)

	// Prepare
	{
		file, err := fs.Create(filepath)
		assert.NoError(t, err)
		_, err = file.Write(content)
		assert.NoError(t, err)
		err = file.Close()
		assert.NoError(t, err)

		err = fs.Symlink(filepath, linkpath)
		assert.NoError(t, err)

		err = fs.Symlink(linkpath, deeppath)
		assert.NoError(t, err)
	}

	// Open(linkpath) -> file
	{
		file, err := fs.Open(linkpath)
		assert.NoError(t, err)

		buf := make([]byte, bufsiz)
		_, err = file.Read(buf)
		assert.NoError(t, err)
		assert.Equal(t, buf, content)

		err = file.Close()
		assert.NoError(t, err)
	}

	// Open(deeppath) -> file
	{
		file, err := fs.Open(deeppath)
		assert.NoError(t, err)

		buf := make([]byte, bufsiz)
		_, err = file.Read(buf)
		assert.NoError(t, err)
		assert.Equal(t, buf, content)

		err = file.Close()
		assert.NoError(t, err)
	}
}

func TestFS_Symlink(t *testing.T) {
	var (
		fs       = memfs.New()
		filepath = "/file"
		linkpath = "/link"
		content  = []byte("Hello, World!")
	)

	// Create linkpath as file
	{
		file, err := fs.Create(linkpath)
		assert.NoError(t, err)
		_, err = file.Write(content)
		assert.NoError(t, err)
		err = file.Close()
		assert.NoError(t, err)
	}

	// Overwrite linkpath as symlink
	{
		err := fs.Symlink(filepath, linkpath)
		assert.Error(t, err)
	}

	err := fs.Remove(linkpath)
	assert.NoError(t, err)

	// Create linkpath as link
	{
		err := fs.Symlink(filepath, linkpath)
		assert.NoError(t, err)
	}

	// Overwrite linkpath as symlink
	{
		err := fs.Symlink(filepath, linkpath)
		assert.Error(t, err)
	}

	// Conclusion: fs.Symlink() doesn't allow overwriting any existing file
}

func TestWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-windows platform")
		return
	}

	fs := NewOSFS()

	// Create a file
	file, err := fs.Create("testfile.txt")
	assert.NoError(t, err)
	_, err = file.Write([]byte("Hello, World!"))
	assert.NoError(t, err)

	// Create a symlink
	err = fs.Symlink("testfile.txt", "testlink")
	assert.NoError(t, err)
}
