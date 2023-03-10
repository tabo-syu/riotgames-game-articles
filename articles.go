package riotgamesgamearticles

import (
	"context"
	"fmt"
)

type LOLWebsiteArticles struct {
	Articles *Articles[*LOLArticle]
}

func NewLOLWebsiteArticles(locale string) (*LOLWebsiteArticles, error) {
	req, err := newLOLWebsiteRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewLOLWebsiteArticles: %w", err)
	}

	res, err := newFetcher[LOLArticlesResponse, *LOLArticle]().Fetch(req)
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
	return l.Articles.FilterByTagName(lolPatchNoteTagName).Content
}

func (l LOLWebsiteArticles) TFTPatchNotes() []*LOLArticle {
	return l.Articles.FilterByTagName(tftPatchNoteTagName).Content
}

type ValorantWebsiteArticles struct {
	Articles *Articles[*ValorantArticle]
}

func NewValorantWebsiteArticles(locale string) (*ValorantWebsiteArticles, error) {
	req, err := newValorantWebsiteRequest(context.Background(), locale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewValorantWebsiteArticles: %w", err)
	}

	res, err := newFetcher[ValorantArticlesResponse, *ValorantArticle]().Fetch(req)
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
	return l.Articles.FilterByTagName(valorantPatchNoteTagName).Content
}
