package run

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/msisdev/dotato/pkg/config"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
	"github.com/msisdev/dotato/pkg/log"
	"github.com/msisdev/dotato/pkg/state"
)

type Options struct {
	StatePath string
}

// Each command has different requirement of dependencies.
// This object can provide dependencies on demand.
type Runner struct {
	fs		billy.Filesystem
	log		*log.Logger

	base  *gp.GardenPath		// directory of the config file
	cfg 	*config.Config
	rule	*ignore.RuleTree
	state	*state.State
}

func NewRunner(fs billy.Filesystem, level log.Level) *Runner {
	return &Runner{
		fs: fs,
		log: log.NewLogger(level),
	}
}

func (r *Runner) initState() (err error) {
	if r.state != nil {
		return
	}

	r.state, err = state.NewState(state.PathDefault)
	return
}

// initConfig tries to find config file by
// walking up the directory tree.
// It returns the directory of the config file.
func (r *Runner) initConfig() (ok bool, err error) {
	if r.cfg != nil {
		return
	}

	// Get pwd
	pwd, err := gp.NewGardenPath(".")
	if err != nil {
		return
	}

	config.ReadConfigFileRecur(r.fs, pwd, config.ConfigFileName)

	return false, nil
}

// initIgnore requires the location of the config file.
// If config file exists, it will read only the root
// ignore file and return.
func (r *Runner) initIgnore(path gp.GardenPath) (err error) {
	if r.rule != nil {
		return
	}

	// if base is not set, call initConfig
	if r.base == nil {
		ok, err := r.initConfig()
		if err != nil {
			return err
		}
		if !ok {
			
		}
	}
	

	rules, ok, err := ignore.CompileIgnoreFile(
		osfs.New("/"),
		path,
	)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	
	r.rule = ignore.NewRuleTreeFromPath(path)
	r.rule.Add(path, rules)
	return
}
