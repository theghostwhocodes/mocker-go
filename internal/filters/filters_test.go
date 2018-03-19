package filters

import (
	"fmt"
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

func TestFilterMockHeaderContentSimpleHeader(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header3": []string{"Value5", "Value6"},
				"Header4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Header1": []string{"Value1", "Value2"},
	}
	filtered, err := FilterMockHeaderContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentDoubleHeader(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header3": []string{"Value5", "Value6"},
				"Header4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Header1": []string{"Value1", "Value2"},
		"Header2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockHeaderContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 2 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentNoMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header3": []string{"Value5", "Value6"},
				"Header4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Header10": []string{"Value1", "Value2"},
		"Header20": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockHeaderContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 0 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentNoHeaderWithAMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header3": []string{"Value5", "Value6"},
				"Header4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Header1": []string{"Value1", "Value2"},
		"Header2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockHeaderContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%v\n", filtered)
	fmt.Printf("%v", len(filtered))
	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentNoHeaderWithAMatch2(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header3": []string{"Value5", "Value6"},
				"Header4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Headers: map[string][]string{
				"Header1": []string{"Value1", "Value2"},
				"Header2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Header5": []string{"Value1", "Value2"},
		"Header6": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockHeaderContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%v\n", filtered)
	fmt.Printf("%v", len(filtered))
	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentSimplePayload(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload3": []string{"Value5", "Value6"},
				"Payload4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Payload1": []string{"Value1", "Value2"},
	}
	filtered, err := FilterMockPayloadContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockHeaderContentDoublePayload(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload3": []string{"Value5", "Value6"},
				"Payload4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Payload1": []string{"Value1", "Value2"},
		"Payload2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 2 {
		t.Fail()
	}
}

func TestFilterMockPayloadContentNoMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload3": []string{"Value5", "Value6"},
				"Payload4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Payload10": []string{"Value1", "Value2"},
		"Payload20": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 0 {
		t.Fail()
	}
}

func TestFilterMockPayloadContentNoPayloadWithAMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload3": []string{"Value5", "Value6"},
				"Payload4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Payload: map[string][]string{
				"Payload1": []string{"Value1", "Value2"},
				"Payload2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	headers := map[string][]string{
		"Payload1": []string{"Value1", "Value2"},
		"Payload2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, headers)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%v\n", filtered)
	fmt.Printf("%v", len(filtered))
	if len(filtered) != 1 {
		t.Fail()
	}
}
