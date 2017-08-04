package models

type Message struct{
  Type int `json:"type"`
  Speech []string `json:"speech"`
}
