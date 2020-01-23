package ui

// Response is a container for contents of a template http response
type Response struct {
	HTTPStatus int
	Message    string
	Payload    interface{}
}
