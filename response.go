package riotgamesgamearticles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	lolPatchNoteTagName      = "patch_notes"
	tftPatchNoteTagName      = "teamfight_tactics_patch_notes"
	valorantPatchNoteTagName = "patch_notes"
)

type Article interface {
	GetTags() []ArticleTag
}

type URL struct {
	URL string `json:"url"`
}

type ArticleTag struct {
	UID         string `json:"uid"`
	Title       string `json:"title"`
	MachineName string `json:"machine_name"`
	IsHidden    bool   `json:"is_hidden"`
	URL         URL    `json:"url"`
}

type LOLBanner struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

type LOLCategory struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type LOLAuthor struct {
	Title string `json:"title"`
}

type LOLArticle struct {
	UID          string        `json:"uid"`
	Title        string        `json:"title"`
	Date         time.Time     `json:"date"`
	Description  string        `json:"description"`
	ArticleType  string        `json:"article_type"`
	URL          URL           `json:"url"`
	ExternalLink string        `json:"external_link"`
	YoutubeLink  string        `json:"youtube_link"`
	Banner       LOLBanner     `json:"banner"`
	Category     []LOLCategory `json:"category"`
	Author       []LOLAuthor   `json:"author"`
	ArticleTags  []ArticleTag  `json:"article_tags"`
}

func (a LOLArticle) GetTags() []ArticleTag {
	return a.ArticleTags
}

type ValorantCategory struct {
	Title       string `json:"title"`
	MachineName string `json:"machine_name"`
}

type ValorantBanner struct {
	URL       string `json:"url"`
	Dimension struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"dimension"`
	ContentType string `json:"content_type"`
	FileSize    string `json:"file_size"`
	Filename    string `json:"filename"`
}

type ValorantArticle struct {
	ID           string             `json:"id"`
	UID          string             `json:"uid"`
	Title        string             `json:"title"`
	Date         time.Time          `json:"date"`
	Description  string             `json:"description"`
	ArticleType  string             `json:"article_type"`
	ExternalLink string             `json:"external_link"`
	ArticleTags  []ArticleTag       `json:"article_tags"`
	Category     []ValorantCategory `json:"category"`
	PathSlug     string             `json:"pathSlug"`
	URL          URL                `json:"url"`
	Banner       ValorantBanner     `json:"banner"`
}

func (a ValorantArticle) GetTags() []ArticleTag {
	return a.ArticleTags
}

type Articles[T Article] struct {
	Content []T
}

func (a Articles[T]) FilterByTagName(name string) *Articles[T] {
	var articles []T
	for _, article := range a.Content {
		for _, tag := range article.GetTags() {
			if tag.MachineName == name {
				articles = append(articles, article)
			}
		}
	}

	return &Articles[T]{articles}
}

type ArticlesResponse[T Article] interface {
	ToArticles() *Articles[T]
}

type LOLArticlesResponse struct {
	Result struct {
		Data struct {
			AllArticles struct {
				Edges []struct {
					Article LOLArticle `json:"node"`
				} `json:"edges"`
			} `json:"allArticles"`
		} `json:"data"`
	} `json:"result"`
}

func (r LOLArticlesResponse) ToArticles() *Articles[*LOLArticle] {
	var articles []*LOLArticle
	for _, edge := range r.Result.Data.AllArticles.Edges {
		a := edge.Article
		articles = append(articles, &a)
	}

	return &Articles[*LOLArticle]{articles}
}

type ValorantArticlesResponse struct {
	Result struct {
		Data struct {
			AllContentstackArticles struct {
				Articles []ValorantArticle `json:"nodes"`
			} `json:"allContentstackArticles"`
		} `json:"data"`
	} `json:"result"`
}

func (r ValorantArticlesResponse) ToArticles() *Articles[*ValorantArticle] {
	var articles []*ValorantArticle
	for _, article := range r.Result.Data.AllContentstackArticles.Articles {
		a := article
		articles = append(articles, &a)
	}

	return &Articles[*ValorantArticle]{articles}
}

type WebsiteResponse[T ArticlesResponse[U], U Article] struct {
	Res *http.Response
}

func (r WebsiteResponse[T, U]) Body() (*T, error) {
	body, err := io.ReadAll(r.Res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &result, nil
}
