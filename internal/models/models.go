package models

// MockHTTPRequest is the struct for the request part of the mock content
type MockHTTPRequest struct {
	Method  string
	Params  map[string]string
	Payload interface{}
	Headers map[string][]string
}

// MockHTTPResponse is the struct for the reponse part of the mock content
type MockHTTPResponse struct {
	Headers map[string]string
	Body    interface{}
}

// MockHTTP is the struct for the mock content
type MockHTTP struct {
	Request  MockHTTPRequest
	Response MockHTTPResponse
}
