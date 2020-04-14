package viewmodel

// JourneyPlanVM ...
type JourneyPlanVM struct {
	ID                    uint                ` json:"id"`
	Code                  string              ` json:"code"`
	JourneyName           string              ` json:"journeyName"`
	JourneySchedule       string              ` json:"journeySchedule"`
	Signatures            string              ` json:"signatures"`
	RequireSelfie         bool                ` json:"requireSelfie"`
	Person                string              ` json:"person"`
	StartJourney          string              ` json:"startJourney"`
	FinishJourney         string              ` json:"finishJourney"`
	IsDueToday            bool                ` json:"isDueToday"`
	IsDraft               bool                ` json:"isDraft"`
	IsMakeUp              bool                ` json:"isMakeUp"`
	TodayCompletedCount   int                 ` json:"todayCompletedCount"`
	CompletedCount        int                 ` json:"completedCount"`
	TodayScheduleCount    int                 ` json:"todayScheduleCount"`
	IsCompletedToday      bool                ` json:"isCompletedToday"`
	IsCompletedThisPeriod bool                ` json:"isCompletedThisPeriod"`
	ScheduleCount         int                 ` json:"scheduleCount"`
	IsScheduleThisPeriod  bool                ` json:"isScheduleThisPeriod"`
	CreatedAt             string              ` json:"createdAt"`
	CreatedBy             string              ` json:"createdBy"`
	UpdatedAt             string              ` json:"updatedAt"`
	UpdatedBy             string              ` json:"updatedBy"`
	Sites                 []SitesVM           ` json:"sites"`
	Questionnaires        []QuestionnairesVM  ` json:"questionnaires"`
	AssignedAuditor       []AssignedAuditorVM ` json:"assignedAuditor"`
	EmailTargets          []EmailTargetsVM    ` json:"emailTargets"`
	Activity              []ActivityVM        ` json:"activity"`
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

// EmailTargetsVM ...
type EmailTargetsVM struct {
	Email string `json:"email"`
}

// ActivityVM ...
type ActivityVM struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Datetime string `json:"datetime"`
}
