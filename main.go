package main

import (
	"flag"
	"log"
)

func main() {
	var esURL = flag.String("esurl", "", "URL of your Elasticsearch service")
	var esPWD = flag.String("espwd", "", "Password for your Elasticsearch service")
	var ghToken = flag.String("ghtoken", "", "GitHub Access Token")
	flag.Parse()

	// gocron.Every(1).Hour().Do(githubCronPlaceholder)
	// gocron.Every(1).Day().Do(blogCronPlaceholder)
	// <-gocron.Start()

	esIndexGitHub(*esURL, *esPWD, *ghToken)
}

func githubCronPlaceholder() {
	log.Printf("Fetching GitHub Repos")
}

func blogCronPlaceholder() {
	log.Printf("Fetching Blog Posts")
}
