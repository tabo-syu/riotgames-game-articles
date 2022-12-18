package riotgamespatchnotes

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewLOLArticles(t *testing.T) {
	for _, locale := range LOLLocales {
		t.Run(fmt.Sprintf("locale: %s", locale), func(t *testing.T) {
			articles, err := NewLOLArticles(locale)
			if err != nil {
				t.Errorf("NewLOLArticles() error = %v", err)
			}
			t.Logf(articles.LOLPatchNotes()[0].Title)

			time.Sleep(1 * time.Second)
		})
	}
}

func TestLOLArticles_All(t *testing.T) {
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
			l := LOLArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLArticles.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOLArticles_LOLPatchNotes(t *testing.T) {
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
			l := LOLArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.LOLPatchNotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLArticles.LOLPatchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOLArticles_TFTPatchNotes(t *testing.T) {
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
			l := LOLArticles{
				Articles: tt.fields.Articles,
			}
			if got := l.TFTPatchNotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLArticles.TFTPatchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
