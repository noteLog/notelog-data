package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v29/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/joho/godotenv"
)

func getGitHubClient() *github.Client {
	if len(os.Getenv("GITHUB_ACCESS_TOKEN")) == 0 {
		loadDotEnvErr := godotenv.Load()
		if loadDotEnvErr != nil || len(os.Getenv("GITHUB_ACCESS_TOKEN")) == 0 {
			log.Printf("Error loading environment variables")
		}
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

// GetUserRepos returns a slice of the user's public GitHub Repositories
// Implementation Credit: https://github.com/lox/alfred-github-jump/repos.go
func githubGetUserRepos() ([]*github.Repository, error) {
	client := getGitHubClient()
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
func githubGetStarredRepos() ([]*github.Repository, error) {
	client := getGitHubClient()
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
