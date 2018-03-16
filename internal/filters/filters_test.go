package filters

import (
	"testing"

	"github.com/theghostwhocodes/mocker-go/internal/models"
)

func TestCheckArrayEquality(t *testing.T) {
	slice1 := []string{"string1", "string2", "string3"}
	slice2 := []string{"string1", "string2", "string3"}

	equals := checkArrayEquality(slice1, slice2)

	if !equals {
		t.Fail()
	}
}

func TestCheckArrayInequality1(t *testing.T) {
	slice1 := []string{"string1", "string2", "string3"}
	slice2 := []string{"string1", "string2", "string4"}

	equals := checkArrayEquality(slice1, slice2)

	if equals {
		t.Fail()
	}
}

func TestCheckArrayInequality2(t *testing.T) {
	slice1 := []string{"string1", "string2"}
	slice2 := []string{"string1", "string2", "string4"}

	equals := checkArrayEquality(slice1, slice2)

	if equals {
		t.Fail()
	}
}

func TestFilterMockHTTPMethod(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Method: "GET",
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Method: "POST",
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Method: "GET",
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	filtered, err := FilterMockHTTPMethod(mocks, "GET")

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 2 {
		t.Fail()
	}
}
