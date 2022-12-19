package riotgamespatchnotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type URL struct {
	URL string `json:"url"`
}

type Banner struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

type Category struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type Author struct {
	Title string `json:"title"`
}

type ArticleTag struct {
	UID         string `json:"uid"`
	Title       string `json:"title"`
	MachineName string `json:"machine_name"`
	IsHidden    bool   `json:"is_hidden"`
	URL         URL    `json:"url"`
}

type Article struct {
	UID          string       `json:"uid"`
	Title        string       `json:"title"`
	Date         time.Time    `json:"date"`
	Description  string       `json:"description"`
	ArticleType  string       `json:"article_type"`
	URL          URL          `json:"url"`
	ExternalLink string       `json:"external_link"`
	YoutubeLink  string       `json:"youtube_link"`
	Banner       Banner       `json:"banner"`
	Category     []Category   `json:"category"`
	Author       []Author     `json:"author"`
	ArticleTags  []ArticleTag `json:"article_tags"`
}

const (
	LOLPatchNoteTagName = "patch_notes"
	TFTPatchNoteTagName = "teamfight_tactics_patch_notes"
)

type Articles struct {
	Content []*Article
}

func (a Articles) FilterByTagName(tagName string) *Articles {
	var patchNotes []*Article
	for _, article := range a.Content {
		for _, tag := range article.ArticleTags {
			if tag.MachineName == tagName {
				patchNotes = append(patchNotes, article)
			}
		}
	}

	return &Articles{Content: patchNotes}
}

type ArticlesResponse struct {
	Result struct {
		Data struct {
			AllArticles struct {
				Edges []struct {
					Article Article `json:"node"`
				} `json:"edges"`
			} `json:"allArticles"`
		} `json:"data"`
	} `json:"result"`
}

func (g ArticlesResponse) ToArticles() *Articles {
	var articles []*Article
	for _, edge := range g.Result.Data.AllArticles.Edges {
		a := edge.Article
		articles = append(articles, &a)
	}

	return &Articles{Content: articles}
}

type WebsiteResponse struct {
	Res *http.Response
}

func (r WebsiteResponse) Body() (*ArticlesResponse, error) {
	body, err := io.ReadAll(r.Res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	var result ArticlesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &result, nil
}
