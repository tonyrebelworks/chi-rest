package handler

import (
	"chi-rest/server/request"
	"chi-rest/usecase"
	"encoding/json"
	"net/http"
)

// MemberHandler ...
type MemberHandler struct {
	Handler
}

// Profile Get member profile.
func (h *MemberHandler) Profile(w http.ResponseWriter, r *http.Request) {
	uc := usecase.NewMemberUsecase(h.DB, h.Cfg)
	me, err := uc.FindByID(r.URL.Query().Get("s"))
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	sendSuccess(w, me)

	return
}

// RegisterHandler register action in member controller
func (h *MemberHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.RegisterRequest{}
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	memberUc := usecase.NewMemberUsecase(h.DB, h.Cfg)
	err = memberUc.Register(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	sendSuccess(w, req)
	return
}
