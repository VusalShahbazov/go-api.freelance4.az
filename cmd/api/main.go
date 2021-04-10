package main

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/api/server"
	"github.com/joho/godotenv"
	"log"
)

func main()  {
	err := godotenv.Load()

	apiServer := server.NewApiServer()

	err = apiServer.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
