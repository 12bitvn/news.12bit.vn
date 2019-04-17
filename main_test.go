package main

import (
	"github.com/mmcdole/gofeed"
	"testing"
	"time"
)

func Test_createCommit(t *testing.T) {
	type args struct {
		accessToken string
		feedItem    gofeed.Item
	}
	now, _ := time.Parse(time.RFC3339, "2019-04-17 00:57:11 +0700")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid access token",
			args: args{
				accessToken: "fda8b31b2bcf7766ddfc029d67c94d148d2677d0",
				feedItem: gofeed.Item{
					Title:           "title",
					Description:     "description",
					PublishedParsed: &now,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createCommit(tt.args.accessToken, tt.args.feedItem); (err != nil) != tt.wantErr {
				t.Errorf("createCommit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
