package state

import (
	"errors"
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
func NewDB(path string) (*gorm.DB, Version, error) {
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
	_, ok, err := GetVersion(db)
	if err != nil {
		return nil, VersionUnknown, err
	}

	// Is this db new?
	if !ok {
		// Set version
		if err := SetVersion(db, Version1); err != nil {
			return nil, VersionUnknown, err
		}
	}

	return db, Version1, nil
}

// Select version from db
func GetVersion(db *gorm.DB) (Version, bool, error) {
	// Query version
	store := Store{ Key: KeyVersion }
	if err := db.First(&store).Error; err != nil {
		// not found error?
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", false, nil
		}

		// It is an error
		return "", false, err
	}

	// Map value to version
	switch store.Value {
	case string(Version1):
		return Version1, true, nil
	default:
		return VersionUnknown, true, nil
	}
}

// Upsert version into db
func SetVersion(db *gorm.DB, version Version) error {
	store := Store{
		Key:   KeyVersion,
		Value: string(version),
	}
	if err := db.Save(&store).Error; err != nil {
		return err
	}
	return nil
}
