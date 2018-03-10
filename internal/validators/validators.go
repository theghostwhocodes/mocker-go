package validators

import (
	"encoding/json"
)

// IsValidJSON return true if content is valid JSON, else false
func IsValidJSON(content []byte) bool {
	var result map[string]interface{}
	return json.Unmarshal(content, &result) == nil
}
