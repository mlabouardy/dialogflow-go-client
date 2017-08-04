package models

type Query struct {
	Query           string          `json:"query,omitempty"`
	E               Event           `json:"e,omitempty"`
	Event           Event           `json:"event,omitempty"`
	V               string          `json:"v,omitempty"`
	SessionID       string          `json:"sessionId,omitempty"`
	Lang            string          `json:"lang,omitempty"`
	Contexts        []Context       `json:"contexts,omitempty"`
	ResetContexts   bool            `json:"resetContexts,omitempty"`
	Entities        []Entity        `json:"entities,omitempty"`
	Timezone        string          `json:"timezone,omitempty"`
	Location        Location        `json:"location,omitempty"`
	OriginalRequest OriginalRequest `json:"originalRequest,omitempty"`
}

type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type OriginalRequest struct {
	Source string `json:"source,omitempty"`
	Data   string `json:"data,omitempty"`
}
