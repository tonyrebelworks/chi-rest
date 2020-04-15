package request

// AddJourneyRequest ...
type AddJourneyRequest struct {
	JourneyName     string            `json:"journeyName" validate:"required"`
	JourneySchedule int64             `json:"journeySchedule" validate:"required"`
	AssignedAuditor []AssignedAuditor `json:"assignedAuditor" validate:"required"`
	Sites           []Sites           `json:"sites" validate:"required"`
	Questionnaires  []Questionnaires  `json:"questionnaires" validate:"required"`
	Signatures      int64             `json:"signatures" validate:"required"`
	RequireSelfie   int64             `json:"requireSelfie" validate:"required"`
	EmailTo         []EmailTo         `json:"emailTo" validate:"required"`
	Person          string            `json:"person"`
	StartJourney    string            `json:"startJourney"`
	FinishJourney   string            `json:"finishJourney"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
}

// UpdateJourneyRequest ...
type UpdateJourneyRequest struct {
	JourneyName     string `json:"journeyName" validate:"required"`
	JourneySchedule int64  `json:"journeySchedule" validate:"required"`
	Salesman        string `json:"salesman" validate:"required"`
	Sites           string `json:"sites" validate:"required"`
	Questionnaires  string `json:"questionnaires" validate:"required"`
	Signatures      int64  `json:"signatures" validate:"required"`
	RequireSelfie   int64  `json:"requireSelfie" validate:"required"`
	EmailTo         string `json:"emailTo" validate:"required"`
	Activity        string `json:"activity" validate:"required"`
	StartJourney    string `json:"startJourney"`
	FinishJourney   string `json:"finishJourney"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// AssignedAuditor ...
type AssignedAuditor struct {
	UserID string `json:"userID" `
}

// Sites ...
type Sites struct {
	SiteID string `json:"siteID" `
}

// Questionnaires ...
type Questionnaires struct {
	QuestionnaireID string `json:"questionnaireID" `
}

// EmailTo ...
type EmailTo struct {
	Email string `json:"email" `
}

// UpdateTimeJourneyRequest ...
type UpdateTimeJourneyRequest struct {
	JourneyID string `json:"journeyID" validate:"required"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
