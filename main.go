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
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
	"time"
)

var feedParser = gofeed.NewParser()

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	filenames := []string{}
	for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strings.Join(filenames, ";"),
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
