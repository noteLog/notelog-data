package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {
	loadDotEnvErr := godotenv.Load()
	if loadDotEnvErr != nil {
		log.Printf("loadDotEnvErr = %+v\n", loadDotEnvErr)
	}

	log.Println("Creating Cron")
	c := cron.New()

	blogSpec := "0 1 * * *"
	githubSpec := "0 * * * *"

	log.Println("Adding Blog Table Update job")
	c.AddFunc(blogSpec, func() {
		CreateUpdateBlogTable(os.Getenv("DB_CONNECTION"))
	})

	log.Println("Adding GitHub Table Update job")
	c.AddFunc(githubSpec, func() {
		log.Println("GitHub Table Update PlaceHolder")
	})

	log.Println("Starting Cron")
	c.Start()
	log.Println("Cron is Running")
	<-make(chan struct{})
}

func cronFunc() {

}
