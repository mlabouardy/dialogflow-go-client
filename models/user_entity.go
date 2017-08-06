package models

type UserEntity struct {
	SessionID string  `json:"sessionId,omitempty"`
	Name      string  `json:"name,omitempty"`
	Extend    bool    `json:"extend,omitempty"`
	Entries   []Entry `json:"entries,omitempty"`
}
