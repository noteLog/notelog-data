package main

import "testing"

func TestGitHubGetUserRepos(t *testing.T) {
	_, err := githubGetUserRepos()
	if err != nil {
		t.Errorf("githubGetUserRepos() returned error %v, none expected", err)
	}
}

func TestGitHubGetStarredRepos(t *testing.T) {
	_, err := githubGetStarredRepos()
	if err != nil {
		t.Errorf("githubGetStarredRepos() returned error %v, none expected", err)
	}
}
