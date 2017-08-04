package models

type IntentAgent struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	ContextIn  []string `json:"contextIn,omitempty"`
	ContextOut []struct {
		Name     string `json:"name,omitempty"`
		Lifespan int    `json:"lifespan,omitempty"`
	} `json:"contextOut,omitempty"`
	Actions    []string `json:"actions,omitempty"`
	Parameters []struct {
		Name         string   `json:"name,omitempty"`
		Value        string   `json:"value,omitempty"`
		DefaultValue string   `json:"defaultValue,omitempty"`
		Required     bool     `json:"required,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Prompts      []string `json:"prompts,omitempty"`
	} `json:"parameters,omitempty"`
	Priority       int  `json:"priority,omitempty"`
	FallbackIntent bool `json:"fallbackIntent,omitempty"`
}
