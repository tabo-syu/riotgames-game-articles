package riotgamespatchnotes

import (
	"context"
	"fmt"
)

func LOL(ctx context.Context, locale string) ([]*Article, error) {
	req, err := NewLOLGameUpdatesRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize LOLGameUpdatesRequest: %w", err)
	}

	res, err := NewFetcher().Fetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	result, err := res.Body()
	if err != nil {
		return nil, fmt.Errorf("failed to parse to GameUpdates: %w", err)
	}

	articles := result.ToArticles()

	return articles.FilterByTagName(LOLPatchNoteTagName).Content, nil
}

func TFT(ctx context.Context, locale string) ([]*Article, error) {
	req, err := NewLOLGameUpdatesRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize LOLGameUpdatesRequest: %w", err)
	}

	res, err := NewFetcher().Fetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	result, err := res.Body()
	if err != nil {
		return nil, fmt.Errorf("failed to parse to GameUpdates: %w", err)
	}

	articles := result.ToArticles()

	return articles.FilterByTagName(TFTPatchNoteTagName).Content, nil
}
