package models

import (
	"fmt"
	"reflect"
)

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

func (query Query) ToMap() map[string]string {
	params := make(map[string]string)

	if query.Query != "" {
		params["query"] = query.Query
	}

	if !reflect.DeepEqual(query.E, Event{}) {
		params["e"] = query.Event.Name
	}

	if !reflect.DeepEqual(query.Contexts, []Context{}) {
		params["contexts"] = query.Contexts[0].Name
	}

	if !reflect.DeepEqual(query.Location, Location{}) {
		params["latitude"] = fmt.Sprintf("%f", query.Location.Latitude)
		params["longitude"] = fmt.Sprintf("%f", query.Location.Longitude)
	}

	params["v"] = query.V
	params["sessionId"] = query.SessionID
	params["lang"] = query.Lang

	return params
}
