package models

import "time"

type QueryResponse struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Result    struct {
		Source           string            `json:"source"`
		Action           string            `json:"action"`
		ResolvedQuery    string            `json:"resolvedQuery"`
		ActionIncomplete bool              `json:"actionIncomplete"`
		Parameters       map[string]string `json:"parameters"`
		Contexts         []struct {
			Name       string            `json:"name"`
			Parameters map[string]string `json:"parameters"`
			Lifespan   int               `json:"lifespan"`
		} `json:"contexts"`
		Metadata struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string `json:"speech"`
			Messages []struct {
				Type   int    `json:"type"`
				Speech string `json:"speech"`
			} `json:"messages"`
		} `json:"fulfillment"`
		Score float32 `json:"score"`
	} `json:"result"`
	Status struct {
		Code         int    `json:"code"`
		ErrorDetails string `json:"errorDetails"`
		ErrorID      string `json:"errorId"`
		ErrorType    string `json:"errorType"`
	} `json:"status"`
	SessionID string `json:"sessionId"`
	Lang      string `json:"lang"`
}
