package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}

	apiBaseURL := os.Getenv("SERVER_API_BASE_URL")
	if apiBaseURL == "" {
		panic("SERVER_API_BASE_URL not found in .env file")
	}

	apiClient := serverapi.NewClient(time.Second*5, apiBaseURL)
	cfg := &config{
		apiClient: apiClient,
	}

	startCLI(cfg)
}
