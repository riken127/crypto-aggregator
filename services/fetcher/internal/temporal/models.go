package temporal

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

type AssetsInput struct {
	Assets []Asset `json:"assets"`
}
