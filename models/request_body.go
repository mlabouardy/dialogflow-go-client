package models

type RequestBody struct {
	Query     string   `json:"query,omitempty"`
	Event     Event    `json:"event,omitempty"`
	SessionID string   `json:"sessionId"`
	Lang      string   `json:"lang,omitempty"`
	Entities  []Entity `json:"entities,omitempty"`
}

type Event struct {
	Name string            `json:"name,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}

type Entity struct {
	Name    string  `json:"name,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
}

type Entry struct {
	Value    string   `json:"value,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
}
