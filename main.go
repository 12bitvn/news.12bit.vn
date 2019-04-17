package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/oauth2"
	"os"
	"text/template"
	"time"
)

var feedParser = gofeed.NewParser()

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mongoApplyURI, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		return events.APIGatewayProxyResponse{}, errors.New("MONGODB_URI is not defined")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoApplyURI))
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	linksCollection := client.Database("prod").Collection("links")
	if _, err := linksCollection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159}); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success",
	}, nil
}

func removeDuplicate(feed *gofeed.Feed) *gofeed.Feed {
	return feed
}

func commit(accessToken string, feed *gofeed.Feed) error {
	fileContent, err := render(feed)
	if err != nil {
		return err
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("New site"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("vominh"), Email: github.String("nguyenvanduocit@gmail.com")},
	}

	if _, _, err := client.Repositories.CreateFile(ctx, "12bitvn", "news.12bit.vn", fmt.Sprintf("content/links/%s", generateFileName(feed, feed.Items[0])), opts); err != nil {
		return err
	}
	return nil
}

func render(feed *gofeed.Feed) ([]byte, error) {
	mdTemplate := `---
title: "{{ .Title }}"
description: "{{ .Description }}"
date: {{ .PublishedParsed.Format "2006-01-02T15:04:05Z07:00" }}
link: "{{ .Link }}"
site: site_name
draft: false
---
`
	tmpl, err := template.New("link-template").Parse(mdTemplate)
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, feed); err != nil {
		return nil, err
	}
	return tpl.Bytes(), nil
}

func generateFileName(feed *gofeed.Feed, fieldItem *gofeed.Item) string {
	return fmt.Sprintf("%s-%s.md", slug.Make(fieldItem.Title), fieldItem.PublishedParsed.Format(time.RFC3339))
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
