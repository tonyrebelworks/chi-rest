package request

// AddJourneyRequest ...
type AddJourneyRequest struct {
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
