package models

type Context struct {
	Name       string           `json:"name"`
	Lifespan   int              `json:"lifespan"`
	Parameters ContextParameter `json:"parameters"`
}

type ContextParameter struct {
	IntentAction string `json:"intent_action"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}
