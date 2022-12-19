package riotgamespatchnotes

import (
	"context"
	"fmt"
)

type LOLWebsiteArticles struct {
	Articles *Articles
}

func NewLOLWebsiteArticles(locale string) (*LOLWebsiteArticles, error) {
	req, err := NewLOLWebsiteRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewLOLWebsiteArticles: %w", err)
	}

	res, err := NewFetcher().Fetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	result, err := res.Body()
	if err != nil {
		return nil, fmt.Errorf("failed to parse to ArticlesResponse: %w", err)
	}

	return &LOLWebsiteArticles{result.ToArticles()}, err
}

func (l LOLWebsiteArticles) All() []*Article {
	return l.Articles.Content
}

func (l LOLWebsiteArticles) LOLPatchNotes() []*Article {
	return l.Articles.FilterByTagName(LOLPatchNoteTagName).Content
}

func (l LOLWebsiteArticles) TFTPatchNotes() []*Article {
	return l.Articles.FilterByTagName(TFTPatchNoteTagName).Content
}
