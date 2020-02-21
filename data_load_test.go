package main

import (
	"testing"
)

func TestGetESClient(t *testing.T) {
	_, err := getESClient()
	if err != nil {
		t.Errorf("Getting ElasticSearch Client returned an error %v+", err)
	}
}
