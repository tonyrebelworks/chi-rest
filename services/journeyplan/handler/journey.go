package handler

import (
	"chi-rest/services/journeyplan/request"
	"chi-rest/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	validator "gopkg.in/go-playground/validator.v9"
)

// GetAllJourney ...
// GetStringByInt example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string	MsgSuccess
// @Router / [get]
func (h *Contract) GetAllJourney(w http.ResponseWriter, r *http.Request) {
	var (
		types string
		maxID int
		limit int
		err   error
	)

	types = "next"
	maxID = 0
	limit = 10000

	// types = r.URL.Query().Get("types")
	// if types != "prev" && types != "next" {
	// 	h.SendBadRequest(w, "Invalid type value")
	// 	return
	// }
	// maxID, err = strconv.Atoi(r.URL.Query().Get("max_id"))
	// if err != nil {
	// 	h.SendBadRequest(w, "Invalid last id value")
	// 	return
	// }
	// limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
	// if err != nil {
	// 	h.SendBadRequest(w, "Invalid limit value")
	// 	return
	// }

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllJourney(types, maxID, limit)

	h.SendSuccess(w, res, pagination)
	return
}

// GetDetailJourney ...
func (h *Contract) GetDetailJourney(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailJourney(code)
	if err != nil {
		fmt.Println(err)
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddJourney ...
func (h *Contract) AddJourney(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddJourneyRequest{}
	// reqActivity := request.AddJourneyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := xid.New().String()

	JourneyName := req.JourneyName
	DepartmentKey := req.DepartmentKey
	JourneySchedule := req.JourneySchedule

	if len(req.AssignedAuditor) > 0 {

	}
	assignedAuditors := make([]string, 0)
	for _, aa := range req.AssignedAuditor {
		assignedAuditors = append(assignedAuditors, aa.UserID)
	}

	if len(req.Sites) > 0 {

	}
	sitess := make([]string, 0)
	for _, si := range req.Sites {
		sitess = append(sitess, si.SiteID)
	}

	if len(req.Questionnaires) > 0 {

	}
	questionnairess := make([]string, 0)
	for _, qu := range req.Questionnaires {
		questionnairess = append(questionnairess, qu.QuestionnaireID)
	}

	if len(req.EmailTo) > 0 {

	}
	emails := make([]string, 0)
	for _, em := range req.EmailTo {
		emails = append(emails, em.Email)
	}

	if len(req.DatesCustom) > 0 {

	}
	datesCustom := make([]string, 0)
	for _, dc := range req.DatesCustom {
		datesCustom = append(datesCustom, dc.DatesCustom)
	}

	if len(req.DaysOfWeek) > 0 {

	}
	daysOfWeek := make([]string, 0)
	for _, dow := range req.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, dow.DaysOfWeek)
	}

	if len(req.DatesOfMonth) > 0 {

	}
	datesOfMonth := make([]string, 0)
	for _, dom := range req.DatesOfMonth {
		datesOfMonth = append(datesOfMonth, dom.DateOfMonth)
	}

	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	Person := req.Person

	lastID, err := usecase.UC{h.App}.StoreJourney(
		code,
		JourneyName,
		DepartmentKey,
		JourneySchedule,
		datesCustom,
		daysOfWeek,
		datesOfMonth,
		assignedAuditors,
		sitess,
		questionnairess,
		Signatures,
		RequireSelfie,
		Person,
		emails,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// UpdateJourney ...
func (h *Contract) UpdateJourney(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateJourneyRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	JourneyName := req.JourneyName
	DepartmentKey := req.DepartmentKey
	JourneySchedule := req.JourneySchedule

	assignedAuditors := make([]string, 0)
	for _, aa := range req.AssignedAuditor {
		assignedAuditors = append(assignedAuditors, aa.UserID)
	}

	if len(req.Sites) > 0 {

	}
	sitess := make([]string, 0)
	for _, si := range req.Sites {
		sitess = append(sitess, si.SiteID)
	}

	if len(req.Questionnaires) > 0 {

	}
	questionnairess := make([]string, 0)
	for _, qu := range req.Questionnaires {
		questionnairess = append(questionnairess, qu.QuestionnaireID)
	}

	if len(req.EmailTo) > 0 {

	}
	emails := make([]string, 0)
	for _, em := range req.EmailTo {
		emails = append(emails, em.Email)
	}

	if len(req.DatesCustom) > 0 {

	}
	datesCustom := make([]string, 0)
	for _, dc := range req.DatesCustom {
		datesCustom = append(datesCustom, dc.DatesCustom)
	}

	if len(req.DaysOfWeek) > 0 {

	}
	daysOfWeek := make([]string, 0)
	for _, dow := range req.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, dow.DaysOfWeek)
	}

	if len(req.DatesOfMonth) > 0 {

	}
	datesOfMonth := make([]string, 0)
	for _, dom := range req.DatesOfMonth {
		datesOfMonth = append(datesOfMonth, dom.DateOfMonth)
	}
	// Activity := req.Activity

	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	Person := req.Person

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateJourney(
		code,
		JourneyName,
		DepartmentKey,
		JourneySchedule,
		datesCustom,
		daysOfWeek,
		datesOfMonth,
		assignedAuditors,
		sitess,
		questionnairess,
		Signatures,
		RequireSelfie,
		Person,
		emails,
		// Activity,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// DeleteJourney ...
func (h *Contract) DeleteJourney(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteJourney(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// UpdateTimeJourney ...
func (h *Contract) UpdateTimeJourney(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateTimeJourneyRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	JourneyID := req.JourneyID
	StartTime := req.StartTimeJourney
	EndTime := req.EndTimeJourney

	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateTimeJourney(
		JourneyID,
		StartTime,
		EndTime)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// GetDetailJourneyMobile ...
func (h *Contract) GetDetailJourneyMobile(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailJourneyMobile(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetReportJourney ...
func (h *Contract) GetReportJourney(w http.ResponseWriter, r *http.Request) {
	journeyID := chi.URLParam(r, "journeyid")
	res, err := usecase.UC{h.App}.GetReportJourney(journeyID)
	if err != nil {
		fmt.Println(err)
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddTrackingTimeJourney ...
func (h *Contract) AddTrackingTimeJourney(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddTrackingTimeJourneyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	ReportJourneyID := req.ReportJourneyID
	Latitude := req.Latitude
	Longitude := req.Longitude
	UserCode := ""

	lastID, err := usecase.UC{h.App}.AddTrackingTimeJourney(
		ReportJourneyID,
		UserCode,
		Latitude,
		Longitude,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// GetAllJourneyMobile ...
func (h *Contract) GetAllJourneyMobile(w http.ResponseWriter, r *http.Request) {
	res, err := usecase.UC{h.App}.GetAllJourneyMobile()
	if err != nil {
		fmt.Println(err)
		return
	}
	h.SendSuccess(w, res, nil)
	return
}

// GetURLFirebase ...
func (h *Contract) GetURLFirebase(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userid")
	journeyID := chi.URLParam(r, "journeyid")
	reporDate := chi.URLParam(r, "reportdate")

	res, err := usecase.UC{h.App}.GetReportJourneyByParam(userID, journeyID, reporDate)
	if err != nil {
		fmt.Println(err)
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetURLFirebaseStarted ...
func (h *Contract) GetURLFirebaseStarted(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userid")
	reportDate := chi.URLParam(r, "reportdate")

	res, err := usecase.UC{h.App}.GetReportJourneyByParamStarted(userID, reportDate)
	if err != nil {
		fmt.Println(err)
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetByReportID ...
func (h *Contract) GetByReportID(w http.ResponseWriter, r *http.Request) {
	reportID := chi.URLParam(r, "reportid")

	res, err := usecase.UC{h.App}.GetByReportID(reportID)
	if err != nil {
		fmt.Println(err)
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddURLFirebase ...
func (h *Contract) AddURLFirebase(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddURLFirebaseRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	Code := req.ReportID
	UserID := req.UserID
	JourneyID := req.JourneyID
	URL := req.URL
	Start := req.Start
	End := req.End
	Status := req.Status
	ReportDate := req.ReportDate

	if Start == "" {
		fmt.Println("update")
		res, err := usecase.UC{h.App}.UpdateReportJourney(
			Code, URL, End)
		if err != nil {
			h.SendBadRequest(w, err.Error())
			return
		}
		h.SendSuccess(w, res, nil)

	} else {
		res, err := usecase.UC{h.App}.GetReportJourneyByUJR(UserID, JourneyID, URL, Start, End, Status, ReportDate)
		if err != nil {
			// fmt.Println(err)
			fmt.Println("insert new row")

			code := xid.New().String()

			lastID, err := usecase.UC{h.App}.StoreReportJourney(
				code,
				UserID,
				JourneyID,
				URL,
				Start,
				End,
				Status,
				ReportDate,
			)
			if err != nil {
				h.SendBadRequest(w, err.Error())
				return
			}

			h.SendSuccess(w, map[string]interface{}{}, lastID)
			return
		}

		h.SendSuccess(w, res, nil)
	}

	return
}

//GetInterval ...
func (h *Contract) GetInterval(w http.ResponseWriter, r *http.Request) {
	res, err := usecase.UC{h.App}.GetInterval()
	if err != nil {
		fmt.Println(err)
		return
	}
	h.SendSuccess(w, res, nil)
	return
}

// DeleteReportByID ...
func (h *Contract) DeleteReportByID(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteReportByID(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// UpdateInterval ...
func (h *Contract) UpdateInterval(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateInterval{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	TimePerSecond := req.TimePerSecond

	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateInterval(
		TimePerSecond)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}
