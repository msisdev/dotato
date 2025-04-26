package dotato

import (
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// func (d Dotato) CreateConfig(cfg *config.Config) (err error) {
// 	if d.base != nil {
// 		return nil
// 	}
	
// 	return config.Write(d.fs, getConfigFileName(), cfg)
// }

func (d Dotato) CreateConfigSample() (err error) {
	if d.base != nil {
		return nil
	}

	//  Set base to working directory
	*d.base, err = gp.New(".")
	if err != nil {
		return
	}
	
	// Create file
	file, err := d.fs.Create(getConfigFileName())
	if err != nil {
		return
	}
	defer file.Close()

	// Write sample config
	_, err = file.Write([]byte(config.GetSampleConfigStr()))
	
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

func (d *Dotato) WriteConfig(cfg *config.Config) (err error) {
	if d.base == nil {
		return nil
	}

	return config.Write(d.fs, getConfigFileName(), cfg)
}
