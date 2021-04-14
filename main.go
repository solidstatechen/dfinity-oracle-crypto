package main

import (
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func main() {
	config := models.Config{
		CanisterName:   "crypto_oracle",
		UpdateInterval: 1 * time.Minute,
	}

	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{
				Key: "bitcoin",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true&include_last_updated_at=true",
						JSONPaths: map[string]string{
							"usd_per_token": "$.bitcoin.usd",
							"market_cap": "$.bitcoin.usd_market_cap",
							"24h_vol": "$.bitcoin.usd_24h_vol",
							"24h_perc_change": "$.bitcoin.usd_24h_change",
							"last_updated": "$.bitcoin.last_updated_at",
						},
					},
				},
			},
			{
				Key: "ethereum",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true&include_last_updated_at=true",
						JSONPaths: map[string]string{
							"usd_per_token": "$.ethereum.usd",
							"market_cap": "$.ethereum.usd_market_cap",
							"24h_vol": "$.ethereum.usd_24h_vol",
							"24h_perc_change": "$.ethereum.usd_24h_change",
							"last_updated": "$.ethereum.last_updated_at",
						},
					},
				},
			},
			{
				Key: "ampleforth",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=ampleforth&vs_currencies=usd&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true&include_last_updated_at=true",
						JSONPaths: map[string]string{
							"usd_per_token": "$.ampleforth.usd",
							"market_cap": "$.ampleforth.usd_market_cap",
							"24h_vol": "$.ampleforth.usd_24h_vol",
							"24h_perc_change": "$.ampleforth.usd_24h_change",
							"last_updated": "$.ampleforth.last_updated_at",
						},
					},
				},
			},
		},
	}

	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}
