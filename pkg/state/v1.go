package state

import (
	"time"

	"github.com/msisdev/dotato/pkg/config"
)

type V1History struct {
	TargetPath			string			`gorm:"primaryKey"`
	SourcePath			string			`gorm:"uniqueIndex"`
	Mode						config.Mode	`gorm:"not null"`
	TargetUpdatedAt	time.Time 	`gorm:"not null"`
	SourceUpdatedAt	time.Time		`gorm:"not null"`
	Hash						string			`gorm:"not null"`
}

func (d DB) V1_Migrate() error {
	// Create the v1history table
	if err := d.db.AutoMigrate(&V1History{}); err != nil {
		return err
	}
	return nil
}

func (d DB) V1_GetAll() ([]V1History, error) {
	var hs []V1History
	if err := d.db.Find(&hs).Error; err != nil {
		return nil, err
	}
	return hs, nil
}

func (d DB) V1_GetOne(targetPath string) (*V1History, error) {
	h := V1History{ TargetPath: targetPath }
	if err := d.db.First(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

func (d DB) V1_UpsertOne(h V1History) error {
	return d.db.Save(&h).Error
}

func (d DB) V1_DeleteMany(targetPath []string) error {
	var hs = []V1History{}
	for _, path := range targetPath {
		hs = append(hs, V1History{TargetPath: path})
	}

	return d.db.Delete(&hs).Error
}
