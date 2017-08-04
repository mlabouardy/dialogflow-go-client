package models

import "time"

type QueryResponse struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Result    struct {
		Source           string            `json:"source,omitempty"`
		Action           string            `json:"action,omitempty"`
		ResolvedQuery    string            `json:"resolvedQuery,omitempty"`
		ActionIncomplete bool              `json:"actionIncomplete,omitempty"`
		Parameters       map[string]string `json:"parameters,omitempty"`
		Contexts         []struct {
			Name       string            `json:"name,omitempty"`
			Parameters map[string]string `json:"parameters,omitempty"`
			Lifespan   int               `json:"lifespan,omitempty"`
		} `json:"contexts,omitempty"`
		Metadata struct {
			IntentID                  string `json:"intentId,omitempty"`
			WebhookUsed               string `json:"webhookUsed,omitempty"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed,omitempty"`
			IntentName                string `json:"intentName,omitempty"`
		} `json:"metadata,omitempty"`
		Fulfillment struct {
			Speech   string `json:"speech,omitempty"`
			Messages []struct {
				Type   int    `json:"type,omitempty"`
				Speech string `json:"speech,omitempty"`
			} `json:"messages,omitempty"`
		} `json:"fulfillment,omitempty"`
		Score float32 `json:"score,omitempty"`
	} `json:"result,omitempty"`
	Status struct {
		Code         int    `json:"code,omitempty"`
		ErrorDetails string `json:"errorDetails,omitempty"`
		ErrorID      string `json:"errorId,omitempty"`
		ErrorType    string `json:"errorType,omitempty"`
	} `json:"status,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	Lang      string `json:"lang,omitempty"`
}
