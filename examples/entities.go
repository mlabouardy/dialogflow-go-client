package main

import (
	"fmt"
	"log"

	. "github.com/mlabouardy/dialogflow-go-client"
	. "github.com/mlabouardy/dialogflow-go-client/models"
)

func main() {
	err, client := NewDialogFlowClient(Options{
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
