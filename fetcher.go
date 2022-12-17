package riotgamespatchnotes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Fetcher struct {
	client *http.Client
}

func NewFetcher() *Fetcher {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 2 * time.Second

	return &Fetcher{client: client.StandardClient()}
}

func (f Fetcher) Fetch(req *GameUpdatesRequest) (*GameUpdatesResponse, error) {
	res, err := f.client.Do(req.Req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resources: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code was not 200: %w", err)
	}

	return &GameUpdatesResponse{Res: res}, nil
}
