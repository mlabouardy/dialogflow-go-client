package models

type EntityRequest struct {
	Name    string  `json:"name,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
}
