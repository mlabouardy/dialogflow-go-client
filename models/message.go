package models

type Message struct {
	Type   int      `json:"type,omitempty"`
	Speech []string `json:"speech,omitempty"`
}
