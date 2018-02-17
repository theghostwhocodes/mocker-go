package validators_test

import (
	"testing"

	"github.com/theghostwhocodes/mocker-go/internal/validators"
)

func TestIsValidJSON(t *testing.T) {
	content := "{\"key\": \"value\"}"
	isValid := validators.IsValidJSON([]byte(content))

	if !isValid {
		t.Fail()
	}
}

func TestIsNotValidJSON(t *testing.T) {
	content := "{\"key\": \"value}"
	notValid := validators.IsValidJSON([]byte(content))

	if notValid {
		t.Fail()
	}
}

func TestHasRequest(t *testing.T) {
	content := map[string]interface{}{
		"key":     "value",
		"request": "another value",
	}
	hasRequest := validators.HasRequest(content)

	if !hasRequest {
		t.Fail()
	}
}

func TestHasNotRequest(t *testing.T) {
	content := map[string]interface{}{
		"key": "value",
		"req": "another value",
	}
	hasNotRequest := validators.HasRequest(content)

	if hasNotRequest {
		t.Fail()
	}
}

func TestHasResponse(t *testing.T) {
	content := map[string]interface{}{
		"key":      "value",
		"response": "another value",
	}
	hasResponse := validators.HasResponse(content)

	if !hasResponse {
		t.Fail()
	}
}

func TestHasNotResponse(t *testing.T) {
	content := map[string]interface{}{
		"key": "value",
		"req": "another value",
	}
	hasNotResponse := validators.HasResponse(content)

	if hasNotResponse {
		t.Fail()
	}
}
