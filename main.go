package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accessToken, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	if !ok {
		return events.APIGatewayProxyResponse{}, errors.New("Can not get Github access token from GITHUB_ACCESS_TOKEN")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	fileContent := []byte("This is the content of my file\nand the 2nd line of it")
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("This is my commit message"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("vominh"), Email: github.String("nguyenvanduocit@gmail.com")},
	}
	_, _, err := client.Repositories.CreateFile(ctx, "12bitvn", "news.12bit.vn", "myNewFile.md", opts)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
