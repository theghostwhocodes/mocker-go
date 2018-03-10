package contentmanagers

// MockHTTPParam is the structure for the HTTP query parameters
type MockHTTPParam struct {
	key   string
	value string
}

// MockHTTPHeader is the structure for the HTTP headers
type MockHTTPHeader struct {
	key   string
	value string
}

// MockHTTPRequest is the struct for the request part of the mock content
type MockHTTPRequest struct {
	Method      string
	HTTPParams  []MockHTTPParam
	Payload     string
	HTTPHeaders []MockHTTPHeader
}

// MockHTTPResponse is the struct for the reponse part of the mock content
type MockHTTPResponse struct {
	HTTPHeaders []MockHTTPHeader
	body        interface{}
}

// MockHTTP is the struct for the mock content
type MockHTTP struct {
	request  MockHTTPRequest
	response MockHTTPResponse
}
