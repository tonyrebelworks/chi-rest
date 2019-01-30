package handler

import (
	"chi-rest/lib"
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
)

// Handler ...
type Handler struct {
	DB  *mgo.Session
	Cfg lib.Config
}

func emptyJSONArr() []map[string]interface{} {
	return []map[string]interface{}{}
}

func sendSuccess(w http.ResponseWriter, payload interface{}) {
	RespondWithJSON(w, 200, 200, "Success", payload)
}

func SendBadRequest(w http.ResponseWriter, message string) {
	RespondWithJSON(w, 400, 400, message, emptyJSONArr())
}

func sendValidationError(w http.ResponseWriter, payload interface{}) {
	RespondWithJSON(w, 400, 405, "validation error", payload)
}

// RespondWithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, httpCode int, statCode int, message string, payload interface{}) {
	respPayload := map[string]interface{}{
		"statCode": statCode,
		"statMsg":  message,
		"data":     payload,
	}

	response, _ := json.Marshal(respPayload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}
