package helper

import (
	"encoding/json"
	"net/http"
)

// Response format
type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespondWithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	data := Response{
		Status:  int64(code),
		Message: msg,
		Data:    nil,
	}
	RespondWithJSON(w, code, data)
}
