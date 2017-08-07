package apiai

import (
	"encoding/json"
	"errors"
	"reflect"

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

// Create API.AI instance
func NewApiAiClient(options Options) (error, *ApiAiClient) {
	if (reflect.DeepEqual(options, Options{}) || options.AccessToken == "") {
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

// Takes natural language text and information as query parameters and returns information as JSON
func (client *ApiAiClient) QueryFindRequest(query Query) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(query, Query{}) {
		return response, errors.New("query cannot be empty")
	}

	if query.V != "" {
		query.V = client.GetApiVersion()
	}

	if query.Lang != "" {
		query.Lang = client.GetApiLang()
	}

	if query.SessionID != "" {
		query.SessionID = client.GetSessionID()
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:         client.GetBaseUrl() + "query",
			Method:      "GET",
			Body:        nil,
			QueryParams: query.ToMap(),
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Takes natural language text and information as JSON in the POST body and returns information as JSON
func (client *ApiAiClient) QueryCreateRequest(query Query) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(query, Query{}) {
		return response, errors.New("query cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   query,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves a list of all entities for the agent
func (client *ApiAiClient) EntitiesFindAllRequest() ([]Entity, error) {
	var response []Entity

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves the specified entity
func (client *ApiAiClient) EntitiesFindByIdRequest(eid string) (Entity, error) {
	var response Entity

	if eid == "" {
		return response, errors.New("eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Creates a new entity
func (client *ApiAiClient) EntitiesCreateRequest(entity Entity) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(entity, Entity{}) {
		return response, errors.New("entity cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   entity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Adds entries to the specified entity.
func (client *ApiAiClient) EntitiesAddEntryRequest(eid string, entries []Entry) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(entries, []Entry{}) || eid == "" {
		return response, errors.New("entries and eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   entries,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Creates or updates an array of entities
func (client *ApiAiClient) EntitiesUpdateRequest(entities []Entity) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(entities, []Entity{}) {
		return response, errors.New("entities cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entities,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Updates the specified entity
func (client *ApiAiClient) EntitiesUpdateEntityRequest(eid string, entity Entity) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(entity, Entity{}) || eid == "" {
		return response, errors.New("entity and eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Updates entity entries
func (client *ApiAiClient) EntitiesUpdateEntityEntriesRequest(eid string, entries []Entry) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(entries, Entry{}) || eid == "" {
		return response, errors.New("entries and eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entries,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes the specified entity
func (client *ApiAiClient) EntitiesDeleteRequest(eid string) (QueryResponse, error) {
	var response QueryResponse

	if eid == "" {
		return response, errors.New("eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes entity entries
func (client *ApiAiClient) EntitiesDeleteEntriesRequest(eid string, values []string) (QueryResponse, error) {
	var response QueryResponse

	if len(values) == 0 || eid == "" {
		return response, errors.New("values and eid cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
			Method: "DELETE",
			Body:   values,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Adds one or multiple user entities for a session.
func (client *ApiAiClient) UserEntitiesCreateRequest(userEntities []UserEntity) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(userEntities, []UserEntity{}) {
		return response, errors.New("user entities cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities?v=" + client.GetApiVersion(),
			Method: "POST",
			Body: struct {
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
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Updates user entity specified by name
func (client *ApiAiClient) UserEntitiesUpdateRequest(name string, userEntity UserEntity) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(userEntity, UserEntity{}) || name == "" {
		return response, errors.New("user entity and name cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   userEntity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Gets a user entity object by name
func (client *ApiAiClient) UserEntitiesFindByNameRequest(name string) (UserEntity, error) {
	var response UserEntity

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes a user entity object with a specified name
func (client *ApiAiClient) UserEntitiesDeleteByNameRequest(name string) (QueryResponse, error) {
	var response QueryResponse

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves a list of all intents for the agent
func (client *ApiAiClient) IntentsFindAllRequest() ([]IntentAgent, error) {
	var response []IntentAgent

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves the specified intent
func (client *ApiAiClient) IntentsFindByIdRequest(id string) (Intent, error) {
	var response Intent

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Creates a new intent
func (client *ApiAiClient) IntentsCreateRequest(intent Intent) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(intent, Intent{}) {
		return response, errors.New("intent cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   intent,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Updates the specified intent
func (client *ApiAiClient) IntentsUpdateRequest(id string, intent Intent) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(intent, Intent{}) || id == "" {
		return response, errors.New("intent and id cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   intent,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes the specified intent
func (client *ApiAiClient) IntentsDeleteRequest(id string) (QueryResponse, error) {
	var response QueryResponse

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// retrieves the list of all currently active contexts for the specified session
func (client *ApiAiClient) ContextsFindAllRequest() ([]Context, error) {
	var response []Context

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves the specified context for the specified session
func (client *ApiAiClient) ContextsFindByNameRequest(name string) (Context, error) {
	var response Context

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Adds new active contexts to the specified session
func (client *ApiAiClient) ContextsCreateRequest(contexts []Context) (QueryResponse, error) {
	var response QueryResponse

	if reflect.DeepEqual(contexts, []Context{}) {
		return response, errors.New("contexts cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "POST",
			Body:   contexts,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes all contexts from the specified session
func (client *ApiAiClient) ContextsDeleteRequest() (QueryResponse, error) {
	var response QueryResponse

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Deletes the specified context from the specified session
func (client *ApiAiClient) ContextsDeleteByNameRequest(name string) (QueryResponse, error) {
	var response QueryResponse

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GET API.AI access token
func (client *ApiAiClient) GetAccessToken() string {
	return client.accessToken
}

// GET API.AI version
func (client *ApiAiClient) GetApiVersion() string {
	if client.apiVersion != "" {
		return client.apiVersion
	}
	return DEFAULT_API_VERSION
}

// GET API.AI language
func (client *ApiAiClient) GetApiLang() string {
	if client.apiLang != "" {
		return client.apiLang
	}
	return DEFAULT_CLIENT_LANG
}

// Get API.AI base url
func (client *ApiAiClient) GetBaseUrl() string {
	if client.apiBaseUrl != "" {
		return client.apiBaseUrl
	}
	return DEFAULT_BASE_URL
}

// Get current session ID
func (client *ApiAiClient) GetSessionID() string {
	return client.sessionID
}

// Set a new seesion ID
func (client *ApiAiClient) SetSessionID(sessionID string) {
	client.sessionID = sessionID
}
