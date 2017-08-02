package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/mlabouardy/apiai-go-client/models"
)

type Request struct {
	URI           string
	RequestMethod string
	Headers       map[string]string
	RequestBody   interface{}
	QueryParams   map[string]string
}

func NewRequest(client *ApiAiClient, overridedRequestOptions RequestOptions) *Request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetAccessToken(),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	request := &Request{
		URI:           overridedRequestOptions.URI,
		RequestMethod: overridedRequestOptions.RequestMethod,
		Headers:       headers,
		QueryParams:   overridedRequestOptions.QueryParams,
		RequestBody: RequestBody{
			Lang:      client.GetApiLang(),
			SessionID: client.GetSessionID(),
			Query:     overridedRequestOptions.RequestBody.Query,
			Event:     overridedRequestOptions.RequestBody.Event,
			Entities:  overridedRequestOptions.RequestBody.Entities,
		},
	}
	return request
}

func (r *Request) Perform() ([]byte, error) {
	var data []byte
	client := &http.Client{}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r.RequestBody)

	req, err := http.NewRequest(r.RequestMethod, r.URI, b)
	if r.RequestMethod == "GET" {
		req, err = http.NewRequest(r.RequestMethod, r.URI, nil)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range r.QueryParams {
		query.Add(key, value)
	}

	if err != nil {
		return data, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
