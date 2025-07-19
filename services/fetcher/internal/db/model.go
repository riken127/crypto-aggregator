package db

import (
	"time"

	"gorm.io/gorm"
)

// Asset represents a cryptocurrency asset with its details.
type Asset struct {
	gorm.Model
	ID       string `gorm:"uniqueIndex"`
	Symbol   string
	Name     string
	Explorer string
	Records  []AssetRecord `gorm:"foreignKey:AssetID"`
}

// AssetRecord represents a record of an asset's data at a specific timestamp.
// It includes fields such as asset ID, timestamp, price in USD,
// volume in the last 24 hours, change percentage in the last 24 hours,
// market cap in USD, volume-weighted average price in the last 24 hours,
// maximum supply, and current supply.
type AssetRecord struct {
	ID                uint      `gorm:"primaryKey"`
	AssetID           string    `gorm:"index;not null"`
	Timestamp         time.Time `gorm:"index"`
	PriceUsd          string
	VolumeUsd24Hr     string
	ChangePercent24Hr string
	MarketCapUsd      string
	Vwap24Hr          string
	MaxSupply         string
	Supply            string
}
