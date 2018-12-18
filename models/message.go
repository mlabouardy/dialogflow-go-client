package models

type Message struct {
	Type   interface{} `json:"type,omitempty"`
	Speech string      `json:"speech,omitempty"`
}
