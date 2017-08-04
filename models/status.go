package models

type Status struct {
	Code         int    `json:"code,omitempty"`
	ErrorDetails string `json:"errorDetails,omitempty"`
	ErrorID      string `json:"errorId,omitempty"`
	ErrorType    string `json:"errorType,omitempty"`
}
