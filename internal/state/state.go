package state

import (
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"gorm.io/gorm"
)

type State struct {
	DB *gorm.DB
}

// Latest version of history schema
type History = HistoryV1

// What New does:
//   - Initialize db instance
//   - Migrate db to latest version
func New(fs billy.Filesystem, statePath string) (*State, error) {
	// Create directories
	err := fs.MkdirAll(filepath.Dir(statePath), 0644)
	if err != nil {
		return nil, err
	}

	// Open db
	db, ver, err := newDB(statePath)
	if err != nil {
		return nil, err
	}

	// Migrate
	switch ver {
	case Version1:
		// Migrate V1
		if err := v1_migrate(db); err != nil {
			return nil, err
		}

	default:
		return nil, ErrVersionUnknown
	}

	return &State{DB: db}, nil
}

///////////////////////////////////////////////////////////

func (s State) GetAllByMode(mode string) ([]History, error) {
	return s.v1_getAllByMode(mode)
}

func (s State) Upsert(h History) error {
	return s.v1_upsert(h)
}

func (s State) Delete(h History) error {
	return s.v1_delete(h)
}

// Doc at https://gorm.io/docs/transactions.html#Transaction
func (s State) Tx(fn func(tx *gorm.DB) error) error {
	return s.v1_tx(fn)
}

// Intercept error from fn, run tx safely and return the error
func (s State) TxSafe(fn func(tx *gorm.DB) error) error {
	var fnErr error

	// Must return nil
	s.v1_tx(func(tx *gorm.DB) error {
		fnErr = fn(tx)

		return nil
	})

	return fnErr
}

func (s State) TxUpsert(tx *gorm.DB, h History) error {
	return s.v1_tx_upsert(tx, h)
}

func (s State) TxDelete(tx *gorm.DB, h History) error {
	return s.v1_tx_delete(tx, h)
}
