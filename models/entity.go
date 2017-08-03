package models

type Entity struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Count              int     `json:"count,omitempty"`
	Preview            string  `json:"preview,omitempty"`
	IsOverridable      bool    `json:"isOverridable,omitempty"`
	IsEnum             bool    `json:"isEnum,omitempty"`
	AutomatedExpansion bool    `json:"automatedExpansion,omitempty"`
	Entries            []Entry `json:"entries,omitempty"`
}
