package riotgamesgamearticles

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type fetcher[T ArticlesResponse[U], U Article] struct {
	client *http.Client
}

func newFetcher[T ArticlesResponse[U], U Article]() *fetcher[T, U] {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 2 * time.Second

	return &fetcher[T, U]{client: client.StandardClient()}
}

func (f fetcher[T, U]) Fetch(req *websiteRequest) (*WebsiteResponse[T, U], error) {
	res, err := f.client.Do(req.Req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resources: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code was not 200: %w", err)
	}

	return &WebsiteResponse[T, U]{Res: res}, nil
}
