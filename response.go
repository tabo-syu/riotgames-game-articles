package riotgamespatchnotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	LOLPatchNoteTagName      = "patch_notes"
	TFTPatchNoteTagName      = "teamfight_tactics_patch_notes"
	ValorantPatchNoteTagName = "patch_notes"
)

type WebsiteResponse interface {
	ToArticles()
}

type URL struct {
	URL string `json:"url"`
}

type ArticleTag struct {
	UID         string `json:"uid"`
	Title       string `json:"title"`
	MachineName string `json:"machine_name"`
	IsHidden    bool   `json:"is_hidden"`
	URL         LOLURL `json:"url"`
}

type LOLURL URL

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

type LOLArticleTag ArticleTag

type LOLArticle struct {
	UID          string          `json:"uid"`
	Title        string          `json:"title"`
	Date         time.Time       `json:"date"`
	Description  string          `json:"description"`
	ArticleType  string          `json:"article_type"`
	URL          LOLURL          `json:"url"`
	ExternalLink string          `json:"external_link"`
	YoutubeLink  string          `json:"youtube_link"`
	Banner       LOLBanner       `json:"banner"`
	Category     []LOLCategory   `json:"category"`
	Author       []LOLAuthor     `json:"author"`
	ArticleTags  []LOLArticleTag `json:"article_tags"`
}

type LOLArticles struct {
	Content []*LOLArticle
}

func (a LOLArticles) FilterByTagName(tagName string) *LOLArticles {
	var patchNotes []*LOLArticle
	for _, article := range a.Content {
		for _, tag := range article.ArticleTags {
			if tag.MachineName == tagName {
				patchNotes = append(patchNotes, article)
			}
		}
	}

	return &LOLArticles{Content: patchNotes}
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

func (r LOLArticlesResponse) ToArticles() *LOLArticles {
	var articles []*LOLArticle
	for _, edge := range r.Result.Data.AllArticles.Edges {
		a := edge.Article
		articles = append(articles, &a)
	}

	return &LOLArticles{Content: articles}
}

type LOLWebsiteResponse struct {
	Res *http.Response
}

func (r LOLWebsiteResponse) Body() (*LOLArticlesResponse, error) {
	body, err := io.ReadAll(r.Res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	var result LOLArticlesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &result, nil
}

type ValorantURL URL

type ValorantArticleTag ArticleTag

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
	ID           string               `json:"id"`
	UID          string               `json:"uid"`
	Title        string               `json:"title"`
	Date         time.Time            `json:"date"`
	Description  string               `json:"description"`
	ArticleType  string               `json:"article_type"`
	ExternalLink string               `json:"external_link"`
	ArticleTags  []ValorantArticleTag `json:"article_tags"`
	Category     []ValorantCategory   `json:"category"`
	PathSlug     string               `json:"pathSlug"`
	URL          ValorantURL          `json:"url"`
	Banner       ValorantBanner       `json:"banner"`
}

type ValorantArticles struct {
	Content []*ValorantArticle
}

func (a ValorantArticles) FilterByTagName(tagName string) *ValorantArticles {
	var articles []*ValorantArticle
	for _, article := range a.Content {
		for _, tag := range article.ArticleTags {
			if tag.MachineName == tagName {
				articles = append(articles, article)
			}
		}
	}

	return &ValorantArticles{Content: articles}
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

func (r ValorantArticlesResponse) ToArticles() *ValorantArticles {
	var articles []*ValorantArticle
	for _, article := range r.Result.Data.AllContentstackArticles.Articles {
		a := article
		articles = append(articles, &a)
	}

	return &ValorantArticles{Content: articles}
}

type ValorantWebsiteResponse struct {
	Res *http.Response
}

func (r ValorantWebsiteResponse) Body() (*ValorantArticlesResponse, error) {
	body, err := io.ReadAll(r.Res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	var result ValorantArticlesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &result, nil
}
