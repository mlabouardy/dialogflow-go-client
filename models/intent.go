package models

type Intent struct {
	ID                    string     `json:"id,omitempty"`
	Name                  string     `json:"name,omitempty"`
	Auto                  bool       `json:"auto,omitempty"`
	Contexts              []string   `json:"contexts,omitempty"`
	Templates             []string   `json:"templates,omitempty"`
	UserSays              []UserSay  `json:"userSays,omitempty"`
	Responses             []Response `json:"responses,omitempty"`
	Priority              int        `json:"priority,omitempty"`
	WebhookUsed           bool       `json:"webhookUsed,omitempty"`
	WebhookForSlotFilling bool       `json:"webhookForSlotFilling,omitempty"`
	FallbackIntent        bool       `json:"fallbackIntent,omitempty"`
	CortanaCommand        struct {
		NavigationOrService string
		Target              string
	}
	Events []struct {
		Name string
	}
}

type UserSay struct {
	ID   string `json:"userSays,omitempty"`
	Data struct {
		Text        string `json:"text,omitempty"`
		Meta        string `json:"meta,omitempty"`
		Alias       string `json:"alias,omitempty"`
		UserDefined bool   `json:"userDefined,omitempty"`
	} `json:"data,omitempty"`
	IsTemplate bool `json:"isTemplate,omitempty"`
	Count      int  `json:"count,omitempty"`
}

type Response struct {
	Action           string `json:"action,omitempty"`
	ResetContexts    bool   `json:"resetContexts,omitempty"`
	AffectedContexts []struct {
		Name     string `json:"name,omitempty"`
		Lifespan bool   `json:"lifespan,omitempty"`
	} `json:"affectedContexts,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Messages   []Message   `json:"message,omitempty"`
}
