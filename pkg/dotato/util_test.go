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
