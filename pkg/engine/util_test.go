package engine

import (
	"testing"

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
	e := Engine{
		cdir: cdir,
	}
	resultDtt := e.DotToDtt(base, dot, group)
	assert.Equal(t, dtt.Abs(), resultDtt.Abs())
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
	e := Engine{
		cdir: cdir,
	}
	path := e.DttToDot(base, dtt)
	assert.Equal(t, dot.Abs(), path.Abs())
}
