package viewmodel

// JourneyPlanVM ...
type JourneyPlanVM struct {
	ID              uint   ` json:"id"`
	Code            string ` json:"code"`
	JourneyName     string ` json:"journeyName"`
	JourneySchedule string ` json:"journeySchedule"`
	Salesman        string ` json:"assignedAuditor"`
	Sites           string ` json:"sites"`
	Questionnaires  string ` json:"questionnaires"`
	Activity        string ` json:"activity"`
	Signatures      string ` json:"signatures"`
	RequireSelfie   string ` json:"requireSelfie"`
	EmailTo         string ` json:"emailTargets"`
	StartJourney    string ` json:"startJourney"`
	FinishJourney   string ` json:"finishJourney"`
	CreatedAt       string ` json:"createdAt"`
	UpdatedAt       string ` json:"updatedAt"`
	AssignedAuditor []AssignedAuditorVM
}

// AssignedAuditorVM ...
type AssignedAuditorVM struct {
	UserID string `json:"userID"`
}
