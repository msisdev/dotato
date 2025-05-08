package dotato

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/state"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/internal/ignore"
)

// Dotato is a kind of engine.
// It handles building blocks and exposes high level functions.
type Dotato struct {
	fs      billy.Filesystem
	isMem   bool
	maxIter int // max iteration for file system execution

	cdir  gp.GardenPath // config directory
	cfg   *config.Config
	ig    *ignore.Ignore // ignore file in config directory
	state *state.State
}

// New Dotato instance with filesystem
func New() *Dotato {
	return &Dotato{
		fs:    osfs.New(""),
		isMem: false,
		maxIter: useEnvOrDefaultInt(
			MaxFileSystemIterEnv,
			MaxFileSystemIterDefault,
		),
	}
}

func NewMemfs() *Dotato {
	return &Dotato{
		fs:    memfs.New(),
		isMem: true,
		maxIter: useEnvOrDefaultInt(
			MaxFileSystemIterEnv,
			MaxFileSystemIterDefault,
		),
	}
}

func NewWithFS(fs billy.Filesystem, isMem bool) *Dotato {
	return &Dotato{
		fs:    fs,
		isMem: isMem,
		maxIter: useEnvOrDefaultInt(
			MaxFileSystemIterEnv,
			MaxFileSystemIterDefault,
		),
	}
}

func (d *Dotato) setConfig() (err error) {
	if d.cfg != nil {
		return
	}

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
	if err = d.setConfig(); err != nil {
		return
	}

	d.ig, err = readIgnore(d.fs, d.cdir)
	return
}

///////////////////////////////////////////////////////////////////////////////

func (d Dotato) GetConfigDir() (gp.GardenPath, error) {
	if err := d.setConfig(); err != nil {
		return nil, err
	}

	return d.cdir, nil
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

func (d Dotato) GetConfigGroups(
	plan string,
) (
	groups map[string]bool,
	err error,
) {
	if err = d.setConfig(); err != nil {
		return
	}

	// Get groups
	groups = d.cfg.GetGroups(plan)
	return
}

// Get resolved dotfile base of group.
// May return not found env vars
func (d Dotato) GetConfigGroupBase(
	group, resolver string,
) (
	base gp.GardenPath,
	notFound []string,
	err error,
) {
	if err = d.setConfig(); err != nil {
		return
	}

	// Get resolved dotfile base of group
	base, notFound, err = d.cfg.GetGroupBase(group, resolver)
	return
}

func (d Dotato) GetConfigResolvers() (rs map[string]string, err error) {
	if err = d.setConfig(); err != nil {
		return
	}

	rs = make(map[string]string)

	// For each group
	for _, resolvers := range d.cfg.Groups {
		// For each resolver
		for name, resolver := range resolvers {
			// Collect resolver
			rs[name] = resolver
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////

func (d Dotato) GetAllHistoryByMode(mode string) ([]state.History, error) {
	if err := d.setState(); err != nil {
		return nil, err
	}

	return d.state.GetAllByMode(mode)
}

func (d Dotato) PutHistory(h state.History) (err error) {
	if err = d.setState(); err != nil {
		return
	}

	return d.state.UpsertOne(h)
}

func (d Dotato) DeleteHistory(h state.History) (err error) {
	if err = d.setState(); err != nil {
		return
	}

	return d.state.DeleteOne(h)
}

// Init ///////////////////////////////////////////////////////////////////////

func (d Dotato) Init() (bool, error) {
	wd, err := os.Getwd()
	if err != nil {
		return false, err
	}

	configPath := filepath.Join(wd, dotatoFileNameConfig)

	// Check if config file exists
	_, err = d.fs.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// do nothing
		} else {
			return false, err
		}
	} else {
		return false, nil
	}

	// Create config file
	cf, err := d.fs.Create(configPath)
	if err != nil {
		return false, err
	}
	defer cf.Close()

	// Write example config
	_, err = cf.Write([]byte(config.GetExample()))
	if err != nil {
		return false, err
	}

	return true, nil
}
