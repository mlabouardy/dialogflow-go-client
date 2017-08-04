package models

type RequestOptions struct {
	URI         string
	Method      string
	Body        interface{}
	QueryParams map[string]string
}
