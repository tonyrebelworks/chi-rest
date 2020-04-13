package usecase

import (
	"chi-rest/model"
	"strings"
	"time"
)

// GetAllJourney ...
func (uc UC) GetAllJourney() ([]map[string]interface{}, error) {
	data, err := model.JourneyOp.GetAll(uc.DB)
	if err != nil {
		return nil, err
	}

	// fmt.Println(data)
	resMap := make([]map[string]interface{}, 0)
	for _, a := range data {
		assignedAuditorRes := make([]map[string]interface{}, 0)

		assignAud := a.Salesman
		arrAssignAud := strings.Split(assignAud, "|")
		for i := range arrAssignAud {
			assignedAuditorRes = append(assignedAuditorRes, map[string]interface{}{
				"userID": arrAssignAud[i],
			})
		}

		resMap = append(resMap, map[string]interface{}{
			"id":              a.ID,
			"code":            a.Code,
			"journeyName":     a.JourneyName,
			"journeySchedule": a.JourneySchedule,
			"sites":           a.Sites,
			"questionnaires":  a.Sites,
			"activity":        a.Activity,
			"signatures":      a.Signatures,
			"requireSelfie":   a.RequireSelfie,
			"emailTo":         a.EmailTo,
			"startJourney":    a.StartJourney.String,
			"finishJourney":   a.FinishJourney.String,
			"createdAt":       a.CreatedAt.String,
			"updatedAt":       a.UpdatedAt.String,
			"assignedAuditor": assignedAuditorRes,
		})
	}

	return resMap, err
}

// GetDetailJourney ...
func (uc UC) GetDetailJourney(code string) ([]*model.JourneyEntity, error) {

	dt, err := model.JourneyOp.GetDetail(uc.DB, code)
	return dt, err
}

// StoreJourney ...
func (uc UC) StoreJourney(
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
	dt, err := model.JourneyOp.Store(uc.DB, code, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, time.Now().UTC())
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
