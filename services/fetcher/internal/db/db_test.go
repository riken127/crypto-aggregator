package db

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Should save asset and record successfully
func TestRepository_ShouldSaveAssetWithRecordSuccessfully(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Asset{}, &AssetRecord{})
	repo := NewRepository(db)

	asset := Asset{
		ID:       "bitcoin",
		Symbol:   "BTC",
		Name:     "Bitcoin",
		Explorer: "https://blockchain.info/",
	}
	record := AssetRecord{
		AssetID:           "bitcoin",
		Timestamp:         time.Now(),
		PriceUsd:          "100",
		VolumeUsd24Hr:     "200",
		ChangePercent24Hr: "1",
		MarketCapUsd:      "1000",
		Vwap24Hr:          "99",
		MaxSupply:         "21000000",
		Supply:            "19000000",
	}
	err := repo.SaveAssetWithRecord(asset, record)
	if err != nil {
		t.Fatalf("should save asset and record, got error: %v", err)
	}

	var got Asset
	if err := db.First(&got, "id = ?", "bitcoin").Error; err != nil {
		t.Fatalf("should find asset, got error: %v", err)
	}
	if got.Symbol != "BTC" {
		t.Errorf("should have symbol BTC, got %s", got.Symbol)
	}

	var gotRecord AssetRecord
	if err := db.First(&gotRecord, "asset_id = ?", "bitcoin").Error; err != nil {
		t.Fatalf("should find asset record, got error: %v", err)
	}
	if gotRecord.PriceUsd != "100" {
		t.Errorf("should have priceUsd 100, got %s", gotRecord.PriceUsd)
	}
}

// Should upsert asset (update on conflict)
func TestRepository_ShouldUpsertAsset(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Asset{})
	repo := NewRepository(db)

	asset := Asset{
		ID:       "bitcoin",
		Symbol:   "BTC",
		Name:     "Bitcoin",
		Explorer: "https://blockchain.info/",
	}
	_ = repo.UpsertAsset(asset)

	// Update asset
	asset.Name = "Bitcoin Updated"
	err := repo.UpsertAsset(asset)
	if err != nil {
		t.Fatalf("should upsert asset, got error: %v", err)
	}

	var got Asset
	if err := db.First(&got, "id = ?", "bitcoin").Error; err != nil {
		t.Fatalf("should find asset, got error: %v", err)
	}
	if got.Name != "Bitcoin Updated" {
		t.Errorf("should update asset name, got %s", got.Name)
	}
}

// Should insert asset record only
func TestRepository_ShouldInsertAssetRecordOnly(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&AssetRecord{})
	repo := NewRepository(db)

	record := AssetRecord{
		AssetID:   "bitcoin",
		Timestamp: time.Now(),
		PriceUsd:  "123",
	}
	err := repo.InsertAssetRecord(record)
	if err != nil {
		t.Fatalf("should insert asset record, got error: %v", err)
	}

	var got AssetRecord
	if err := db.First(&got, "asset_id = ?", "bitcoin").Error; err != nil {
		t.Fatalf("should find asset record, got error: %v", err)
	}
	if got.PriceUsd != "123" {
		t.Errorf("should have priceUsd 123, got %s", got.PriceUsd)
	}
}
