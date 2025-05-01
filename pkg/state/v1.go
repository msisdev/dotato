package state

import (
	"time"

	"gorm.io/gorm"
)

type HistoryV1 struct {
	RemotePath			string			`gorm:"primaryKey"`
	LocalPath				string			`gorm:"uniqueIndex"`
	Mode						string			`gorm:"not null"`
	CreatedAt				time.Time		`gorm:"autoCreateTime"`
	UpdatedAt				time.Time		`gorm:"autoUpdateTime"`
}

func v1_migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&HistoryV1{}); err != nil {
		return err
	}
	return nil
}

func (s State) v1_upsertOne(h HistoryV1) error {
	return s.DB.Save(&h).Error
}

func (s State) v1_getAllByMode(mode string) (hs []HistoryV1, err error) {
	err = s.DB.Where("mode = ?", mode).Find(&hs).Error
	
	return
}
