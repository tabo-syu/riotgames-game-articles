package riotgamespatchnotes

import (
	"context"
	"fmt"
)

type LOLWebsiteArticles struct {
	Articles *LOLArticles
}

func NewLOLWebsiteArticles(locale string) (*LOLWebsiteArticles, error) {
	req, err := NewLOLWebsiteRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewLOLWebsiteArticles: %w", err)
	}

	res, err := NewFetcher().FetchLOL(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	result, err := res.Body()
	if err != nil {
		return nil, fmt.Errorf("failed to parse to ArticlesResponse: %w", err)
	}

	return &LOLWebsiteArticles{result.ToArticles()}, err
}

func (l LOLWebsiteArticles) All() []*LOLArticle {
	return l.Articles.Content
}

func (l LOLWebsiteArticles) LOLPatchNotes() []*LOLArticle {
	return l.Articles.FilterByTagName(LOLPatchNoteTagName).Content
}

func (l LOLWebsiteArticles) TFTPatchNotes() []*LOLArticle {
	return l.Articles.FilterByTagName(TFTPatchNoteTagName).Content
}

type ValorantWebsiteArticles struct {
	Articles *ValorantArticles
}

func NewValorantWebsiteArticles(locale string) (*ValorantWebsiteArticles, error) {
	req, err := NewValorantWebsiteRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewValorantWebsiteArticles: %w", err)
	}

	res, err := NewFetcher().FetchValorant(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	result, err := res.Body()
	if err != nil {
		return nil, fmt.Errorf("failed to parse to ArticlesResponse: %w", err)
	}

	return &ValorantWebsiteArticles{result.ToArticles()}, err
}

func (l ValorantWebsiteArticles) All() []*ValorantArticle {
	return l.Articles.Content
}

func (l ValorantWebsiteArticles) PatchNotes() []*ValorantArticle {
	return l.Articles.FilterByTagName(ValorantPatchNoteTagName).Content
}
