package filters

import (
	"net/url"
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

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockPayloadContentSimple(t *testing.T) {
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
	payload := url.Values{
		"Payload1": []string{"Value1", "Value2"},
	}
	filtered, err := FilterMockPayloadContent(mocks, payload)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockPayloadContentDoublePayload(t *testing.T) {
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
	payload := url.Values{
		"Payload1": []string{"Value1", "Value2"},
		"Payload2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, payload)

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
	payload := map[string][]string{
		"Payload10": []string{"Value1", "Value2"},
		"Payload20": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, payload)

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
	payload := map[string][]string{
		"Payload1": []string{"Value1", "Value2"},
		"Payload2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, payload)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockPayloadContentNoPayloadWithAMatch2(t *testing.T) {
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
	payload := map[string][]string{
		"Payload5": []string{"Value1", "Value2"},
		"Payload6": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockPayloadContent(mocks, payload)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 0 {
		t.Fail()
	}
}

func TestFilterMockParametersContentSimple(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	params := url.Values{
		"Param1": []string{"Value1", "Value2"},
	}
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockParametersContentSimpleNoParameter(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	var params url.Values
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockParametersContentDoubleParameter(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	params := url.Values{
		"Param1": []string{"Value1", "Value2"},
		"Param2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 2 {
		t.Fail()
	}
}

func TestFilterMockParameterContentNoMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			// Params: url.Values{
			// 	"Param1": []string{"Value1", "Value2"},
			// 	"Param2": []string{"Value3", "Value4"},
			// },
		},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	params := url.Values{
		"Param1": []string{"Value10", "Value2"},
		"Param2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 0 {
		t.Fail()
	}
}

func TestFilterMockParameterContentNoParameterWithAMatch(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	params := url.Values{
		"Param1": []string{"Value1", "Value2"},
		"Param2": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestFilterMockParameterContentNoParameterWithAMatch2(t *testing.T) {
	mockHTTP1 := models.MockHTTP{
		Request: models.MockHTTPRequest{},
	}
	mockHTTP2 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param3": []string{"Value5", "Value6"},
				"Param4": []string{"Value7", "Value8"},
			},
		},
	}
	mockHTTP3 := models.MockHTTP{
		Request: models.MockHTTPRequest{
			Params: url.Values{
				"Param1": []string{"Value1", "Value2"},
				"Param2": []string{"Value3", "Value4"},
			},
		},
	}

	mocks := []models.MockHTTP{mockHTTP1, mockHTTP2, mockHTTP3}
	params := url.Values{
		"Param5": []string{"Value1", "Value2"},
		"Param6": []string{"Value3", "Value4"},
	}
	filtered, err := FilterMockParamsContent(mocks, params)

	if err != nil {
		t.Fail()
	}

	if len(filtered) != 0 {
		t.Fail()
	}
}
