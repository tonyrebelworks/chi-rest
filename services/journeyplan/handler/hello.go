package handler

import (
	"chi-rest/usecase"
	"fmt"
	"net/http"
)

// Hello ...
// GetStringByInt example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string	MsgSuccess
// @Router / [get]
func (h *Contract) Hello(w http.ResponseWriter, r *http.Request) {
	res, err := usecase.UC{h.App}.GetAllJourney()
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}
