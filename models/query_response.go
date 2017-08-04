package models

import "time"

type QueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Result    Result    `json:"result,omitempty"`
	Status    Status    `json:"status,omitempty"`
	SessionID string    `json:"sessionId,omitempty"`
	Lang      string    `json:"lang,omitempty"`
}

type Result struct {
	Source           string            `json:"source,omitempty"`
	Action           string            `json:"action,omitempty"`
	ResolvedQuery    string            `json:"resolvedQuery,omitempty"`
	ActionIncomplete bool              `json:"actionIncomplete,omitempty"`
	Parameters       map[string]string `json:"parameters,omitempty"`
	Contexts         []Context         `json:"contexts,omitempty"`
	Metadata         Metadata          `json:"metadata,omitempty"`
	Fulfillment      Fulfillment       `json:"fulfillment,omitempty"`
	Score            float32           `json:"score,omitempty"`
}

type Fulfillment struct {
	Speech   string    `json:"speech,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}
