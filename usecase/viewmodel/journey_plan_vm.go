package viewmodel

// JourneyPlanVM ...
type JourneyPlanVM struct {
	ID                    uint                ` json:"id"`
	Code                  string              ` json:"code"`
	JourneyName           string              ` json:"journeyName"`
	JourneySchedule       int                 ` json:"journeySchedule"`
	DateCustom            []int               ` json:"dateCustom"`
	DaysOfWeek            []int               ` json:"daysOfWeek"`
	DateOfMonth           []int               ` json:"dateOfMonth"`
	Signatures            string              ` json:"signatures"`
	RequireSelfie         bool                ` json:"requireSelfie"`
	Person                string              ` json:"person"`
	StartTime             string              ` json:"startTime"`
	EndTime               string              ` json:"endTime"`
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
	EmailTargets          []string            ` json:"emailTargets"`
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

// JourneyPlanMobileVM ...
type JourneyPlanMobileVM struct {
	// ID                    uint               ` json:"id"`
	Code                  string             ` json:"id"`
	Name                  string             ` json:"name"`
	StartTime             string             ` json:"startTime"`
	EndTime               string             ` json:"endTime"`
	Type                  string             ` json:"type"`
	Schedule              int                ` json:"schedule"`
	Priority              bool               ` json:"priority"`
	Language              string             ` json:"language"`
	Signatures            string             ` json:"signatures"`
	SelfieSignature       bool               ` json:"selfieSignature"`
	Person                string             ` json:"person"`
	Questionnaires        []QuestionnairesVM ` json:"questionnaires"`
	Sites                 []SitesVM          ` json:"site"`
	IsDueToday            bool               ` json:"isDueToday"`
	IsDraft               bool               ` json:"isDraft"`
	IsMakeUp              bool               ` json:"isMakeUp"`
	TodayCompletedCount   int                ` json:"todayCompletedCount"`
	CompletedCount        int                ` json:"completedCount"`
	TodayScheduleCount    int                ` json:"todayScheduleCount"`
	IsCompletedToday      bool               ` json:"isCompletedToday"`
	IsCompletedThisPeriod bool               ` json:"isCompletedThisPeriod"`
	ScheduleCount         int                ` json:"scheduleCount"`
	IsScheduleThisPeriod  bool               ` json:"isScheduleThisPeriod"`
	// CreatedAt             string             ` json:"createdAt"`
	// CreatedBy             string             ` json:"createdBy"`
	// UpdatedAt             string             ` json:"updatedAt"`
	// UpdatedBy             string             ` json:"updatedBy"`
}

// ReportJourneyPlanVM ...
type ReportJourneyPlanVM struct {
	ID              uint                ` json:"id"`
	Code            string              ` json:"code"`
	JourneyName     string              ` json:"journeyName"`
	JourneySchedule int                 ` json:"journeySchedule"`
	AssignedAuditor []AssignedAuditorVM ` json:"assignedAuditor"`
	Sites           []SitesVM           ` json:"sites"`
	Questionnaires  []QuestionnairesVM  ` json:"questionnaires"`
	// EmailTargets    []EmailTargetsVM    ` json:"emailTargets"`
	// Activity        []ActivityVM        ` json:"activity"`
	// Reports         []ReportsVM         ` json:"reports"`
	Reports         []ReportsVM         ` json:"reports"`
	TrackingTimeGPS []TrackingTimeGPSVM ` json:"trackingTimeGPS"`
	Signatures      string              ` json:"signatures"`
	StartJourney    string              ` json:"startJourney"`
	FinishJourney   string              ` json:"finishJourney"`
	CreatedAt       string              ` json:"createdAt"`
}

// TrackingTimeGPSVM ...
type TrackingTimeGPSVM struct {
	TrackingTime string `json:"trackingTime"`
	Lat          string `json:"lat"`
	Long         string `json:"long"`
}

// GetAllJourneyPlanMobileVM ...
type GetAllJourneyPlanMobileVM struct {
	Code     string ` json:"id"`
	Name     string ` json:"name"`
	Type     string ` json:"type"`
	Schedule int    ` json:"schedule"`
	Priority bool   ` json:"priority"`
	Language string ` json:"language"`
	// IsDueToday            bool   ` json:"isDueToday"`
	// IsDraft               bool   ` json:"isDraft"`
	// IsMakeUp              bool   ` json:"isMakeUp"`
	TodayCompletedCount int ` json:"todayCompletedCount"`
	CompletedCount      int ` json:"completedCount"`
	// TodayScheduleCount    int    ` json:"todayScheduleCount"`
	// IsCompletedToday      bool   ` json:"isCompletedToday"`
	// IsCompletedThisPeriod bool   ` json:"isCompletedThisPeriod"`
	// ScheduleCount         int    ` json:"scheduleCount"`
	// IsScheduleThisPeriod  bool   ` json:"isScheduleThisPeriod"`
}

// ReportsVM ...
type ReportsVM struct {
	URL string `json:"url"`
}
