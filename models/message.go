package models

type Message struct {
	Type     string `json:"type,omitempty"`
	Speech   string `json:"textToSpeech,omitempty"`
	Platform string `json:"platform,omitempty"`
}
