package models

type Metadata struct {
	IntentID                  string `json:"intentId,omitempty"`
	WebhookUsed               string `json:"webhookUsed,omitempty"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed,omitempty"`
	IntentName                string `json:"intentName,omitempty"`
}
