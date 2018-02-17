package validators

import "encoding/json"

// IsValidJSON return true if content is valid JSON, else false
func IsValidJSON(content []byte) bool {
	var result map[string]interface{}
	return json.Unmarshal(content, &result) == nil
}

// HasRequest return true if 'request' key is in JSON, else false
func HasRequest(content map[string]interface{}) bool {
	_, ok := content["request"]
	return ok
}

// HasResponse return true if 'response' key is in JSON, else false
func HasResponse(content map[string]interface{}) bool {
	_, ok := content["response"]
	return ok
}
