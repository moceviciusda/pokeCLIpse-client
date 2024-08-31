package serverapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	baseURL    string
}

func NewClient(timeout time.Duration, baseURL string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}
