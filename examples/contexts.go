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

	contexts, err := client.ContextsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, context := range contexts {
		fmt.Printf("%s:%d\n", context.Name, context.Lifespan)
	}
}
