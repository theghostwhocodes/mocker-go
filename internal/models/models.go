package models

import (
	"net/http"
	"net/url"
)

// MockHTTPRequest is the struct for the request part of the mock content
type MockHTTPRequest struct {
	Method  string
	Params  url.Values
	Payload url.Values
	Headers http.Header
}

// MockHTTPResponse is the struct for the reponse part of the mock content
type MockHTTPResponse struct {
	Headers http.Header
	Body    interface{}
}

// MockHTTP is the struct for the mock content
type MockHTTP struct {
	Request  MockHTTPRequest
	Response MockHTTPResponse
}
