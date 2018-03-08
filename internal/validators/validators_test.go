package validators

import (
	"testing"
)

func TestIsValidJSON(t *testing.T) {
	content := "{\"key\": \"value\"}"
	isValid := IsValidJSON([]byte(content))

	if !isValid {
		t.Fail()
	}
}

func TestIsNotValidJSON(t *testing.T) {
	content := "{\"key\": \"value}"
	notValid := IsValidJSON([]byte(content))

	if notValid {
		t.Fail()
	}
}

func TestHasRequest(t *testing.T) {
	content := map[string]interface{}{
		"key":     "value",
		"request": "another value",
	}
	hasRequest := HasRequest(content)

	if !hasRequest {
		t.Fail()
	}
}

func TestHasNotRequest(t *testing.T) {
	content := map[string]interface{}{
		"key": "value",
		"req": "another value",
	}
	hasNotRequest := HasRequest(content)

	if hasNotRequest {
		t.Fail()
	}
}

func TestHasResponse(t *testing.T) {
	content := map[string]interface{}{
		"key":      "value",
		"response": "another value",
	}
	hasResponse := HasResponse(content)

	if !hasResponse {
		t.Fail()
	}
}

func TestHasNotResponse(t *testing.T) {
	content := map[string]interface{}{
		"key": "value",
		"req": "another value",
	}
	hasNotResponse := HasResponse(content)

	if hasNotResponse {
		t.Fail()
	}
}
