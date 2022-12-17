package riotgamespatchnotes

import (
	"context"
	"testing"
)

func TestNewLOLGameUpdatesRequest(t *testing.T) {
	type args struct {
		ctx    context.Context
		locale string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Test invalid locale",
			args{context.Background(), "foo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLOLGameUpdatesRequest(tt.args.ctx, tt.args.locale)
			if err == nil {
				t.Errorf("NewLOLGameUpdatesRequest() error = %v", err)
				return
			}
		})
	}
}