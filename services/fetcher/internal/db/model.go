package db

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	ID       string `gorm:"uniqueIndex"`
	Symbol   string
	Name     string
	Explorer string
	Records  []AssetRecord `gorm:"foreignKey:AssetID"`
}
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
