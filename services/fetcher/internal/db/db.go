package db

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpsertAsset(asset Asset) error {
	return r.db.Save(&asset).Error
}
func (r *Repository) InsertAssetRecord(asset AssetRecord) error {
	return r.db.Create(&asset).Error
}
func (r *Repository) SaveAssetWithRecord(asset Asset, record AssetRecord) error {
	tx := r.db.Begin()
	if err := tx.Save(&asset).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&record).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
