package apiai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/mlabouardy/apiai-go-client/models"
)

type Request struct {
	URI         string
	Method      string
	Headers     map[string]string
	Body        interface{}
	QueryParams map[string]string
}

// Initialize a new HTTP request
func NewRequest(client *ApiAiClient, overridedRequestOptions RequestOptions) *Request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetAccessToken(),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	request := &Request{
		URI:         overridedRequestOptions.URI,
		Method:      overridedRequestOptions.Method,
		Headers:     headers,
		QueryParams: overridedRequestOptions.QueryParams,
		Body:        overridedRequestOptions.Body,
	}

	return request
}

// Execute an HTTP request
func (r *Request) Perform() ([]byte, error) {
	var data []byte
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URI, nil)

	if r.Method != "GET" {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(r.Body)
		req, err = http.NewRequest(r.Method, r.URI, b)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range r.QueryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

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
