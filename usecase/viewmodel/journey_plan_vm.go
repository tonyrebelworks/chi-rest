package viewmodel

// JourneyPlanVM ...
type JourneyPlanVM struct {
	ID              uint   ` json:"id"`
	Code            string ` json:"code"`
	JourneyName     string ` json:"journeyName"`
	JourneySchedule string ` json:"journeySchedule"`
	// Salesman        string              ` json:"assignedAuditor"`
	// Activity        string              ` json:"activity"`
	Signatures      string              ` json:"signatures"`
	RequireSelfie   string              ` json:"requireSelfie"`
	EmailTo         string              ` json:"emailTargets"`
	StartJourney    string              ` json:"startJourney"`
	FinishJourney   string              ` json:"finishJourney"`
	CreatedAt       string              ` json:"createdAt"`
	UpdatedAt       string              ` json:"updatedAt"`
	Sites           []SitesVM           ` json:"sites"`
	Questionnaires  []QuestionnairesVM  ` json:"questionnaires"`
	AssignedAuditor []AssignedAuditorVM ` json:"assignedAuditor"`
	Activity        []ActivityVM        ` json:"activity"`
}

// SitesVM ...
type SitesVM struct {
	SiteID string `json:"siteID"`
}

// QuestionnairesVM ...
type QuestionnairesVM struct {
	QuestionnairesID string `json:"questionnairesID"`
}

// AssignedAuditorVM ...
type AssignedAuditorVM struct {
	UserID string `json:"userID"`
}

// ActivityVM ...
type ActivityVM struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Datetime string `json:"datetime"`
}
