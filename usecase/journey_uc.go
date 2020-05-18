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
		dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, a.Code)
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
		for _, a := range dataActivity {
			tempRes := viewmodel.ActivityVM{
				UserID:   a.UserID,
				Username: a.Username,
				Datetime: a.Datetime.String,
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
			"departmentKey":    "",
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
			"createdBy":        a.CreatedBy.String,
			"updatedAt":        a.UpdatedAt.String,
			"updatedBy":        a.UpdatedBy.String,
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

	if journeySchedule == 1 {
		typeJourneySchedule = "daily"
	}
	if journeySchedule == 2 {
		typeJourneySchedule = "weekly"
	}
	if journeySchedule == 3 {
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

	dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
	}

	activityRes := []viewmodel.ActivityVM{}
	for _, a := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			UserID:   a.UserID,
			Username: a.Username,
			Datetime: a.Datetime.String,
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

	res := viewmodel.JourneyPlanVM{
		ID:               data.ID,
		Code:             data.Code,
		JourneyName:      data.JourneyName,
		AssignedAuditor:  strAssignAud,
		Auditors:         arrAssignAud,
		DepartmentKey:    "",
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
		CreatedBy:        "",
		UpdatedAt:        data.UpdatedAt.String,
		UpdatedBy:        "",
	}

	return res, err
}

// StoreJourney ...
func (uc UC) StoreJourney(
	code string,
	journeyName string,
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

	dt, err := model.JourneyOp.Store(uc.DB, code, journeyName, journeySchedule, datesCustom, daysOfWeek, datesOfMonth, salesman, sites, questionnaires, signatures, requireSelfie, person, emailTo, time.Now().UTC())
	return dt, err
}

// UpdateJourney ...
func (uc UC) UpdateJourney(
	code string,
	journeyName string,
	journeySchedule int64,
	salesman string,
	sites string,
	questionnaires string,
	signatures int64,
	requireSelfie int64,
	emailTo string,
	activity string,

) (int64, error) {
	dt, err := model.JourneyOp.Update(uc.DB, code, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, time.Now().UTC())
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
	for _, a := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			UserID:   a.UserID,
			Username: a.Username,
			Datetime: a.Datetime.String,
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
func (uc UC) GetReportJourney(code string) (viewmodel.ReportJourneyPlanVM, error) {
	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
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

	dataRep, err := model.ReportFirebaseOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
	}

	reportsRes := []viewmodel.ReportsVM{}
	for _, a := range dataRep {
		tempRes := viewmodel.ReportsVM{
			URL: a.URL,
		}
		reportsRes = append(reportsRes, tempRes)

	}

	assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	for i := range arrAssignAud {
		assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
			UserID: arrAssignAud[i],
		})
	}

	dataTraTi, err := model.TrackingTimeOp.GetByJourneyCode(uc.DB, code, "")
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
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

	datesCustom := data.DatesCustom.String
	daysOfWeek := data.DaysOfWeek.String
	datesOfMonth := data.DatesOfMonth.String

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

	res := viewmodel.ReportJourneyPlanVM{
		ID:              data.ID,
		Code:            data.Code,
		JourneyName:     data.JourneyName,
		JourneySchedule: data.JourneySchedule,
		DateCustom:      datesCustomToInt,
		DaysOfWeek:      daysOfWeekToInt,
		DateOfMonth:     datesOfMonthToInt,
		AssignedAuditor: assignedAuditorRes,
		Sites:           sitesRes,
		Questionnaires:  questionnairesRes,
		Reports:         reportsRes,
		Signatures:      data.Signatures,
		StartJourney:    data.StartJourney.String,
		FinishJourney:   data.FinishJourney.String,
		CreatedAt:       data.CreatedAt.String,
		TrackingTimeGPS: traTiRes,
	}

	return res, err
}

// AddTrackingTimeJourney ...
func (uc UC) AddTrackingTimeJourney(
	journeyCode string,
	userCode string,
	latitude string,
	longitude string,

) (int64, error) {

	dt, err := model.TrackingTimeOp.Store(uc.DB, journeyCode, userCode, latitude, longitude, time.Now().UTC())
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

// AddURLFirebase ...
func (uc UC) AddURLFirebase(
	url string,
	journeyID string,

) (int64, error) {

	dt, err := model.ReportFirebaseOp.Store(uc.DB, url, journeyID, time.Now().UTC())
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
