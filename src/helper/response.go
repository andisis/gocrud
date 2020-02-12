package helper

import (
	"encoding/json"
	"net/http"
)

// ResponseSuccess format
type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

// ResponseError format
type ResponseError struct {
	Error interface{} `json:"error"`
}

// ErrPayload struct
type ErrPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSONResponse write json response format
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
