package serverapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	baseURL    string
	Token      string
	Party      []Pokemon
}

func NewClient(timeout time.Duration, baseURL string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		Party:   []Pokemon{},
	}
}
