package coin

// Asset represents a cryptocurrency asset with its details.
// It includes fields such as ID, symbol, name, explorer URL, price in USD,
// volume in the last 24 hours, change percentage in the last 24 hours,
// market cap in USD, volume-weighted average price in the last 24 hours,
// maximum supply, and current supply.
type Asset struct {
	ID                string `json:"id"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Explorer          string `json:"explorer"`
	PriceUsd          string `json:"priceUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	MarketCapUsd      string `json:"marketCapUsd"`
	Vwap24Hr          string `json:"vwap24Hr"`
	MaxSupply         string `json:"maxSupply"`
	Supply            string `json:"supply"`
}
