package request

// AddJourneyRequest ...
type AddJourneyRequest struct {
	JourneyName     string            `json:"journeyName" validate:"required"`
	DepartmentKey   string            `json:"departmentID" validate:"required"`
	JourneySchedule int64             `json:"journeySchedule" validate:"required"`
	AssignedAuditor []AssignedAuditor `json:"assignedAuditor" validate:"required"`
	Sites           []Sites           `json:"sites" validate:"required"`
	Questionnaires  []Questionnaires  `json:"questionnaires" validate:"required"`
	Signatures      int64             `json:"signatures" validate:"required"`
	RequireSelfie   int64             `json:"requireSelfie" validate:"required"`
	DatesCustom     []DatesCustom     `json:"datesCustom" `
	DaysOfWeek      []DaysOfWeek      `json:"daysOfWeek" `
	DatesOfMonth    []DatesOfMonth    `json:"datesOfMonth"`
	EmailTo         []EmailTo         `json:"emailTo"`
	Person          string            `json:"person"`
	StartJourney    string            `json:"startJourney"`
	FinishJourney   string            `json:"finishJourney"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
}

// UpdateJourneyRequest ...
type UpdateJourneyRequest struct {
	JourneyName     string            `json:"journeyName" validate:"required"`
	DepartmentKey   string            `json:"departmentID" validate:"required"`
	JourneySchedule int64             `json:"journeySchedule" validate:"required"`
	AssignedAuditor []AssignedAuditor `json:"assignedAuditor" validate:"required"`
	Sites           []Sites           `json:"sites" validate:"required"`
	Questionnaires  []Questionnaires  `json:"questionnaires" validate:"required"`
	Signatures      int64             `json:"signatures" validate:"required"`
	RequireSelfie   int64             `json:"requireSelfie" validate:"required"`
	Person          string            `json:"person"`
	DatesCustom     []DatesCustom     `json:"datesCustom" `
	DaysOfWeek      []DaysOfWeek      `json:"daysOfWeek" `
	DatesOfMonth    []DatesOfMonth    `json:"datesOfMonth"`
	EmailTo         []EmailTo         `json:"emailTo"`
	Activity        []string          `json:"activity"`
	StartJourney    string            `json:"startJourney"`
	FinishJourney   string            `json:"finishJourney"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
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
	JourneyID        string `json:"journeyID" validate:"required"`
	StartTimeJourney string `json:"startTimeJourney"`
	EndTimeJourney   string `json:"endTimeJourney"`
}

// AddTrackingTimeJourneyRequest ...
type AddTrackingTimeJourneyRequest struct {
	JourneyID string `json:"journeyID" validate:"required"`
	Latitude  string `json:"latitude" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
}

// AddURLFirebaseRequest ...
type AddURLFirebaseRequest struct {
	JourneyID string `json:"journeyID" validate:"required"`
	URL       string `json:"url" validate:"required"`
}

// DatesCustom ...
type DatesCustom struct {
	DatesCustom string `json:"dateCustom" `
}

// DaysOfWeek ...
type DaysOfWeek struct {
	DaysOfWeek string `json:"daysOfWeek" `
}

// DatesOfMonth ...
type DatesOfMonth struct {
	DateOfMonth string `json:"datesOfMonth" `
}

// UpdateInterval ...
type UpdateInterval struct {
	TimePerSecond int `json:"timePerSecond" validate:"required"`
}
