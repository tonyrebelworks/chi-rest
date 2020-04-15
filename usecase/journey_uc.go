package usecase

import (
	"chi-rest/model"
	"chi-rest/usecase/viewmodel"
	"strings"
	"time"
)

// GetAllJourney ...
func (uc UC) GetAllJourney() ([]map[string]interface{}, error) {
	data, err := model.JourneyOp.GetAll(uc.DB)
	if err != nil {
		return nil, err
	}

	resMap := make([]map[string]interface{}, 0)
	for _, a := range data {
		dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, a.Code)
		if err != nil {
			return nil, err
		}

		sitesRes := make([]map[string]interface{}, 0)
		site := a.Sites
		arrSites := strings.Split(site, "|")
		for i := range arrSites {
			sitesRes = append(sitesRes, map[string]interface{}{
				"siteID": arrSites[i],
			})
		}

		questionnairesRes := make([]map[string]interface{}, 0)
		question := a.Questionnaires
		arrQuestion := strings.Split(question, "|")
		for i := range arrQuestion {
			questionnairesRes = append(questionnairesRes, map[string]interface{}{
				"questionnaireID": arrQuestion[i],
			})
		}

		emailRes := make([]map[string]interface{}, 0)
		email := a.EmailTo
		arrEmail := strings.Split(email, "|")
		for i := range arrEmail {
			emailRes = append(emailRes, map[string]interface{}{
				"email": arrEmail[i],
			})
		}

		assignedAuditorRes := make([]map[string]interface{}, 0)
		assignAud := a.Salesman
		arrAssignAud := strings.Split(assignAud, "|")
		for i := range arrAssignAud {
			assignedAuditorRes = append(assignedAuditorRes, map[string]interface{}{
				"userID": arrAssignAud[i],
			})
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

		resMap = append(resMap, map[string]interface{}{
			"id":                    a.ID,
			"code":                  a.Code,
			"journeyName":           a.JourneyName,
			"journeySchedule":       a.JourneySchedule,
			"activity":              activityRes,
			"signatures":            a.Signatures,
			"requireSelfie":         a.RequireSelfie,
			"person":                a.Person.String,
			"startJourney":          a.StartJourney.String,
			"finishJourney":         a.FinishJourney.String,
			"isDueToday":            true,
			"isDraft":               false,
			"isMakeUp":              false,
			"todayCompletedCount":   0,
			"completedCount":        0,
			"todayScheduleCount":    1,
			"isCompletedToday":      false,
			"isCompletedThisPeriod": false,
			"scheduleCount":         7,
			"isScheduleThisPeriod":  true,
			"createdAt":             a.CreatedAt.String,
			"createdBy":             a.CreatedBy.String,
			"updatedAt":             a.UpdatedAt.String,
			"updatedBy":             a.UpdatedBy.String,
			"sites":                 sitesRes,
			"questionnaires":        questionnairesRes,
			"emailTargets":          emailRes,
			"assignedAuditor":       assignedAuditorRes,
		})
	}

	return resMap, err
}

// GetDetailJourney ...
func (uc UC) GetDetailJourney(code string) (viewmodel.JourneyPlanVM, error) {
	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
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

	res := viewmodel.JourneyPlanVM{
		ID:                    data.ID,
		Code:                  data.Code,
		JourneyName:           data.JourneyName,
		JourneySchedule:       data.JourneySchedule,
		AssignedAuditor:       assignedAuditorRes,
		Sites:                 sitesRes,
		EmailTargets:          emailRes,
		Questionnaires:        questionnairesRes,
		Activity:              activityRes,
		Signatures:            data.Signatures,
		RequireSelfie:         data.RequireSelfie,
		Person:                data.Person.String,
		CreatedAt:             data.CreatedAt.String,
		UpdatedAt:             data.UpdatedAt.String,
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

// StoreJourney ...
func (uc UC) StoreJourney(
	code string,
	journeyName string,
	journeySchedule int64,
	salesman []string,
	sites []string,
	questionnaires []string,
	signatures int64,
	requireSelfie int64,
	person string,
	emailTo []string,

) (int64, error) {

	dt, err := model.JourneyOp.Store(uc.DB, code, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, person, emailTo, time.Now().UTC())
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
