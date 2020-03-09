package main

import (
	"os"

	"github.com/jasonlvhit/gocron"
)

func main() {
	esURL := os.Getenv("ES_URL")
	esPWD := os.Getenv("ES_PWD")
	ghToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	gocron.Every(1).Hour().Do(esIndexGitHub, esURL, esPWD, ghToken)
	gocron.Every(1).Day().Do(esIndexBlog, esURL, esPWD)
	<-gocron.Start()
}
