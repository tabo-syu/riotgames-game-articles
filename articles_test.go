package riotgamespatchnotes

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewLOLWebsiteArticles(t *testing.T) {
	for _, locale := range LOLLocales {
		t.Run(fmt.Sprintf("locale: %s", locale), func(t *testing.T) {
			articles, err := NewLOLWebsiteArticles(locale)
			if err != nil {
				t.Errorf("NewLOLWebsiteArticles() error = %v", err)
			}
			t.Logf(articles.LOLPatchNotes()[0].Title)

			time.Sleep(1 * time.Second)
		})
	}
}

func TestLOLWebsiteArticles_All(t *testing.T) {
	type fields struct {
		Articles *Articles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Article
	}{
		{
			"Test all articles",
			fields{
				&Articles{
					[]*Article{
						{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []ArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []ArticleTag{}},
					},
				},
			},
			[]*Article{
				{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title2", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}}},
				{Title: "title3", ArticleTags: []ArticleTag{{MachineName: TFTPatchNoteTagName}}},
				{Title: "title4", ArticleTags: []ArticleTag{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LOLWebsiteArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLWebsiteArticles.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOLWebsiteArticles_LOLPatchNotes(t *testing.T) {
	type fields struct {
		Articles *Articles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Article
	}{
		{
			"Test all articles",
			fields{
				&Articles{
					[]*Article{
						{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []ArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []ArticleTag{}},
					},
				},
			},
			[]*Article{
				{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title2", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LOLWebsiteArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.LOLPatchNotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLWebsiteArticles.LOLPatchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOLWebsiteArticles_TFTPatchNotes(t *testing.T) {
	type fields struct {
		Articles *Articles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Article
	}{
		{
			"Test all articles",
			fields{
				&Articles{
					[]*Article{
						{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []ArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []ArticleTag{}},
					},
				},
			},
			[]*Article{
				{Title: "title1", ArticleTags: []ArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title3", ArticleTags: []ArticleTag{{MachineName: TFTPatchNoteTagName}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LOLWebsiteArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.TFTPatchNotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLWebsiteArticles.TFTPatchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
