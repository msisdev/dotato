package dotato

import (
	"os"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func TestDotToDtt(t *testing.T) {
	var (
		basepath 	= "/home/user"
		dotpath		= "/home/user/.bashrc"
		cdirpath	= "/home/user/Documents/dotato"
		group 		= "bash"
		dttpath 	= "/home/user/Documents/dotato/bash/.bashrc"
	)

	// Base path
	base, err := gp.New(basepath)
	assert.NoError(t, err)

	// Dotfile
	dot, err := gp.New(dotpath)
	assert.NoError(t, err)
	
	// Config dir
	cdir, err := gp.New(cdirpath)
	assert.NoError(t, err)

	// Dotato
	dtt, err := gp.New(dttpath)
	assert.NoError(t, err)

	// Test the function
	d := Dotato{
		cdir: cdir,
	}
	path := d.DotToDtt(base, dot, group)
	assert.Equal(t, dtt.Abs(), path.Abs())
}

func TestDttToDot(t *testing.T) {
	var (
		cdirpath 	= "/home/user/Documents/dotato"
		dttpath		= "/home/user/Documents/dotato/bash/.bashrc"
		basepath	= "/home/user"
		dotpath 	= "/home/user/.bashrc"
	)

	// Base path
	base, err := gp.New(basepath)
	assert.NoError(t, err)

	// Dotfile
	dot, err := gp.New(dotpath)
	assert.NoError(t, err)

	// Config dir
	cdir, err := gp.New(cdirpath)
	assert.NoError(t, err)

	// Dotato
	dtt, err := gp.New(dttpath)
	assert.NoError(t, err)

	// Test the function
	d := Dotato{
		cdir: cdir,
	}
	path := d.DttToDot(dtt, base)
	assert.Equal(t, dot.Abs(), path.Abs())
}

func TestCompareFile(t *testing.T) {
	var (
		fs = memfs.New()
		dtt = Dotato{fs: fs}
		path1 = "/file1.txt"
		path2 = "/file2.txt"
		content = []byte("Hello, World!")
	)

	// Compare equal files
	{
		file1, err := fs.Create(path1)
		assert.NoError(t, err)
		_, err = file1.Write(content)
		assert.NoError(t, err)
		err = file1.Close()
		assert.NoError(t, err)
	
		file2, err := fs.Create(path2)
		assert.NoError(t, err)
		_, err = file2.Write(content)
		assert.NoError(t, err)
		err = file2.Close()
		assert.NoError(t, err)
	
		// Compare
		equal, err := dtt.compareFile(path1, path2)
		assert.NoError(t, err)
		assert.True(t, equal)
	}

	// Compare different files
	{
		file1, err := fs.OpenFile(path1, os.O_RDWR, 0644)
		assert.NoError(t, err)
		_, err = file1.Write([]byte("Changed content"))
		assert.NoError(t, err)
		err = file1.Close()
		assert.NoError(t, err)

		// Compare
		equal, err := dtt.compareFile(path1, path2)
		assert.NoError(t, err)
		assert.False(t, equal)
	}
}

func TestMemfs(t *testing.T) {
	var (
		fs = memfs.New()
	)

	assert.Equal(t, "/", fs.Root())
	fs.MkdirAll("/home/user", os.ModePerm)
	fs.Create("/home/user/file.txt")
	fs.Chroot("/home/user")
	assert.Equal(t, "/", fs.Root())
	
	_, err := fs.Open("/file.txt")
	assert.Error(t, err)
}

func TestEvalSymlinks(t *testing.T) {
	var (
		fs = memfs.New()
		d = Dotato{fs: fs, isMem: true}
		path1 = "/link1"
		path2 = "/link2"
		path3 = "/link3"
		path3Alt = "./link3"
		path4 = "/link4"
	)
	
	// No symlink
	{
		// Create a file
		file, err := fs.Create(path1)
		assert.NoError(t, err)
		assert.NoError(t, file.Close())

		path, err := d.evalSymlinks(path1)
		assert.NoError(t, err)
		assert.Equal(t, path1, path)
	}

	// 1 symlink
	{
		// Create a symlink
		err := fs.Symlink(path1, path2)
		assert.NoError(t, err)

		path, err := d.evalSymlinks(path2)
		assert.NoError(t, err)
		assert.Equal(t, path1, path)
	}

	// 2 symlinks
	{
		// Create a symlink
		err := fs.Symlink(path2, path3)
		assert.NoError(t, err)

		path, err := d.evalSymlinks(path3)
		assert.NoError(t, err)
		assert.Equal(t, path1, path)
	}

	// 3 symlinks, test with relative path
	{
		// Create a symlink
		err := fs.Symlink(path3Alt, path4)
		assert.NoError(t, err)

		path, err := d.evalSymlinks(path4)
		assert.NoError(t, err)
		assert.Equal(t, path1, path)
	}
}
