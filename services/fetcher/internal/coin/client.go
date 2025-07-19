package coin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Asset represents a cryptocurrency asset with its details.
type Client struct {
	apiKey string
	apiURL string
}

// NewClient creates a new CoinCap API client with the provided API key.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		apiURL: "https://rest.coincap.io/v3/assets",
	}
}

// NewClientWithURL creates a new CoinCap API client with the provided API key and custom URL (testing purposes).
func NewClientWithURL(apiKey, apiURL string) *Client {
	return &Client{
		apiKey: apiKey,
		apiURL: apiURL,
	}
}

// FetchAssets retrieves the list of cryptocurrency assets from the CoinCap API.
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
