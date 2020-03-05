package main

import (
	"os"
	"testing"
)

func TestESGetClient(t *testing.T) {
	_, err := esGetClient(os.Getenv("ELASTICSEARCH_URL"), os.Getenv("ELASTICSEARCH_PWD"))
	if err != nil {
		t.Errorf("Getting ElasticSearch Client returned an error %v+", err)
	}
}

func TestESIndexGitHub(t *testing.T) {
	esIndexGitHub(os.Getenv("ELASTICSEARCH_URL"), os.Getenv("ELASTICSEARCH_PWD"), os.Getenv("GITHUB_ACCESS_TOKEN"))
}
