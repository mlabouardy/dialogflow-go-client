package main

import (
	"encoding/json"
	"errors"
	"fmt"

	. "github.com/mlabouardy/apiai-go-client/models"
	uuid "github.com/satori/go.uuid"
)

type ApiAiClient struct {
	accessToken string
	apiLang     string
	apiVersion  string
	apiBaseUrl  string
	sessionID   string
}

func NewApiAiClient(options Options) (error, *ApiAiClient) {
	if (options == Options{} || options.AccessToken == "") {
		return errors.New("Access token is required for new ApiAiClient instance"), nil
	}

	client := &ApiAiClient{
		accessToken: options.AccessToken,
	}

	client.apiBaseUrl = options.ApiBaseUrl
	if client.apiBaseUrl == "" {
		client.apiBaseUrl = DEFAULT_BASE_URL
	}

	client.apiLang = options.ApiLang
	if client.apiLang == "" {
		client.apiLang = DEFAULT_CLIENT_LANG
	}

	client.apiVersion = options.ApiVersion
	if client.apiVersion == "" {
		client.apiVersion = DEFAULT_API_VERSION
	}

	client.sessionID = options.SessionID
	if client.sessionID == "" {
		client.sessionID = uuid.NewV4().String()
	}

	return nil, client
}

func (client *ApiAiClient) TextRequest(query string) (QueryResponse, error) {
	var queryResponse QueryResponse

	if query == "" {
		return queryResponse, errors.New("Query should not be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: RequestBody{
				Query:     query,
				Lang:      client.GetApiLang(),
				SessionID: client.GetSessionID(),
			},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

func (client *ApiAiClient) EventRequest(eventName string, eventData map[string]string) (QueryResponse, error) {
	var queryResponse QueryResponse

	if eventName == "" {
		return queryResponse, errors.New("Event name can not be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: RequestBody{
				Lang:      client.GetApiLang(),
				SessionID: client.GetSessionID(),
				Event: Event{
					Name: eventName,
					Data: eventData,
				},
			},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Retrieves a list of all entities for the agent
func (client *ApiAiClient) EntitiesFindAllRequest() ([]Entity, error) {
	var entityResponse []Entity

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody:   RequestBody{},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return entityResponse, err
	}

	err = json.Unmarshal(data, &entityResponse)

	return entityResponse, err
}

// Retrieves the specified entity
func (client *ApiAiClient) EntitiesFindByIdRequest(eid string) (Entity, error) {
	var entityResponse Entity

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody:   RequestBody{},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return entityResponse, err
	}

	err = json.Unmarshal(data, &entityResponse)

	return entityResponse, err
}

func (client *ApiAiClient) EntitiesCreateRequest(entity Entity) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: RequestBody{
				Lang:      client.GetApiLang(),
				SessionID: client.GetSessionID(),
				Name:      entity.Name,
				Entries:   entity.Entries,
			},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Adds entries to the specified entity.
func (client *ApiAiClient) EntitiesAddEntryRequest(eid string, entries []Entry) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody:   entries,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Creates or updates an array of entities
func (client *ApiAiClient) EntitiesUpdateRequest(entities []Entity) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			RequestMethod: "PUT",
			RequestBody:   entities,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Updates the specified entity
func (client *ApiAiClient) EntitiesUpdateEntityRequest(eid string, entity Entity) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			RequestMethod: "PUT",
			RequestBody:   entity,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Updates entity entries
func (client *ApiAiClient) EntitiesUpdateEntityEntriesRequest(eid string, entries []Entry) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			RequestMethod: "PUT",
			RequestBody:   entries,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Deletes the specified entity
func (client *ApiAiClient) EntitiesDeleteRequest(eid string) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			RequestMethod: "DELETE",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Deletes entity entries
func (client *ApiAiClient) EntitiesDeleteEntriesRequest(eid string, values []string) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			RequestMethod: "DELETE",
			RequestBody:   values,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Adds one or multiple user entities for a session.
func (client *ApiAiClient) UserEntitiesCreateRequest(userEntities []UserEntity) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: struct {
				SessionID string
				Entities  []UserEntity
			}{
				SessionID: client.GetSessionID(),
				Entities:  userEntities,
			},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Updates user entity specified by name
func (client *ApiAiClient) UserEntitiesUpdateRequest(name string, userEntity UserEntity) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
			RequestMethod: "PUT",
			RequestBody:   userEntity,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Gets a user entity object by name
func (client *ApiAiClient) UserEntitiesFindByNameRequest(name string) (UserEntity, error) {
	var userEntity UserEntity

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
			RequestMethod: "GET",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return userEntity, err
	}

	err = json.Unmarshal(data, &userEntity)

	return userEntity, err
}

// Deletes a user entity object with a specified name
func (client *ApiAiClient) UserEntitiesDeleteByNameRequest(name string) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
			RequestMethod: "DELETE",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Retrieves a list of all intents for the agent
func (client *ApiAiClient) IntentsFindAllRequest() ([]IntentAgent, error) {
	var intents []IntentAgent

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return intents, err
	}

	err = json.Unmarshal(data, &intents)

	return intents, err
}

// Retrieves the specified intent
func (client *ApiAiClient) IntentsFindByIdRequest(id string) (Intent, error) {
	var intent Intent

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return intent, err
	}

	err = json.Unmarshal(data, &intent)

	return intent, err
}

// Creates a new intent
func (client *ApiAiClient) IntentsCreateRequest(intent Intent) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody:   intent,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Updates the specified intent
func (client *ApiAiClient) IntentsUpdateRequest(id string, intent Intent) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			RequestMethod: "PUT",
			RequestBody:   intent,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Deletes the specified intent
func (client *ApiAiClient) IntentsDeleteRequest(id string, intent Intent) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			RequestMethod: "DELETE",
			RequestBody:   RequestBody{},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// retrieves the list of all currently active contexts for the specified session
func (client *ApiAiClient) ContextsFindAllRequest() ([]Context, error) {
	var contexts []Context

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			RequestMethod: "GET",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return contexts, err
	}

	err = json.Unmarshal(data, &contexts)

	return contexts, err
}

// Retrieves the specified context for the specified session
func (client *ApiAiClient) ContextsFindByNameRequest(name string) (Context, error) {
	var contexts Context

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
			RequestMethod: "GET",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return contexts, err
	}

	err = json.Unmarshal(data, &contexts)

	return contexts, err
}

// Adds new active contexts to the specified session
func (client *ApiAiClient) ContextsCreateRequest(contexts []Context) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			RequestMethod: "POST",
			RequestBody:   contexts,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Deletes all contexts from the specified session
func (client *ApiAiClient) ContextsDeleteRequest() (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			RequestMethod: "DELETE",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	fmt.Println(string(data))

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

// Deletes the specified context from the specified session
func (client *ApiAiClient) ContextsDeleteByNameRequest(name string) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
			RequestMethod: "DELETE",
			RequestBody:   nil,
		},
	)

	data, err := request.Perform()

	fmt.Println(string(data))

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}

func (client *ApiAiClient) GetAccessToken() string {
	return client.accessToken
}

func (client *ApiAiClient) GetApiVersion() string {
	if client.apiVersion != "" {
		return client.apiVersion
	}
	return DEFAULT_API_VERSION
}

func (client *ApiAiClient) GetApiLang() string {
	if client.apiLang != "" {
		return client.apiLang
	}
	return DEFAULT_CLIENT_LANG
}

func (client *ApiAiClient) GetBaseUrl() string {
	if client.apiBaseUrl != "" {
		return client.apiBaseUrl
	}
	return DEFAULT_BASE_URL
}

func (client *ApiAiClient) GetSessionID() string {
	return client.sessionID
}

func (client *ApiAiClient) SetSessionID(sessionID string) {
	client.sessionID = sessionID
}
