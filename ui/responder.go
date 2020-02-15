package ui

import (
	"encoding/json"
	"net/http"
)

// Respond performs a http response with json encoding of payload if provided
func Respond(w http.ResponseWriter, res Response, origin string) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(res.HTTPStatus)
	if res.Payload != nil {
		json.NewEncoder(w).Encode(res.Payload)
	}
	w.Write([]byte(res.Message))
}

// RespondWithOptions responds with access control allow POST, OPTIONS methods in response headers
func RespondWithOptions(w http.ResponseWriter, res Response, origin string) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(res.HTTPStatus)
	if res.Payload != nil {
		json.NewEncoder(w).Encode(res.Payload)
	}
	w.Write([]byte(res.Message))
}

// CreateResponse creates a Response object that can be used to write http responses
func CreateResponse(status int, message string, payload interface{}) Response {
	r := Response{HTTPStatus: status, Message: message, Payload: payload}
	return r
}

// MakeErrorResponse creates an error response object with the message provided
func MakeErrorResponse(status int, message string) Response {
	r := Response{HTTPStatus: http.StatusInternalServerError, Message: message, Payload: nil}
	return r
}

// RespondError responsds to http request with an error response
func RespondError(w http.ResponseWriter, errorcode int, message string) {
	r := Response{HTTPStatus: errorcode, Message: message, Payload: nil}
	Respond(w, r, "")
}
