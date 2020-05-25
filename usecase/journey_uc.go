package usecase

import (
	"chi-rest/model"
	"chi-rest/usecase/viewmodel"
	"log"
	"strconv"
	"strings"
	"time"
)

// GetAllJourney ...
func (uc UC) GetAllJourney(types string, maxID, limit int) ([]map[string]interface{}, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.JourneyOp.GetAll(uc.DB, types, maxID, limit)

	if len(data) > 0 {
		firstRecord := data[0]
		firstID := int(firstRecord.ID)
		lastRecord := data[len(data)-1]
		lastID := int(lastRecord.ID)
		pagination = SimplePaginationRes(types, maxID, firstID, lastID, limit)
	}

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]map[string]interface{}, 0)
	for _, a := range data {
		dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, "bq7e2l5hipgeufbrju5g")
		if err != nil {
			return nil, pagination, err
		}

		// sitesRes := make([]map[string]interface{}, 0)
		site := a.Sites
		arrSites := strings.Split(site, "|")
		// for i := range arrSites {
		// 	sitesRes = append(sitesRes, map[string]interface{}{
		// 		"siteID": arrSites[i],
		// 	})
		// }

		question := a.Questionnaires
		arrQuestion := strings.Split(question, "|")

		email := a.EmailTo
		arrEmail := strings.Split(email, "|")

		assignAud := a.Salesman
		arrAssignAud := strings.Split(assignAud, "|")

		var strAssignAud string
		if len(arrAssignAud) > 1 {
			strAssignAud = ""
		} else {
			strAssignAud = strings.Join(arrAssignAud, "|")
		}

		activityRes := []viewmodel.ActivityVM{}
		for _, ac := range dataActivity {
			tempRes := viewmodel.ActivityVM{
				// UserID:   a.UserID,
				Username: ac.Username,
				Datetime: a.CreatedAt.String,
			}
			activityRes = append(activityRes, tempRes)
		}

		datesCustom := a.DatesCustom.String
		daysOfWeek := a.DaysOfWeek.String
		datesOfMonth := a.DatesOfMonth.String

		tmpDC := strings.Split(datesCustom, ",")
		datesCustomToInt := make([]int, 0, len(tmpDC))
		if datesCustom != "" {
			for _, raw := range tmpDC {
				v, err := strconv.Atoi(raw)
				if err != nil {
					// log.Print(err)
					continue
				}
				datesCustomToInt = append(datesCustomToInt, v)
			}
		}

		tmpDow := strings.Split(daysOfWeek, ",")
		daysOfWeekToInt := make([]int, 0, len(tmpDow))
		if daysOfWeek != "" {
			for _, raw := range tmpDow {
				v, err := strconv.Atoi(raw)
				if err != nil {
					// log.Print(err)
					continue
				}
				daysOfWeekToInt = append(daysOfWeekToInt, v)
			}
		}
		tmpDom := strings.Split(datesOfMonth, ",")
		datesOfMonthToInt := make([]int, 0, len(tmpDom))
		if datesOfMonth != "" {
			for _, raw := range tmpDom {
				v, err := strconv.Atoi(raw)
				if err != nil {
					// log.Print(err)
					continue
				}
				datesOfMonthToInt = append(datesOfMonthToInt, v)
			}
		}
		journeySchedule := a.JourneySchedule
		var typeJourneySchedule string

		if journeySchedule == 1 {
			typeJourneySchedule = "daily"
		}
		if journeySchedule == 2 {
			typeJourneySchedule = "weekly"
		}
		if journeySchedule == 3 {
			typeJourneySchedule = "monthly"
		}

		resMap = append(resMap, map[string]interface{}{
			"id":               a.ID,
			"code":             a.Code,
			"journeyName":      a.JourneyName,
			"assignedAuditor":  strAssignAud,
			"auditors":         arrAssignAud,
			"departmentID":     a.DepartmentKey,
			"type":             typeJourneySchedule,
			"sites":            arrSites,
			"questionnaires":   arrQuestion,
			"signatures":       a.Signatures,
			"requireSelfie":    a.RequireSelfie,
			"selfieSignature":  []string{},
			"person":           a.Person.String,
			"emailTargets":     arrEmail,
			"startTimeJourney": a.StartJourney.String,
			"endTimeJourney":   a.FinishJourney.String,
			"datesCustom":      datesCustomToInt,
			"daysOfWeek":       daysOfWeekToInt,
			"datesOfMonth":     datesOfMonthToInt,
			"journeySchedule":  journeySchedule,
			"activity":         activityRes,
			"createdAt":        a.CreatedAt.String,
			"createdBy":        "admin",
			"updatedAt":        a.UpdatedAt.String,
			"updatedBy":        "admin",
		})
	}

	return resMap, pagination, err
}

// GetDetailJourney ...
func (uc UC) GetDetailJourney(code string) (viewmodel.JourneyPlanVM, error) {

	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
	}

	journeySchedule := data.JourneySchedule
	var typeJourneySchedule string

	if journeySchedule == 0 {
		typeJourneySchedule = "daily"
	}
	if journeySchedule == 1 {
		typeJourneySchedule = "weekly"
	}
	if journeySchedule == 2 {
		typeJourneySchedule = "monthly"
	}

	site := data.Sites
	arrSites := strings.Split(site, "|")

	questionnaires := data.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")

	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")

	var strAssignAud string
	if len(arrAssignAud) > 1 {
		strAssignAud = ""
	} else {
		strAssignAud = strings.Join(arrAssignAud, "|")
	}

	email := data.EmailTo
	arrEmail := strings.Split(email, "|")

	dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, "bq7e2l5hipgeufbrju5g")
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
	}

	activityRes := []viewmodel.ActivityVM{}
	for _, ac := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			// UserID:   a.UserID,
			Username: ac.Username,
			Datetime: data.CreatedAt.String,
		}
		activityRes = append(activityRes, tempRes)

	}

	datesCustom := data.DatesCustom.String
	daysOfWeek := data.DaysOfWeek.String
	datesOfMonth := data.DatesOfMonth.String

	tmpDC := strings.Split(datesCustom, ",")
	datesCustomToInt := make([]int, 0, len(tmpDC))
	if datesCustom != "" {
		for _, raw := range tmpDC {
			v, err := strconv.Atoi(raw)
			if err != nil {
				// log.Print(err)
				continue
			}
			datesCustomToInt = append(datesCustomToInt, v)
		}
	}

	tmpDow := strings.Split(daysOfWeek, ",")
	daysOfWeekToInt := make([]int, 0, len(tmpDow))
	if daysOfWeek != "" {
		for _, raw := range tmpDow {
			v, err := strconv.Atoi(raw)
			if err != nil {
				// log.Print(err)
				continue
			}
			daysOfWeekToInt = append(daysOfWeekToInt, v)
		}
	}
	tmpDom := strings.Split(datesOfMonth, ",")
	datesOfMonthToInt := make([]int, 0, len(tmpDom))
	if datesOfMonth != "" {
		for _, raw := range tmpDom {
			v, err := strconv.Atoi(raw)
			if err != nil {
				// log.Print(err)
				continue
			}
			datesOfMonthToInt = append(datesOfMonthToInt, v)
		}
	}

	res := viewmodel.JourneyPlanVM{
		ID:               data.ID,
		Code:             data.Code,
		JourneyName:      data.JourneyName,
		AssignedAuditor:  strAssignAud,
		Auditors:         arrAssignAud,
		DepartmentKey:    data.DepartmentKey,
		Type:             typeJourneySchedule,
		Sites:            arrSites,
		Questionnaires:   arrQuestionnaires,
		Signatures:       data.Signatures,
		RequireSelfie:    data.RequireSelfie,
		SelfieSignature:  []string{},
		Person:           data.Person.String,
		EmailTargets:     arrEmail,
		StartTimeJourney: data.StartJourney.String,
		EndTimeJourney:   data.FinishJourney.String,
		JourneySchedule:  data.JourneySchedule,
		DateCustom:       datesCustomToInt,
		DaysOfWeek:       daysOfWeekToInt,
		DateOfMonth:      datesOfMonthToInt,
		Activity:         activityRes,
		CreatedAt:        data.CreatedAt.String,
		CreatedBy:        "admin",
		UpdatedAt:        data.UpdatedAt.String,
		UpdatedBy:        "admin",
	}

	return res, err
}

// StoreJourney ...
func (uc UC) StoreJourney(
	code string,
	journeyName string,
	departmentKey string,
	journeySchedule int64,
	datesCustom []string,
	daysOfWeek []string,
	datesOfMonth []string,
	salesman []string,
	sites []string,
	questionnaires []string,
	signatures int64,
	requireSelfie int64,
	person string,
	emailTo []string,

) (int64, error) {

	dt, err := model.JourneyOp.Store(uc.DB, code, journeyName, departmentKey, journeySchedule, datesCustom, daysOfWeek, datesOfMonth, salesman, sites, questionnaires, signatures, requireSelfie, person, emailTo, time.Now().UTC())
	return dt, err
}

// UpdateJourney ...
func (uc UC) UpdateJourney(
	code string,
	journeyName string,
	departmentKey string,
	journeySchedule int64,
	datesCustom []string,
	daysOfWeek []string,
	datesOfMonth []string,
	salesman []string,
	sites []string,
	questionnaires []string,
	signatures int64,
	requireSelfie int64,
	person string,
	emailTo []string,
	// activity string,

) (int64, error) {
	dt, err := model.JourneyOp.Update(uc.DB, code, journeyName, departmentKey, journeySchedule, datesCustom, daysOfWeek, datesOfMonth, salesman, sites, questionnaires, signatures, requireSelfie, person, emailTo, time.Now().UTC())
	return dt, err
}

// DeleteJourney ...
func (uc UC) DeleteJourney(code string) ([]*model.JourneyEntity, error) {

	dt, err := model.JourneyOp.DeleteJourney(uc.DB, code, time.Now().UTC())
	return dt, err
}

// UpdateTimeJourney ...
func (uc UC) UpdateTimeJourney(
	JourneyID string,
	StartTime string,
	EndTime string,
) ([]*model.JourneyEntity, error) {

	dt, err := model.JourneyOp.UpdateTimeJourney(uc.DB, JourneyID, StartTime, EndTime, time.Now().UTC())
	return dt, err
}

// GetDetailJourneyMobile ...
func (uc UC) GetDetailJourneyMobile(code string) (viewmodel.JourneyPlanMobileVM, error) {
	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanMobileVM{}, err
	}

	sitesRes := make([]viewmodel.SitesVM, 0)
	site := data.Sites
	arrSites := strings.Split(site, "|")
	for i := range arrSites {
		sitesRes = append(sitesRes, viewmodel.SitesVM{
			SiteID: arrSites[i],
		})
	}

	questionnairesRes := make([]viewmodel.QuestionnairesVM, 0)
	questionnaires := data.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")
	for i := range arrQuestionnaires {
		questionnairesRes = append(questionnairesRes, viewmodel.QuestionnairesVM{
			QuestionnairesID: arrQuestionnaires[i],
		})
	}

	assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	for i := range arrAssignAud {
		assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
			UserID: arrAssignAud[i],
		})
	}

	emailRes := make([]viewmodel.EmailTargetsVM, 0)
	email := data.EmailTo
	arrEmail := strings.Split(email, "|")
	for i := range arrEmail {
		emailRes = append(emailRes, viewmodel.EmailTargetsVM{
			Email: arrEmail[i],
		})
	}

	dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanMobileVM{}, err
	}

	activityRes := []viewmodel.ActivityVM{}
	for _, ac := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			// UserID:   a.UserID,
			Username: ac.Username,
			Datetime: data.CreatedAt.String,
		}
		activityRes = append(activityRes, tempRes)

	}

	res := viewmodel.JourneyPlanMobileVM{
		Code:                  data.Code,
		Name:                  data.JourneyName,
		StartTime:             data.StartJourney.String,
		EndTime:               data.FinishJourney.String,
		Type:                  "basic",
		Schedule:              data.JourneySchedule,
		Language:              "en",
		Signatures:            data.Signatures,
		SelfieSignature:       data.RequireSelfie,
		Person:                data.Person.String,
		Questionnaires:        questionnairesRes,
		Sites:                 sitesRes,
		IsDueToday:            true,
		IsDraft:               false,
		IsMakeUp:              false,
		TodayCompletedCount:   0,
		CompletedCount:        0,
		TodayScheduleCount:    1,
		IsCompletedToday:      false,
		IsCompletedThisPeriod: false,
		ScheduleCount:         7,
		IsScheduleThisPeriod:  true,
	}

	return res, err
}

// GetReportJourney ...
func (uc UC) GetReportJourney(journeyID string) ([]map[string]interface{}, error) {

	data, err := model.ReportFirebaseOp.GetAllReportJourney(uc.DB, journeyID)
	if err != nil {
		return nil, err
	}
	// fmt.Println(journeyID)

	dataJourney, err := model.JourneyOp.GetDetail(uc.DB, journeyID)
	if err != nil {
		return nil, err
	}

	sitesRes := make([]viewmodel.SitesVM, 0)
	site := dataJourney.Sites
	arrSites := strings.Split(site, "|")
	for i := range arrSites {
		sitesRes = append(sitesRes, viewmodel.SitesVM{
			SiteID: arrSites[i],
		})
	}

	questionnairesRes := make([]viewmodel.QuestionnairesVM, 0)
	questionnaires := dataJourney.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")
	for i := range arrQuestionnaires {
		questionnairesRes = append(questionnairesRes, viewmodel.QuestionnairesVM{
			QuestionnairesID: arrQuestionnaires[i],
		})
	}
	assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := dataJourney.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	for i := range arrAssignAud {
		assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
			UserID: arrAssignAud[i],
		})
	}

	datesCustom := dataJourney.DatesCustom.String
	daysOfWeek := dataJourney.DaysOfWeek.String
	datesOfMonth := dataJourney.DatesOfMonth.String

	tmpDC := strings.Split(datesCustom, ",")
	datesCustomToInt := make([]int, 0, len(tmpDC))
	if datesCustom != "" {
		for _, raw := range tmpDC {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesCustomToInt = append(datesCustomToInt, v)
		}
	}

	tmpDow := strings.Split(daysOfWeek, ",")
	daysOfWeekToInt := make([]int, 0, len(tmpDow))
	if daysOfWeek != "" {
		for _, raw := range tmpDow {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			daysOfWeekToInt = append(daysOfWeekToInt, v)
		}
	}
	tmpDom := strings.Split(datesOfMonth, ",")
	datesOfMonthToInt := make([]int, 0, len(tmpDom))
	if datesOfMonth != "" {
		for _, raw := range tmpDom {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesOfMonthToInt = append(datesOfMonthToInt, v)
		}
	}

	resMap := make([]map[string]interface{}, 0)
	for _, rep := range data {
		if err != nil {
			return nil, err
		}

		repURL := rep.URL.String
		arrRepURL := strings.Split(repURL, "|")

		dataTraTi, err := model.TrackingTimeOp.GetByJourneyCode(uc.DB, rep.ReportID)
		if err != nil {
			return nil, err
		}

		traTiRes := []viewmodel.TrackingTimeGPSVM{}
		for _, a := range dataTraTi {
			tempRes := viewmodel.TrackingTimeGPSVM{
				TrackingTime: a.CreatedAt.String,
				Coordinates: viewmodel.CoordinatesVM{
					Lat:  a.Latitude,
					Long: a.Longitude,
				},
			}
			traTiRes = append(traTiRes, tempRes)

		}
		resMap = append(resMap, map[string]interface{}{
			"id":              dataJourney.ID,
			"code":            dataJourney.Code,
			"journeyName":     dataJourney.JourneyName,
			"departmentKey":   dataJourney.DepartmentKey,
			"journeySchedule": dataJourney.JourneySchedule,
			"dateCustom":      datesCustomToInt,
			"daysOfWeek":      daysOfWeekToInt,
			"dateOfMonth":     datesOfMonthToInt,
			"assignedAuditor": assignedAuditorRes,
			"sites":           sitesRes,
			"questionnaires":  questionnairesRes,
			"reports": map[string]interface{}{
				"reportID":   rep.ReportID,
				"userID":     rep.UserID,
				"journeyID":  rep.JourneyID,
				"url":        arrRepURL,
				"start":      rep.Start.String,
				"end":        rep.End.String,
				"reportDate": rep.ReportDate.String,
			},
			"signatures":      dataJourney.Signatures,
			"startJourney":    dataJourney.StartJourney.String,
			"finishJourney":   dataJourney.FinishJourney.String,
			"createdAt":       dataJourney.CreatedAt.String,
			"trackingTimeGPS": traTiRes,

			// ReportID:   rep.ReportID,
			// UserID:     rep.UserID,
			// JourneyID:  rep.JourneyID,
			// URL:        rep.URL.String,
			// Start:      rep.Start.String,
			// End:        rep.End.String,
			// ReportDate: rep.ReportDate.String,
		})
	}

	return resMap, err
}

// DeleteReportByID ...
func (uc UC) DeleteReportByID(code string) ([]*model.ReportFirebaseEntity, error) {

	dt, err := model.ReportFirebaseOp.DeleteReportByID(uc.DB, code, time.Now().UTC())
	return dt, err
}

// AddTrackingTimeJourney ...
func (uc UC) AddTrackingTimeJourney(
	reportJourneyID string,
	userCode string,
	latitude string,
	longitude string,

) (int64, error) {

	dt, err := model.TrackingTimeOp.Store(uc.DB, reportJourneyID, userCode, latitude, longitude, time.Now().UTC())
	return dt, err
}

// GetAllJourneyMobile ...
func (uc UC) GetAllJourneyMobile() ([]viewmodel.GetAllJourneyPlanMobileVM, error) {
	data, err := model.JourneyOp.GetAll(uc.DB, "next", 0, 10000)
	if err != nil {
		return nil, err
	}

	resMap := make([]viewmodel.GetAllJourneyPlanMobileVM, 0)
	for _, a := range data {

		resMap = append(resMap, viewmodel.GetAllJourneyPlanMobileVM{
			Code:                a.Code,
			Name:                a.JourneyName,
			Schedule:            a.JourneySchedule,
			Type:                "basic",
			Priority:            true,
			Language:            "en",
			TodayCompletedCount: 0,
			CompletedCount:      0,
		})
	}

	return resMap, err
}

// GetReportJourneyByParam ...
func (uc UC) GetReportJourneyByParam(userID string, journeyID string, reportDate string) (map[string]interface{}, error) {

	resMap := make(map[string]interface{}, 0)

	data, err := model.ReportFirebaseOp.GetByParam(uc.DB, userID, journeyID, reportDate)
	if err != nil {
		return nil, err
	}

	resMap = map[string]interface{}{
		"reportID":   data.ReportID,
		"userID":     data.UserID,
		"journeyID":  data.JourneyID,
		"url":        data.URL.String,
		"start":      data.Start.String,
		"end":        data.End.String,
		"status":     data.Status,
		"reportDate": data.ReportDate.String,
	}

	return resMap, err
}

// GetReportJourneyByParamStarted ...
func (uc UC) GetReportJourneyByParamStarted(userID string, reportDate string) (map[string]interface{}, error) {

	resMap := make(map[string]interface{}, 0)

	data, err := model.ReportFirebaseOp.GetByParamStarted(uc.DB, userID, reportDate)
	if err != nil {
		return nil, err
	}

	resMap = map[string]interface{}{
		"reportID":   data.ReportID,
		"userID":     data.UserID,
		"journeyID":  data.JourneyID,
		"url":        data.URL.String,
		"start":      data.Start.String,
		"end":        data.End.String,
		"status":     data.Status,
		"reportDate": data.ReportDate.String,
	}

	return resMap, err
}

// GetByReportID ...
func (uc UC) GetByReportID(reportID string) (map[string]interface{}, error) {

	resMap := make(map[string]interface{}, 0)

	data, err := model.ReportFirebaseOp.GetByReportID(uc.DB, reportID)
	if err != nil {
		return nil, err
	}

	resMap = map[string]interface{}{
		"reportID":   data.ReportID,
		"userID":     data.UserID,
		"journeyID":  data.JourneyID,
		"url":        data.URL.String,
		"start":      data.Start.String,
		"end":        data.End.String,
		"status":     data.Status,
		"reportDate": data.ReportDate.String,
	}

	return resMap, err
}

// GetReportJourneyByUJR ...
func (uc UC) GetReportJourneyByUJR(
	userID string,
	journeyID string,
	url string,
	start string,
	end string,
	status int,
	reportDate string,

) ([]map[string]interface{}, error) {

	// if userID != "" && journeyID != "" && reportDate != "" {
	// 	dt, err := model.ReportFirebaseOp.Store(uc.DB, url, journeyID, time.Now().UTC())
	// 	return dt, err
	// }
	resMap := make([]map[string]interface{}, 0)
	data, err := model.ReportFirebaseOp.GetReportID(uc.DB, userID, journeyID, reportDate)

	resMap = append(resMap, map[string]interface{}{
		"reportID":   data.ReportID,
		"userID":     data.UserID,
		"journeyID":  data.JourneyID,
		"url":        data.URL.String,
		"start":      data.Start.String,
		"end":        data.End.String,
		"status":     data.Status,
		"reportDate": data.ReportDate.String,
	})

	return resMap, err

}

// StoreReportJourney ...
func (uc UC) StoreReportJourney(
	code string,
	userID string,
	journeyID string,
	url string,
	start string,
	end string,
	status int,
	reportDate string,

) (int64, error) {

	dt, err := model.ReportFirebaseOp.Store(uc.DB, code, userID, journeyID, url, start, end, status, reportDate, time.Now().UTC())
	return dt, err
}

// UpdateReportJourney ...
func (uc UC) UpdateReportJourney(
	code string,
	url string,
	end string,

) ([]*model.ReportFirebaseEntity, error) {

	dt, err := model.ReportFirebaseOp.Update(uc.DB, code, url, end)
	return dt, err
}

// GetInterval ...
func (uc UC) GetInterval() ([]viewmodel.GetIntervalVM, error) {
	data, err := model.IntervalOp.GetInterval(uc.DB)
	if err != nil {
		return nil, err
	}

	resMap := make([]viewmodel.GetIntervalVM, 0)
	for _, a := range data {

		resMap = append(resMap, viewmodel.GetIntervalVM{
			TimePerSecond: a.TimePerSecond,
		})
	}

	return resMap, err
}

// UpdateInterval ...
func (uc UC) UpdateInterval(
	TimePerSecond int,
) (int64, error) {

	dt, err := model.IntervalOp.UpdateInterval(uc.DB, TimePerSecond)
	return dt, err
}
