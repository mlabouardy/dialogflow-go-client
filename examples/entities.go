package main

import (
	"fmt"
	"log"

	. "github.com/mlabouardy/apiai-go-client"
	. "github.com/mlabouardy/apiai-go-client/models"
)

func main() {
	err, client := NewApiAiClient(Options{
		AccessToken: "<API.AI TOKEN GOES HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}

	entities, err := client.EntitiesFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Println(entity.Name)
	}
}
