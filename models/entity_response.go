package models

type EntityResponse struct {
	Entities []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Count   int    `json:"count"`
		Preview string `json:"preview"`
	} `json:"entities,omitempty"`
	Status struct {
		Code         int    `json:"code"`
		ErrorDetails string `json:"errorDetails"`
		ErrorID      string `json:"errorId"`
		ErrorType    string `json:"errorType"`
	} `json:"status"`
}
