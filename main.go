package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadDotEnvErr := godotenv.Load()
	if loadDotEnvErr != nil {
		log.Printf("loadDotEnvErr = %+v\n", loadDotEnvErr)
	}
	getESInfo()
}
