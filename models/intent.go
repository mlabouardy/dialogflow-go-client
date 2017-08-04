package models

type Intent struct {
	ID        string
	Name      string
	Auto      bool
	Contexts  []string
	Templates []string
	UserSays  []struct {
		ID   string
		Data struct {
			Text        string
			Meta        string
			Alias       string
			UserDefined bool
		}
		IsTemplate bool
		Count      int
	}
	Responses []struct {
		Action           string
		ResetContexts    bool
		AffectedContexts []struct {
			Name     string
			Lifespan bool
		}
		Parameters []struct {
			Name         string
			Value        string
			DefaultValue string
			Required     bool
			DataType     string
			Prompts      []string
			IsList       bool
		}
		Messages []Message
	}
	Priority              int
	WebhookUsed           bool
	WebhookForSlotFilling bool
	FallbackIntent        bool
	CortanaCommand        struct {
		NavigationOrService string
		Target              string
	}
	Events []struct {
		Name string
	}
}
