package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"github.com/mmcdole/gofeed"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"
)

type Site struct {
	RssURL string `json:"rss_url"`
}

var siteFileUrl = "https://raw.githubusercontent.com/12bitvn/news.12bit.vn/1fc1fd073197714f5f053083bb8853e967a7a972/data/links.json"

var feedParser = gofeed.NewParser()
var httpClient = &http.Client{Timeout: 10 * time.Second}
var accessToken, _ = os.LookupEnv("GITHUB_ACCESS_TOKEN")
var githubClient = newithubClient(accessToken)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sites, err := fetchSiteList(siteFileUrl)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	for _, site := range sites {
		feed, err := feedParser.ParseURL(site.RssURL)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		for _, feedItem := range feed.Items {
			link := Link{feedItem}
			if link.IsDuplicate() {
				continue
			}
			if err := link.Commit(githubClient); err != nil {
				return events.APIGatewayProxyResponse{}, err
			}
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success",
	}, nil
}

func fetchSiteList(fileUrl string) ([]Site, error) {
	var sites []Site
	r, err := httpClient.Get(fileUrl)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&sites); err != nil {
		return nil, err
	}
	return sites, nil
}

func newithubClient(accessToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func main() {
	lambda.Start(handler)
}
