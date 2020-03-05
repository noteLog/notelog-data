package main

import (
	"context"
	"log"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v29/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func getGitHubClient(ghToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

// GetUserRepos returns a slice of the user's public GitHub Repositories
// Implementation Credit: https://github.com/lox/alfred-github-jump/repos.go
func githubGetUserRepos(ghToken string) ([]*github.Repository, error) {
	client := getGitHubClient(ghToken)
	ctx := context.Background()

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 45},
		Sort:        "pushed",
	}

	repos := []*github.Repository{}

	for {
		result, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			return repos, err
		}
		repos = append(repos, result...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	log.Printf("Fetched %v user repos.", len(repos))
	return repos, nil
}

// GetStarredRepos returns a slice of all the repositories the user starred
// Implementation Credit: https://github.com/lox/alfred-github-jump/repos.go
func githubGetStarredRepos(ghToken string) ([]*github.Repository, error) {
	client := getGitHubClient(ghToken)
	ctx := context.Background()

	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 45},
		Sort:        "pushed",
	}

	repos := []*github.Repository{}

	for {
		result, resp, err := client.Activity.ListStarred(ctx, "", opt)
		if err != nil {
			return repos, err
		}
		for _, starred := range result {
			repos = append(repos, starred.Repository)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
	log.Printf("Fetched %v starred repos.", len(repos))
	return repos, nil
}
