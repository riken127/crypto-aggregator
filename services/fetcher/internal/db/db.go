package db

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of Repository with the provided gorm.DB connection.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// UpsertAsset inserts or updates an asset in the database.
func (r *Repository) UpsertAsset(asset Asset) error {
	return r.db.Save(&asset).Error
}

// InsertAssetRecord inserts a new AssetRecord into the database.
func (r *Repository) InsertAssetRecord(asset AssetRecord) error {
	return r.db.Create(&asset).Error
}

// SaveAssetWithRecord inserts or updates an asset and creates a new AssetRecord in a transaction.
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
