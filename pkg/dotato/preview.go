package dotato

import (
	"os"
	"path/filepath"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

type PreviewImportFile struct {
	Dot 			gp.GardenPath
	DotReal 	gp.GardenPath	// nil if dotfile is not symlink
	Dtt 			gp.GardenPath
	DttExists	bool
	Equal			bool
}

// Analyze current state of dotfile and dotato file
func (d Dotato) PreviewImportFile(
	dot gp.GardenPath,
	dtt gp.GardenPath,
) (*PreviewImportFile, error) {
	p := &PreviewImportFile{
		Dot: dot,
		DotReal: nil,
		Dtt: dtt,
		DttExists: false,
		Equal: false,
	}

	// Check dotfile
	var dotPath string	// should be real path
	{
		abs := dot.Abs()

		// Get real path
		// realPath, err := filepath.EvalSymlinks(abs)
		if err != nil {
			return nil, err
		}

		if realPath == abs {
			p.DotReal = nil
			dotPath = abs
		} else {
			p.DotReal, err = gp.New(realPath)
			if err != nil {
				return nil, err
			}
			dotPath = realPath
		}
	}

	// Check dotato file
	var dttPath = dtt.Abs()
	{
		if _, err := d.fs.Stat(dttPath); err != nil {
			if os.IsNotExist(err) {
				p.DttExists = false
				return p, nil
			}

			return nil, err
		}
	}

	// Compare
	if dotPath == dttPath {
		p.Equal = true
		return p, nil
	}
	equal, err := d.compareFile(dotPath, dttPath)
	if err != nil {
		return nil, err
	}
	p.Equal = equal

	return p, nil
}