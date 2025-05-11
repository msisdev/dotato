package state

import (
	// "errors"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Key value store
type Store struct {
	Key		string	`gorm:"primaryKey"`
	Value string
}

// DB schema version
type Version string
const (
	KeyVersion							= "version"
	VersionUnknown	Version = "unknown"
	Version1 				Version = "v1"
	DBVersion	Version = Version1
)

const (
	PathInMemory = ":memory:"
)

var (
	ErrVersionUnknown = fmt.Errorf("unknown version")

)

// What NewDB does:
//  - Open/create db file
//  - Migrate Store table
//
// What NewDB does not do:
//  - Migrate History table
func newDB(path string) (*gorm.DB, Version, error) {
	// Open dbt
	config := gorm.Config{
		// Logger: nil,
	}
	db, err := gorm.Open(sqlite.Open(path), &config)
	if err != nil {
		return nil, VersionUnknown, err
	}

	// Create store table
	if err := db.AutoMigrate(&Store{}); err != nil {
		return nil, VersionUnknown, err
	}

	// Get version
	_, ok, err := getVersion(db)
	if err != nil {
		return nil, VersionUnknown, err
	}

	// Is version not found?
	if !ok {
		// This db is new.
		// Set version to default.
		if err := setVersion(db, DBVersion); err != nil {
			return nil, VersionUnknown, err
		}
	}

	return db, DBVersion, nil
}

// Select version from db
func getVersion(db *gorm.DB) (Version, bool, error) {
	// Query rows that their keys match 'version'
	var stores []Store
	if err := db.Where("key = ?", KeyVersion).Find(&stores).Error; err != nil {
		return "", false, err
	}

	// No rows?
	if len(stores) == 0 {
		return "", false, nil
	}

	// Too many rows?
	if len(stores) > 1 {
		return "", false, fmt.Errorf("multiple version rows found")
	}
	
	// Select first row
	store := stores[0]

	// Type cast
	switch store.Value {
	case string(DBVersion):
		return DBVersion, true, nil
	default:
		return VersionUnknown, true, nil
	}
}

// Upsert version into db
func setVersion(db *gorm.DB, version Version) error {
	store := Store{
		Key:   KeyVersion,
		Value: string(version),
	}
	if err := db.Save(&store).Error; err != nil {
		return err
	}
	return nil
}
