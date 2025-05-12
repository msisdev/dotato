package state

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HistoryV1 struct {
	DotPath		string		`gorm:"primaryKey"`		// file path that is in use for system
	DttPath		string		`gorm:"not null"`	// file path that is in your dotato repository
	Mode			string		`gorm:"not null"`
	CreatedAt	time.Time	`gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime"`
}

func v1_migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&HistoryV1{}); err != nil {
		return err
	}
	return nil
}

func (s State) v1_upsert(h HistoryV1) error {
	return s.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "dot_path"}},
		DoUpdates: clause.AssignmentColumns([]string{"dtt_path", "mode", "updated_at"}),
	}).Create(&h).Error
}

func (s State) v1_getAllByMode(mode string) (hs []HistoryV1, err error) {
	err = s.DB.Where("mode = ?", mode).Find(&hs).Error
	
	return
}

func (s State) v1_delete(h HistoryV1) error {
	return s.DB.Delete(&h).Error
}

func (s State) v1_tx(fn func(tx *gorm.DB) error) error {
	return s.DB.Transaction(fn)
}

func (s State) v1_tx_upsert(tx *gorm.DB, h HistoryV1) error {
	return tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "dot_path"}},
		DoUpdates: clause.AssignmentColumns([]string{"dtt_path", "mode", "updated_at"}),
	}).Create(&h).Error
}

func (s State) v1_tx_delete(tx *gorm.DB, h HistoryV1) error {
	return tx.Delete(&h).Error
}
