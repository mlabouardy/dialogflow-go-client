package main

import (
	"fmt"
	"log"

	. "github.com/mlabouardy/apiai-go-client"
	. "github.com/mlabouardy/apiai-go-client/models"
)

func main() {
	err, client := NewApiAiClient(Options{
		AccessToken: "0b8c9af465a141fc9949fb8e87f0fa1a",
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
