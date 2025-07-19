package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/seuusuario/crypto-aggregator-fetcher/internal/coin"
	"github.com/seuusuario/crypto-aggregator-fetcher/internal/db"
	"github.com/seuusuario/crypto-aggregator-fetcher/internal/fetcher"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// run initializes the database, fetches data from the CoinCap API, and starts the HTTP server.
func run(apiKey, dsn string) error {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := database.AutoMigrate(&db.Asset{}, &db.AssetRecord{}); err != nil {
		return err
	}

	coinClient := coin.NewClient(apiKey)
	repo := db.NewRepository(database)
	f := fetcher.NewFetcher(coinClient, repo)

	if err := f.FetchAndStore(); err != nil {
		return err
	}

	r := gin.Default()
	r.GET("healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	return r.Run(":8080")
}

// main is the entry point of the application.
// It reads the API key and database connection string from environment variables,
// initializes the database, fetches data from the CoinCap API, and starts the HTTP server.
func main() {
	apiKey := os.Getenv("COINCAP_API_KEY")
	postgresDSN := os.Getenv("POSTGRES_DSN")

	if apiKey == "" || postgresDSN == "" {
		panic("COINCAP_API_KEY and POSTGRES_DSN environment variables must be set")
	}

	if err := run(apiKey, postgresDSN); err != nil {
		panic(err)
	}
}
