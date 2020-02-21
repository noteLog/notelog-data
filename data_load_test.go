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

func TestGetESInfo(t *testing.T) {
	client, server := getESInfo()

	expectedClient := "8.0.0-SNAPSHOT"
	expectedServer := "7.6.0"
	if client != expectedClient {
		t.Errorf("Client version %v does not match expected %v", client, expectedClient)
	}
	if server != expectedServer {
		t.Errorf("Client version %v does not match expected %v", server, expectedServer)
	}
}
