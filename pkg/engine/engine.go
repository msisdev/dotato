package engine

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/msisdev/dotato/internal/config"
	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/ignore"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/internal/state"
	"gorm.io/gorm"
)

type Engine struct {
	fs      billy.Filesystem
	isMem   bool
	maxIter int

	cdir 	gp.GardenPath
	cfg  	*config.Config
	ig   	*ignore.Ignore
	state	*state.State
}

func New() *Engine {
	return NewWithFS(filesystem.NewOSFS(), false)
}

func NewMemfs() *Engine {
	return NewWithFS(memfs.New(), true)
}

func NewWithFS(fs billy.Filesystem, isMem bool) *Engine {
	return &Engine{
		fs:      fs,
		isMem:   isMem,
		maxIter: factory.DotatoMaxFSIter,
	}
}

// Config /////////////////////////////////////////////////////////////////////

func (e Engine) GetConfigDir() (gp.GardenPath, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	return e.cdir, nil
}

func (e Engine) GetConfigVersion() (string, error) {
	if err := e.readConfig(); err != nil {
		return "", err
	}

	return e.cfg.Version, nil
}

func (e Engine) GetConfigMode() (string, error) {
	if err := e.readConfig(); err != nil {
		return "", err
	}

	return e.cfg.Mode, nil
}

func (e Engine) GetConfigPlans() (map[string][]string, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	return e.cfg.Plans, nil
}

func (e Engine) GetConfigGroups(plan string) (map[string]bool, bool, error) {
	if err := e.readConfig(); err != nil {
		return nil, false, err
	}

	// Get groups
	groupList, ok := e.cfg.Plans[plan]
	if !ok {
		return nil, false, nil
	}

	// Convert to map[string]bool
	groupSet := make(map[string]bool)
	for _, group := range groupList {
		groupSet[group] = true
	}

	return groupSet, true, nil
}

func (e Engine) GetConfigGroupAll() (map[string]bool, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	// Get groups
	groups := make(map[string]bool)
	for group := range e.cfg.Groups {
		groups[group] = true
	}

	return groups, nil
}

func (e Engine) GetConfigGroupBase(group, resolver string) (gp.GardenPath, []string, error) {
	if err := e.readConfig(); err != nil {
		return nil, nil, err
	}

	// Get group base
	return e.cfg.GetGroupBase(group, resolver)
}

func (e Engine) GetConfigGroupResolvers(group string) (map[string]string, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	rs := make(map[string]string)

	// For each resolver
	for name, resolver := range e.cfg.Groups[group] {
		// Collect resolver
		rs[name] = resolver
	}

	return rs, nil
}

// Ignore /////////////////////////////////////////////////////////////////////

func (e Engine) ReadGroupIgnore(group string) (*ignore.Ignore, error) {
	if err := e.readConfig(); err != nil { return nil, err }

	dir := e.cdir.Copy()
	dir = append(dir, group)
	return factory.ReadIgnoreRecur(e.fs, dir)
}

// State //////////////////////////////////////////////////////////////////////

func (e Engine) GetHistoryByMode(mode string) ([]History, error) {
	if err := e.readState(); err != nil {
		return nil, err
	}

	return e.state.GetAllByMode(mode)
}

func (e Engine) UpsertHistory(h History) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.Upsert(h)
}

func (e Engine) DeleteHistory(h History) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.Delete(h)
}

func (e *Engine) StateTx(fn func(tx *gorm.DB) error) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.Tx(fn)
}

func (e *Engine) StateTxSafe(fn func(tx *gorm.DB) error) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.TxSafe(fn)
}

func (e Engine) StateTxUpsert(tx *gorm.DB, h History) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.TxUpsert(tx, h)
}

func (e Engine) StateTxDelete(tx *gorm.DB, h History) error {
	if err := e.readState(); err != nil {
		return err
	}

	return e.state.TxDelete(tx, h)
}
