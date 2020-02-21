package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
)

// CreateUpdateBlogTable retrieves the list of my blog posts from https://tansawit.me/posts
// And feed each post data into the main MySQL database

func getESClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		// ...
		Username: "elastic",
		Password: os.Getenv("ELASTICSEARCH_PWD"),
	}
	es, err := elasticsearch.NewClient(cfg)
	return es, err
}

func getESInfo() (string, string) {

	log.SetFlags(0)

	var r map[string]interface{}

	//Get cluster info
	es, err := getESClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return elasticsearch.Version, r["version"].(map[string]interface{})["number"].(string)
}
