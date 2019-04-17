package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
	"path/filepath"
	"text/template"
	"time"
)

var markdownTemplate = `---
title: "{{ .Title }}"
description: "{{ .Description }}"
date: {{ .PublishedParsed.Format "2006-01-02T15:04:05Z07:00" }}
link: "{{ .Link }}"
site: site_name
draft: false
---
`

var committer = "crawler"
var committerEmail = "12bitsvn@gmail.com"
var reponame = "news.12bit.vn"
var owner = "12bitvn"
var branch = "master"
var path = "content/links"

type Link struct {
	*gofeed.Item
}

func (link Link) IsDuplicate() bool {
	return false
}

func (link Link) Commit(client *github.Client) error {
	content, err := link.Render()
	if err != nil {
		return err
	}
	ctx := context.Background()
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String(fmt.Sprintf("add new link %s", link.ID())),
		Content:   content,
		Branch:    github.String(branch),
		Committer: &github.CommitAuthor{Name: github.String(committer), Email: github.String(committerEmail)},
	}

	if _, _, err := client.Repositories.CreateFile(ctx, owner, reponame, filepath.Join(path, fmt.Sprintf("%s.md", link.ID())), opts); err != nil {
		return err
	}
	return nil
}

func (link Link) ID() string {
	return fmt.Sprintf("%s-%s", slug.Make(link.Title), link.PublishedParsed.Format(time.RFC3339))
}

func (link Link) Render() ([]byte, error) {
	tmpl, err := template.New("link-template").Parse(markdownTemplate)
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, link); err != nil {
		return nil, err
	}
	return tpl.Bytes(), nil
}
