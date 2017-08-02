package models

type RequestOptions struct {
	URI           string
	RequestMethod string
	RequestBody   RequestBody
	QueryParams   map[string]string
}
