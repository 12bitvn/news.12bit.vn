package main

import (
	"reflect"
	"testing"
)

func Test_fetchSiteList(t *testing.T) {
	type args struct {
		fileUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    []Site
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				fileUrl: "https://raw.githubusercontent.com/12bitvn/news.12bit.vn/1fc1fd073197714f5f053083bb8853e967a7a972/data/links-example.json",
			},
			want: []Site{
				{
					RssURL: "https://12bit.vn/index.xml",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchSiteList(tt.args.fileUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchSiteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchSiteList() = %v, want %v", got, tt.want)
			}
		})
	}
}
