package coin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	apiKey string
	apiURL string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		apiURL: "https://rest.coincap.io/v3/assets",
	}
}

// Para testes, permite criar um client customizado
func NewClientWithURL(apiKey, apiURL string) *Client {
	return &Client{
		apiKey: apiKey,
		apiURL: apiURL,
	}
}

func (c *Client) FetchAssets() ([]Asset, error) {
	url := fmt.Sprintf("%s?apiKey=%s", c.apiURL, c.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CoinCap API returned status %d", resp.StatusCode)
	}

	var result struct {
		Data []Asset `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
