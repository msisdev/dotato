package state

import (
	"time"

	"github.com/msisdev/dotato/pkg/config"
	"gorm.io/gorm"
)

type HistoryV1 struct {
	TargetPath			string			`gorm:"primaryKey"`
	SourcePath			string			`gorm:"uniqueIndex"`
	Mode						config.Mode	`gorm:"not null"`
	TargetUpdatedAt	time.Time 	`gorm:"not null"`
	SourceUpdatedAt	time.Time		`gorm:"not null"`
	Hash						string			`gorm:"not null"`
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

func (s State) v1_getAllByMode(mode config.Mode) (hs []HistoryV1, err error) {
	err = s.DB.Where("mode = ?", mode).Find(&hs).Error
	
	return
}
