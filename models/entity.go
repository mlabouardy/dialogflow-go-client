package models

type Entity struct {
	ID                 string  `json:"id,omitempty"`
	Name               string  `json:"name,omitempty"`
	Count              int     `json:"count,omitempty"`
	Preview            string  `json:"preview,omitempty"`
	IsOverridable      bool    `json:"isOverridable,omitempty"`
	IsEnum             bool    `json:"isEnum,omitempty"`
	AutomatedExpansion bool    `json:"automatedExpansion,omitempty"`
	Entries            []Entry `json:"entries,omitempty"`
}
