package riotgamespatchnotes

import (
	"context"
	"fmt"
)

type LOLArticles struct {
	Articles *Articles
}

func NewLOLArticles(locale string) (*LOLArticles, error) {
	req, err := NewLOLWebsiteRequest(context.Background(), locale)
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

	return &LOLArticles{result.ToArticles()}, err
}

func (l LOLArticles) All() []*Article {
	return l.Articles.Content
}

func (l LOLArticles) LOLPatchNotes() []*Article {
	return l.Articles.FilterByTagName(LOLPatchNoteTagName).Content
}

func (l LOLArticles) TFTPatchNotes() []*Article {
	return l.Articles.FilterByTagName(TFTPatchNoteTagName).Content
}
