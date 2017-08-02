package main

import (
	"fmt"
	"log"

	. "github.com/mlabouardy/apiai-go-client/models"
)

func main() {
	err, client := NewApiAiClient(Options{
		AccessToken: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.TextRequest("hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	/*err, resp = client.ContextsCreateRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)*/

	resp, err = client.EventRequest("test_event", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	resp2, err := client.EntitiesFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp2)
}
