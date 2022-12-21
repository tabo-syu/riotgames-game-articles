package riotgamesgamearticles

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
		want   *Articles[*LOLArticle]
	}{
		{
			"Test filter articles",
			fields{
				[]*LOLArticle{
					{Title: "1", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "2", ArticleTags: []ArticleTag{}},
					{Title: "3", ArticleTags: []ArticleTag{{MachineName: "bar"}}},
					{Title: "4", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "5"},
				},
			},
			args{tagName: "target"},
			&Articles[*LOLArticle]{
				[]*LOLArticle{
					{Title: "1", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "4", ArticleTags: []ArticleTag{{MachineName: "target"}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Articles[*LOLArticle]{
				Content: tt.fields.Content,
			}
			if got := a.FilterByTagName(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesResponse.FilterByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
