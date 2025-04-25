package dotato

import (
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/state"
)

func (d *Dotato) ReadState() (err error) {
	if d.state != nil {
		return
	}

	d.state, err = state.New(getStatePathUnsafe())
	return
}

func (d *Dotato) ReadConfig() (ok bool, err error) {
	if d.cfg != nil {
		return
	}

	filename := getConfigFileName()
	separator := getConfigPathSeparator()

	// Get pwd
	pwd, err := gp.NewWithSep(".", separator)
	if err != nil {
		return
	}

	// Search config file
	cfg, base, err := config.ReadRecur(d.fs, pwd, filename)
	if err != nil {
		return
	}
	d.base = &base
	d.cfg = cfg

	return true, nil
}

func (d *Dotato) ReadBaseIgnore() (err error) {
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
		if err = d.ReadBaseIgnore(); err != nil {
			return
		}
	}

	filename := getIgnoreFileName()
	return ignore.ReadIgnoreFileRecur(d.fs, d.rt, dir, filename)
}
