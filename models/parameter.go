package models

type Parameter struct {
	Name         string   `json:"name,omitempty"`
	Value        string   `json:"value,omitempty"`
	DefaultValue string   `json:"defaultValue,omitempty"`
	Required     bool     `json:"required,omitempty"`
	DataType     string   `json:"dataType,omitempty"`
	Prompts      []string `json:"prompts,omitempty"`
	IsList       bool     `json:"isList,omitempty"`
}
