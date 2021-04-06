package ui

import (
	"encoding/json"
	"net/http"
)

// Respond performs a http response with json encoding of payload if provided
func Respond(w http.ResponseWriter, res Response) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(res.HTTPStatus)
	if res.Payload != nil {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(res.Payload)
	}
	if len(res.Message) > 0 {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(res.Message)
	}
}

// CreateResponse creates a Response object that can be used to write http responses
func CreateResponse(status int, payload interface{}) Response {
	r := Response{HTTPStatus: status, Payload: payload}
	return r
}

// RespondError responsds to http request with an error response
func RespondError(w http.ResponseWriter, errorcode int, message string) {
	http.Error(w, message, errorcode)
}
