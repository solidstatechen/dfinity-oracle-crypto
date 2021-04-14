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
				Key: "bitcoin_price",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd",
						JSONPaths: map[string]string{"usd_per_token": "$.bitcoin.usd"},
					},
				},
			},
			{
				Key: "ethereum_price",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd",
						JSONPaths: map[string]string{"usd_per_token": "$.ethereum.usd"},
					},
				},
			},
			{
				Key: "ampleforth_price",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=ampleforth&vs_currencies=usd",
						JSONPaths: map[string]string{"usd_per_token": "$.ampleforth.usd"},
					},
				},
			},
		},
	}

	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}
