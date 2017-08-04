package models

type UserEntity struct {
	SessionID string  `json:"sessionId"`
	Name      string  `json:"name"`
	Extend    bool    `json:"extend"`
	Entries   []Entry `json:"entries"`
}
