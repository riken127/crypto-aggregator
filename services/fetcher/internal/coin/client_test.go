package coin

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Should fetch assets successfully
func TestFetchAssets_ShouldFetchAssetsSuccessfully(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("apiKey") != "testkey" {
			t.Errorf("expected apiKey=testkey, got %s", r.URL.Query().Get("apiKey"))
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"data":[{"id":"bitcoin","symbol":"BTC","name":"Bitcoin","explorer":"https://blockchain.info/","priceUsd":"100","volumeUsd24Hr":"200","changePercent24Hr":"1","marketCapUsd":"1000","vwap24Hr":"99","maxSupply":"21000000","supply":"19000000"}]}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	assets, err := client.FetchAssets()
	if err != nil {
		t.Fatalf("should not error, got %v", err)
	}
	if len(assets) != 1 || assets[0].ID != "bitcoin" {
		t.Fatalf("should fetch bitcoin asset, got %+v", assets)
	}
}

// Should return error on non-200 status
func TestFetchAssets_ShouldReturnErrorOnNon200Status(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	_, err := client.FetchAssets()
	if err == nil {
		t.Fatalf("should return error on non-200 status")
	}
}

// Should return error on invalid JSON
func TestFetchAssets_ShouldReturnErrorOnInvalidJSON(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{invalid json}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	_, err := client.FetchAssets()
	if err == nil {
		t.Fatalf("should return error on invalid JSON")
	}
}

// Should return error on network failure
func TestFetchAssets_ShouldReturnErrorOnNetworkFailure(t *testing.T) {
	// Use an invalid URL to simulate network error
	client := NewClientWithURL("testkey", "http://127.0.0.1:0")
	_, err := client.FetchAssets()
	if err == nil {
		t.Fatalf("should return error on network failure")
	}
}

// Should call correct URL with apiKey
func TestFetchAssets_ShouldCallCorrectURLWithApiKey(t *testing.T) {
	called := false
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.Query().Get("apiKey") != "testkey" {
			t.Errorf("should call with correct apiKey, got %s", r.URL.Query().Get("apiKey"))
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"data":[]}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	_, _ = client.FetchAssets()
	if !called {
		t.Fatalf("should call the API endpoint")
	}
}

// Should handle empty data array
func TestFetchAssets_ShouldHandleEmptyDataArray(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"data":[]}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	assets, err := client.FetchAssets()
	if err != nil {
		t.Fatalf("should not error, got %v", err)
	}
	if len(assets) != 0 {
		t.Fatalf("should return empty slice, got %+v", assets)
	}
}

// Should handle large response bodies
func TestFetchAssets_ShouldHandleLargeResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		body := `{"data":[`
		for i := 0; i < 100; i++ {
			if i > 0 {
				body += ","
			}
			body += fmt.Sprintf(`{"id":"coin%d","symbol":"SYM%d","name":"Coin%d","explorer":"url","priceUsd":"%d","volumeUsd24Hr":"%d","changePercent24Hr":"%d","marketCapUsd":"%d","vwap24Hr":"%d","maxSupply":"%d","supply":"%d"}`, i, i, i, i, i, i, i, i, i, i)
		}
		body += `]}`
		io.WriteString(w, body)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewClientWithURL("testkey", server.URL)
	assets, err := client.FetchAssets()
	if err != nil {
		t.Fatalf("should not error, got %v", err)
	}
	if len(assets) != 100 {
		t.Fatalf("should return 100 assets, got %d", len(assets))
	}
}
