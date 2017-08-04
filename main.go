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

	entries := []Entry{
		Entry{
			Value:    "Vegetables",
			Synonyms: []string{"tomatoes", "potatoes", "onions"},
		},
	}

	resp5, err := client.EntitiesAddEntryRequest("a441dd36-a1f8-4dbe-b6d3-3b0b2c4031e5", entries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp5)

	/*entities := []Entity{
		Entity{
			Name: "cat",
			Entries: []Entry{
				Entry{
					Value:    "cat",
					Synonyms: []string{"cat", "kitty"},
				},
			},
		},
		Entity{
			Name: "dog",
			Entries: []Entry{
				Entry{
					Value:    "dog",
					Synonyms: []string{"dog", "puppy"},
				},
			},
		},
	}

	resp6, err := client.EntitiesUpdateRequest(entities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp6)*/

	entity2 := Entity{
		ID:   "4f58e9aa-74bc-4a52-9efb-a4cf055c003c",
		Name: "utility types",
		Entries: []Entry{
			Entry{
				Value:    "Electricity",
				Synonyms: []string{"electricity", "electrical"},
			},
			Entry{
				Value:    "Gas",
				Synonyms: []string{"gas", "natural language"},
			},
		},
	}

	resp7, err := client.EntitiesUpdateEntityRequest("4f58e9aa-74bc-4a52-9efb-a4cf055c003c", entity2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp7)

	entries2 := []Entry{
		Entry{
			Value:    "blue",
			Synonyms: []string{"blue", "turquoise"},
		},
	}

	resp8, err := client.EntitiesUpdateEntityEntriesRequest("4f58e9aa-74bc-4a52-9efb-a4cf055c003c", entries2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp8)

	resp9, err := client.EntitiesDeleteRequest("e37ffa28-d33a-41d9-b20c-731eda70edbc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp9)

	userEntities := []UserEntity{
		UserEntity{
			Name: "city",
			Entries: []Entry{
				Entry{
					Value:    "city",
					Synonyms: []string{"electricity", "electrical", "casalablanca"},
				},
			},
		},
	}

	resp10, err := client.UserEntitiesCreateRequest(userEntities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	fmt.Println(resp10)

	resp11, err := client.UserEntitiesFindByNameRequest("city")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp11)

	resp12, err := client.IntentsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp12)

	resp13, err := client.IntentsFindByIdRequest("4505109a-8b7c-4531-b144-0191ffa2a5aa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp13)

	resp14, err := client.ContextsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp14)

	resp15, err := client.ContextsFindByNameRequest("test_event")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp15)

	contexts := []Context{
		Context{
			Name:     "farwell",
			Lifespan: 3,
			Parameters: ContextParameter{
				Name: "John",
			},
		},
	}

	resp16, err := client.ContextsCreateRequest(contexts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp16)

	resp17, err := client.ContextsDeleteRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp17)

}
