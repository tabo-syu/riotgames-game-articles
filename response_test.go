package riotgamespatchnotes

import (
	"reflect"
	"testing"
)

func TestArticles_FilterByTag(t *testing.T) {
	type fields struct {
		Content []*Article
	}
	type args struct {
		tagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Articles
	}{
		{
			"Test filter articles",
			fields{
				[]*Article{
					{Title: "1", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "2", ArticleTags: []ArticleTag{}},
					{Title: "3", ArticleTags: []ArticleTag{{MachineName: "bar"}}},
					{Title: "4", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "5"},
				},
			},
			args{tagName: "target"},
			&Articles{
				[]*Article{
					{Title: "1", ArticleTags: []ArticleTag{{MachineName: "target"}}},
					{Title: "4", ArticleTags: []ArticleTag{{MachineName: "target"}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Articles{
				Content: tt.fields.Content,
			}
			if got := a.FilterByTagName(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Articles.FilterByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
