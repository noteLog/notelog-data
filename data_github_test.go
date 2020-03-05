package main

import (
	"os"
	"testing"
)

func TestGitHubGetUserRepos(t *testing.T) {
	_, err := githubGetUserRepos(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		t.Errorf("githubGetUserRepos() returned error %v, none expected", err)
	}
}

func TestGitHubGetStarredRepos(t *testing.T) {
	_, err := githubGetStarredRepos(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		t.Errorf("githubGetStarredRepos() returned error %v, none expected", err)
	}
}
