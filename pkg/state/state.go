package state

import (
	"fmt"

	"github.com/msisdev/dotato/pkg/config"
	"gorm.io/gorm"
)

var (
	ErrVersionUnknown = fmt.Errorf("unknown version")
)

type State struct {
	DB *gorm.DB
}

// Latest version of history schema
type History = HistoryV1

// What New does:
//  - Initialize db instance
//  - Migrate db to latest version
func New(statePath string) (*State, error) {
	// Open db
	db, ver, err := NewDB(statePath)
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

func (s State) GetAllLink() ([]History, error) {
	return s.v1_getAllByMode(config.ModeLink)
}

func (s State) UpsertOne(h History) error {
	return s.v1_upsertOne(h)
}
