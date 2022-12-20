package riotgamespatchnotes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Fetcher[T ArticlesResponse[U], U Article] struct {
	client *http.Client
}

func NewFetcher[T ArticlesResponse[U], U Article]() *Fetcher[T, U] {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 2 * time.Second

	return &Fetcher[T, U]{client: client.StandardClient()}
}

func (f Fetcher[T, U]) Fetch(req *WebsiteRequest) (*WebsiteResponse[T, U], error) {
	res, err := f.client.Do(req.Req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resources: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code was not 200: %w", err)
	}

	return &WebsiteResponse[T, U]{Res: res}, nil
}
