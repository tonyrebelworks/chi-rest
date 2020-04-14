package handler

import (
	"chi-rest/services/journeyplan/request"
	"chi-rest/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
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
	res, err := usecase.UC{h.App}.GetAllJourney()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(res)
	h.SendSuccess(w, res, nil)
	return
}

// GetDetailJourney ...
func (h *Contract) GetDetailJourney(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailJourney(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddJourney ...
func (h *Contract) AddJourney(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddJourneyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	// if err = h.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	code := xid.New().String()

	JourneyName := req.JourneyName
	JourneySchedule := req.JourneySchedule
	Salesman := req.Salesman
	Sites := req.Sites
	Questionnaires := req.Questionnaires
	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	EmailTo := req.EmailTo
	Activity := req.Activity
	// StartJourney := req.StartJourney
	// FinishJourney := req.FinishJourney

	// mdl := usecase.UC{h.App}.StoreJourney(code, JourneyName, JourneySchedule, Salesman, Sites, Questionnaires, Signatures, RequireSelfie, EmailTo, Activity, StartJourney, FinishJourney)

	lastID, err := usecase.UC{h.App}.StoreJourney(
		code,
		JourneyName,
		JourneySchedule,
		Salesman,
		Sites,
		Questionnaires,
		Signatures,
		RequireSelfie,
		EmailTo,
		Activity,
		// req.StartJourney,
		// req.FinishJourney,
		// req.CreatedAt,
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
	// if err = h.Handler.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	JourneyName := req.JourneyName
	JourneySchedule := req.JourneySchedule
	Salesman := req.Salesman
	Sites := req.Sites
	Questionnaires := req.Questionnaires
	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	EmailTo := req.EmailTo
	Activity := req.Activity

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateJourney(
		code,
		JourneyName,
		JourneySchedule,
		Salesman,
		Sites,
		Questionnaires,
		Signatures,
		RequireSelfie,
		EmailTo,
		Activity,
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
