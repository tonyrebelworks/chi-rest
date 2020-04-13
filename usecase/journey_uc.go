package usecase

import (
	"chi-rest/model"
	"time"
)

// GetAllJourney ...
func (uc UC) GetAllJourney() ([]*model.Journey, error) {
	dt, err := model.JourneyOp.GetAll(uc.DB)

	// Res := model.JourneyVM{}
	// for _, a := range dt {
	// 	tempRes := model.SalesmanEntity{
	// 		UserID: a.Salesman,
	// 	}

	// 	Res = append(Res, tempRes)
	// }
	return dt, err
}

// GetDetailJourney ...
func (uc UC) GetDetailJourney(code string) ([]*model.Journey, error) {

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
func (uc UC) DeleteJourney(code string) ([]*model.Journey, error) {

	dt, err := model.JourneyOp.DeleteJourney(uc.DB, code, time.Now().UTC())
	return dt, err
}
