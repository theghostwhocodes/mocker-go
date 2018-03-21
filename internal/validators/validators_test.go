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
