package state

import (
	"github.com/msisdev/dotato/pkg/config"
)

func (d DB) v1_migrate() error {
	if err := d.db.AutoMigrate(&History{}); err != nil {
		return err
	}
	return nil
}

func (d DB) v1_getAll() ([]History, error) {
	var hs []History
	if err := d.db.Find(&hs).Error; err != nil {
		return nil, err
	}
	return hs, nil
}
func (d DB) v1_getAllByMode(mode config.Mode) ([]History, error) {
	var hs []History
	if err := d.db.Where("mode = ?", mode).Find(&hs).Error; err != nil {
		return nil, err
	}
	return hs, nil
}
func (d DB) v1_getOne(targetPath string) (*History, error) {
	h := History{ TargetPath: targetPath }
	if err := d.db.First(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

func (d DB) v1_upsertOne(h History) error {
	return d.db.Save(&h).Error
}

func (d DB) v1_deleteMany(targetPath []string) error {
	var hs = []History{}
	for _, path := range targetPath {
		hs = append(hs, History{TargetPath: path})
	}

	return d.db.Delete(&hs).Error
}
