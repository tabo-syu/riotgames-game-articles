package riotgamesgamearticles

import (
	"context"
	"testing"
)

func TestNewLOLWebsiteRequest(t *testing.T) {
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
			_, err := NewLOLWebsiteRequest(tt.args.ctx, tt.args.locale)
			if err == nil {
				t.Errorf("NewLOLWebsiteRequest() error = %v", err)
				return
			}
		})
	}
}

func TestNewValorantWebsiteRequest(t *testing.T) {
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
			_, err := NewValorantWebsiteRequest(tt.args.ctx, tt.args.locale)
			if err == nil {
				t.Errorf("NewValorantWebsiteRequest() error = %v", err)
				return
			}
		})
	}
}
