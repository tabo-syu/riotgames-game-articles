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
		Articles *LOLArticles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*LOLArticle
	}{
		{
			"Test all articles",
			fields{
				&LOLArticles{
					[]*LOLArticle{
						{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []LOLArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []LOLArticleTag{}},
					},
				},
			},
			[]*LOLArticle{
				{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title2", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}}},
				{Title: "title3", ArticleTags: []LOLArticleTag{{MachineName: TFTPatchNoteTagName}}},
				{Title: "title4", ArticleTags: []LOLArticleTag{}},
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
		Articles *LOLArticles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*LOLArticle
	}{
		{
			"Test all articles",
			fields{
				&LOLArticles{
					[]*LOLArticle{
						{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []LOLArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []LOLArticleTag{}},
					},
				},
			},
			[]*LOLArticle{
				{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title2", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}}},
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
		Articles *LOLArticles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*LOLArticle
	}{
		{
			"Test all articles",
			fields{
				&LOLArticles{
					[]*LOLArticle{
						{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}}},
						{Title: "title3", ArticleTags: []LOLArticleTag{{MachineName: TFTPatchNoteTagName}}},
						{Title: "title4", ArticleTags: []LOLArticleTag{}},
					},
				},
			},
			[]*LOLArticle{
				{Title: "title1", ArticleTags: []LOLArticleTag{{MachineName: LOLPatchNoteTagName}, {MachineName: TFTPatchNoteTagName}}},
				{Title: "title3", ArticleTags: []LOLArticleTag{{MachineName: TFTPatchNoteTagName}}},
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

func TestNewValorantWebsiteArticles(t *testing.T) {
	for _, locale := range ValorantLocales {
		t.Run(fmt.Sprintf("locale: %s", locale), func(t *testing.T) {
			articles, err := NewValorantWebsiteArticles(locale)
			if err != nil {
				t.Errorf("NewValorantWebsiteArticles() error = %v", err)
			}
			t.Logf(articles.PatchNotes()[0].Title)

			time.Sleep(1 * time.Second)
		})
	}
}

func TestValorantWebsiteArticles_All(t *testing.T) {
	type fields struct {
		Articles *ValorantArticles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*ValorantArticle
	}{
		{
			"Test all articles",
			fields{
				&ValorantArticles{
					[]*ValorantArticle{
						{Title: "title1", ArticleTags: []ValorantArticleTag{{MachineName: ValorantPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []ValorantArticleTag{}},
					},
				},
			},
			[]*ValorantArticle{
				{Title: "title1", ArticleTags: []ValorantArticleTag{{MachineName: ValorantPatchNoteTagName}}},
				{Title: "title2", ArticleTags: []ValorantArticleTag{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ValorantWebsiteArticles{
				Articles: tt.fields.Articles,
			}
			if got := v.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValorantWebsiteArticles.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValorantWebsiteArticles_PatchNotes(t *testing.T) {
	type fields struct {
		Articles *ValorantArticles
	}
	tests := []struct {
		name   string
		fields fields
		want   []*ValorantArticle
	}{
		{
			"Test all articles",
			fields{
				&ValorantArticles{
					[]*ValorantArticle{
						{Title: "title1", ArticleTags: []ValorantArticleTag{{MachineName: ValorantPatchNoteTagName}}},
						{Title: "title2", ArticleTags: []ValorantArticleTag{}},
					},
				},
			},
			[]*ValorantArticle{
				{Title: "title1", ArticleTags: []ValorantArticleTag{{MachineName: ValorantPatchNoteTagName}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ValorantWebsiteArticles{
				Articles: tt.fields.Articles,
			}
			if got := v.PatchNotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LOLWebsiteArticles.PatchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
