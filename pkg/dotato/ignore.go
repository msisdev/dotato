package dotato

import (
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
)

func (d Dotato) CreateIgnoreBase() (err error) {
	return createAndWriteFile(d.fs, ".", []byte(ignore.SampleBaseIgnore))
}

func (d Dotato) CreateIgnoreHome() (err error) {
	return createAndWriteFile(d.fs, ".", []byte(ignore.SampleIgnore))
}

func (d *Dotato) ReadIgnoreBase() (err error) {
	if d.rt != nil {
		return
	}

	// Does config exist?
	if d.base == nil {
		ok, err := d.ReadConfig()
		if err != nil {
			return err
		}
		if !ok {
			return ErrConfigNotFound
		}
	} else {
		// init rule tree
		d.rt = ignore.NewRuleTreeFromBase(*d.base)
	}

	
	filename := getIgnoreFileName()
	return ignore.ReadIgnoreFileAdd(d.fs, d.rt, *d.base, filename)
}

func (d *Dotato) ReadIgnoreRecur(dir gp.GardenPath) (err error) {
	if d.rt == nil {
		if err = d.ReadIgnoreBase(); err != nil {
			return
		}
	}

	filename := getIgnoreFileName()
	return ignore.ReadIgnoreFileRecur(d.fs, d.rt, dir, filename)
}
