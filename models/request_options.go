package models

type RequestOptions struct {
	URI           string
	RequestMethod string
	RequestBody   interface{}
	QueryParams   map[string]string
}
