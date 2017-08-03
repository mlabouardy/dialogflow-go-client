package main

import (
	"encoding/json"
	"errors"

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

/*
func (client *ApiAiClient) EntitiesCreateEntryRequest(entry Entry) (QueryResponse, error) {
	var queryResponse QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: RequestBody{
				Name:    entity.Name,
				Entries: entity.Entries,
			},
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}*/

/*




// Creates a new entity
func (client *ApiAiClient) EntitiesCreateRequest(entityRequest EntityRequest) {
	var queryResponse QueryResponse

	if entityRequest == EntityRequest{} {
		return queryResponse, errors.New("entity name can not be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody:   entityRequest,
		},
	)

	data, err := request.Perform()

	if err != nil {
		return queryResponse, err
	}

	err = json.Unmarshal(data, &queryResponse)

	return queryResponse, err
}*/

/*
func (client *ApiAiClient) UserEntitiesCreateRequest(entities []Entity) (error, ServerResponse) {
	if len(entities) == 0 {
		return errors.New("entities can not be empty"), ServerResponse{}
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities?v=" + client.GetApiVersion(),
			RequestMethod: "POST",
			RequestBody: RequestBody{
				Entities: entities,
			},
		},
	)

	return request.Perform()
}

func (client *ApiAiClient) UserEntitiesRetrieveRequest(name string) (error, ServerResponse) {
	if name == "" {
		return errors.New("name can not be empty"), ServerResponse{}
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody:   RequestBody{},
		},
	)

	return request.Perform()
}

func (client *ApiAiClient) UserEntitiesDeleteRequest(name string) (error, ServerResponse) {
	if name == "" {
		return errors.New("name can not be empty"), ServerResponse{}
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
			RequestMethod: "DELETE",
			RequestBody:   RequestBody{},
		},
	)

	return request.Perform()
}

func (client *ApiAiClient) ContextsCreateRequest() (error, ServerResponse) {
	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
			RequestBody: RequestBody{
				SessionID: client.GetSessionID(),
			},
		},
	)

	return request.Perform()
}

func (client *ApiAiClient) ContextsRetriveRequest(name string) (error, ServerResponse) {
	if name == "" {
		return errors.New("context name can not be empty"), ServerResponse{}
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "contexts/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
			RequestMethod: "GET",
			RequestBody:   RequestBody{},
		},
	)

	return request.Perform()
}*/

/*
func (client *ApiAiClient) UserEntitiesUpdateRequest(name string, entries []Entry) (error, ServerResponse) {
	if name == "" {
		return errors.New("name can not be empty"), ServerResponse{}
	}

	if len(entities) == 0 {
		return errors.New("entities can not be empty"), ServerResponse{}
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:           client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
			RequestMethod: "GET",
		},
		RequestBody{},
	)

	return request.Perform()
}*/

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
