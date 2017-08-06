package models

type Event struct {
	Name string            `json:"name,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}
