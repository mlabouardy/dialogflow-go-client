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

	intents, err := client.IntentsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, intent := range intents {
		fmt.Println(intent.Name)
	}
}
