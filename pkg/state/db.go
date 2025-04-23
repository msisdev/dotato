package state

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/msisdev/dotato/pkg/config"
	"gorm.io/gorm"
)

const (
	DBPathDefault	= "~/.local/share/dotato/dotatostate.sqlite"
	DBPathInMemory = ":memory:"
	
	KeyVersion	 	= "version"
)

// DB is a wrapper of external db driver
type DB struct {
	db *gorm.DB
}

type Store struct {
	Key		string	`gorm:"primaryKey"`
	Value string
}

func NewDB(path string) (*DB, error) {
	var d DB
	{
		// Open db
		db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	
		// Create store table
		if err := db.AutoMigrate(&Store{}); err != nil {
			return nil, err
		}

		d.db = db
	}
	
	// Get version
	ver, ok, err := d.GetVersion()
	if err != nil {
		return nil, err
	}
	
	// Is db empty?
	if !ok {
		// Set version
		d.SetVersion(config.GetDotatoVersion())

		// Migrate to v1
		if err := d.V1_Migrate(); err != nil {
			return nil, err
		}
	}
	
	if ver != config.GetDotatoVersion() {
		// Migrate between different versions
	}

	return &d, nil
}

func (d DB) GetVersion() (string, bool, error) {
	store := Store{ Key: KeyVersion }
	if err := d.db.First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", false, nil
		}
		return "", false, err
	}
	return store.Value, true, nil
}

func (d DB) SetVersion(version string) error {
	store := Store{
		Key:   KeyVersion,
		Value: version,
	}
	if err := d.db.Save(&store).Error; err != nil {
		return err
	}
	return nil
}

