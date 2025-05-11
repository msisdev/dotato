package engine

import (
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Returns corresponding dtt path
func (e Engine) DotToDtt(
	base gp.GardenPath,
	dot gp.GardenPath,
	group string,
) gp.GardenPath {
	path := e.cdir.Copy()
	path = append(path, group)
	path = append(path, dot[len(base):]...)
	return path
}

// Returns corresponding dot path
func (e Engine) DttToDot(
	base gp.GardenPath,
	dtt gp.GardenPath,
) gp.GardenPath {
	if len(dtt) < len(e.cdir)+1 {
		panic("DttToDot: invalid path")
	}

	path := base.Copy()
	for i := len(e.cdir)+1; i < len(dtt); i++ {
		path = append(path, dtt[i])
	}
	return path
}
