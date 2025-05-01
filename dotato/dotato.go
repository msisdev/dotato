package dotato

import (
	"io/fs"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/state"
)

// Dotato is a kind of engine.
// It handles building blocks and exposes high level functions.
type Dotato struct {
	fs    	billy.Filesystem
	isMem 	bool
	maxIter	int								// max iteration for file system execution

	cdir  	gp.GardenPath			// config directory
	cfg   	*config.Config
	ig   		*ignore.Ignore 		// ignore file in config directory
	state 	*state.State
}

// New Dotato instance with filesystem
func New() *Dotato {
	return &Dotato{
		fs:    osfs.New("/"),
		isMem: false,
		maxIter: useEnvOrDefaultInt(MaxFileSystemIterEnv, MaxFileSystemIterDefault),
	}
}

func NewMemfs() *Dotato {
	return &Dotato{
		fs:    memfs.New(),
		isMem: true,
		maxIter: useEnvOrDefaultInt(MaxFileSystemIterEnv, MaxFileSystemIterDefault),
	}
}

func (d *Dotato) setConfig() (err error) {
	if d.cfg != nil {
		return
	}

	// Config
	d.cfg, d.cdir, err = readConfig(d.fs)
	return
}

func (d *Dotato) setState() (err error) {
	if d.state != nil {
		return
	}

	d.state, err = readStateUnsafe(d.fs, d.isMem)
	return
}

func (d *Dotato) setIgnore() (err error) {
	if d.ig != nil {
		return
	}

	// base is required
	err = d.setConfig()
	if err != nil {
		return
	}

	d.ig, err = readIgnore(d.fs, d.cdir)
	return
}

func (d Dotato) GetConfigVersion() (string, error) {
	if err := d.setConfig(); err != nil {
		return "", err
	}

	return d.cfg.Version, nil
}

func (d Dotato) GetConfigMode() (string, error) {
	if err := d.setConfig(); err != nil {
		return "", err
	}

	return d.cfg.Mode, nil
}

func (d Dotato) GetConfigPlans() (map[string][]string, error) {
	if err := d.setConfig(); err != nil {
		return nil, err
	}

	// Get plans
	return d.cfg.Plans, nil
}

func (d Dotato) GetConfigGroups(plan string) (groups map[string]bool, err error) {
	if err = d.setConfig(); err != nil {
		return
	}

	// Get groups
	groups = d.cfg.GetGroups(plan)
	return
}

func (d Dotato) GetConfigGroupBase(group, resolver string) (base gp.GardenPath, notFound []string, err error) {
	if err = d.setConfig(); err != nil {
		return
	}

	// Get base of group
	base, notFound, err = d.cfg.GetGroupBase(group, resolver)
	return
}

/////////////////////////////////////////////////

// A type for both file and directory
type Entity struct {
	Path gp.GardenPath
	Info fs.FileInfo
}

// Scan which files will be imported
func (d Dotato) GetImportPaths(group string, base gp.GardenPath) (es []Entity, err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setIgnore(); err != nil { return }

	ig, err := readIgnoreRecur(d.fs, append(d.cdir, group))
	if err != nil {
		return
	}

	return d.walk(base, ig)
}

func (d Dotato) GetExportPaths(group string) (es []Entity, err error) {
	if err = d.setConfig(); err != nil { return }
	if err = d.setIgnore(); err != nil { return }

	ig, err := readIgnoreRecur(d.fs, append(d.cdir, group))
	if err != nil {
		return
	}

	return d.walk(append(d.cdir, group), ig)
}

/////////////////////////////////////////////////

func (d Dotato) GetAllHistoryByMode(mode string) ([]state.History, error) {
	if err := d.setState(); err != nil { return nil, err }
	
	return d.state.GetAllByMode(mode)
}

func (d Dotato) PutHistory(h state.History) (err error) {
	if err = d.setState(); err != nil { return }

	return d.state.UpsertOne(h)
}
