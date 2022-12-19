package riotgamespatchnotes

import (
	"reflect"
	"testing"
)

func TestArticlesResponse_FilterByTag(t *testing.T) {
	type fields struct {
		Content []*LOLArticle
	}
	type args struct {
		tagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LOLArticles
	}{
		{
			"Test filter articles",
			fields{
				[]*LOLArticle{
					{Title: "1", ArticleTags: []LOLArticleTag{{MachineName: "target"}}},
					{Title: "2", ArticleTags: []LOLArticleTag{}},
					{Title: "3", ArticleTags: []LOLArticleTag{{MachineName: "bar"}}},
					{Title: "4", ArticleTags: []LOLArticleTag{{MachineName: "target"}}},
					{Title: "5"},
				},
			},
			args{tagName: "target"},
			&LOLArticles{
				[]*LOLArticle{
					{Title: "1", ArticleTags: []LOLArticleTag{{MachineName: "target"}}},
					{Title: "4", ArticleTags: []LOLArticleTag{{MachineName: "target"}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := LOLArticles{
				Content: tt.fields.Content,
			}
			if got := a.FilterByTagName(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesResponse.FilterByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
