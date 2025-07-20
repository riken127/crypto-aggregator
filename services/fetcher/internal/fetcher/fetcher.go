package fetcher

import (
	"time"

	"github.com/riken127/crypto-aggregator-fetcher/internal/coin"
	"github.com/riken127/crypto-aggregator-fetcher/internal/db"
	"github.com/riken127/crypto-aggregator-fetcher/internal/temporal"
)

type Fetcher struct {
	coinClient *coin.Client
	repo       *db.Repository
}

// NewFetcher creates a new Fetcher instance with the provided CoinCap API client and database repository.
func NewFetcher(coinClient *coin.Client, repo *db.Repository) *Fetcher {
	return &Fetcher{coinClient, repo}
}

// FetchAndStore retrieves cryptocurrency assets from the CoinCap API and stores them in the database.
func (f *Fetcher) FetchAndStore() error {
	assets, err := f.coinClient.FetchAssets()
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	for _, a := range assets {
		asset := db.Asset{
			ID:       a.ID,
			Symbol:   a.Symbol,
			Name:     a.Name,
			Explorer: a.Explorer,
		}
		record := db.AssetRecord{
			AssetID:           a.ID,
			Timestamp:         now,
			PriceUsd:          a.PriceUsd,
			VolumeUsd24Hr:     a.VolumeUsd24Hr,
			ChangePercent24Hr: a.ChangePercent24Hr,
			MarketCapUsd:      a.MarketCapUsd,
			Vwap24Hr:          a.Vwap24Hr,
			MaxSupply:         a.MaxSupply,
			Supply:            a.Supply,
		}

		if err := f.repo.SaveAssetWithRecord(asset, record); err != nil {
			return err
		}
	}

	var temporalAssets []temporal.Asset

	for _, a := range assets {
		temporalAsset := temporal.Asset{
			ID:                a.ID,
			Symbol:            a.Symbol,
			Name:              a.Name,
			Explorer:          a.Explorer,
			PriceUsd:          a.PriceUsd,
			VolumeUsd24Hr:     a.VolumeUsd24Hr,
			ChangePercent24Hr: a.ChangePercent24Hr,
			MarketCapUsd:      a.MarketCapUsd,
			Vwap24Hr:          a.Vwap24Hr,
			MaxSupply:         a.MaxSupply,
			Supply:            a.Supply,
		}
		temporalAssets = append(temporalAssets, temporalAsset)
	}

	temporal.StartAggregatorWorkflow(temporalAssets)

	return nil
}
