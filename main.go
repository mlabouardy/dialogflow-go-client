package main

import (
	"fmt"
	"log"

	. "github.com/mlabouardy/apiai-go-client/models"
)

func main() {
	err, client := NewApiAiClient(Options{
		AccessToken: "e051cfede76a491899ec942613a89c06",
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

	resp3, err := client.EntitiesFindByIdRequest("a441dd36-a1f8-4dbe-b6d3-3b0b2c4031e5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp3)

	entity := Entity{
		Name: "food",
		Entries: []Entry{
			Entry{
				Value:    "Coffee",
				Synonyms: []string{"maker", "machine", "coffee"},
			},
		},
	}

	resp4, err := client.EntitiesCreateRequest(entity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp4)
}
