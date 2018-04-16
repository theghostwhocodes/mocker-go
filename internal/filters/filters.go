package filters

import (
	"net/http"
	"net/url"
	"reflect"

	"github.com/theghostwhocodes/mocker-go/internal/models"
)

func checkArrayEquality(array1 []string, array2 []string) bool {
	result := false
	if len(array1) != len(array2) {
		return result
	}

	for index, value := range array1 {
		if value == array2[index] {
			result = true
		} else {
			result = false
		}
	}

	return result
}

// FilterMockHTTPMethod filters mock by HTTP method
func FilterMockHTTPMethod(mocks []models.MockHTTP, method string) (results []models.MockHTTP, err error) {
	for _, mock := range mocks {
		if method == mock.Request.Method {
			results = append(results, mock)
		}
	}

	return results, nil
}

// FilterMockHeaderContent filters mocks by HTTP headers content
func FilterMockHeaderContent(mocks []models.MockHTTP, headers http.Header) (results []models.MockHTTP, err error) {
	var emptyHeaderMatches []models.MockHTTP
	var matches []models.MockHTTP
	for _, mock := range mocks {
		counter := 0
		matchCounter := 0

		if len(mock.Request.Headers) == 0 {
			emptyHeaderMatches = append(emptyHeaderMatches, mock)
			continue
		}

		for key, values := range mock.Request.Headers {
			counter++
			headerValues, ok := headers[key]
			if !ok {
				continue
			}

			if reflect.DeepEqual(values, headerValues) {
				matchCounter++
			}
		}

		if matchCounter == counter {
			matches = append(matches, mock)
		}
	}

	if len(matches) > 0 {
		return matches, nil
	}

	if len(emptyHeaderMatches) > 0 && len(headers) > 0 {
		var empty []models.MockHTTP
		return empty, nil
	}

	return emptyHeaderMatches, nil
}

// FilterMockPayloadContent filters mocks by HTTP payload content
func FilterMockPayloadContent(mocks []models.MockHTTP, payload interface{}) (results []models.MockHTTP, err error) {
	var emptyMatches []models.MockHTTP
	var matches []models.MockHTTP
	for _, mock := range mocks {
		mockPayload := mock.Request.Payload

		if mockPayload == nil {
			emptyMatches = append(emptyMatches, mock)
			continue
		}

		if reflect.DeepEqual(mockPayload, payload) {
			matches = append(matches, mock)
		}
	}

	if len(matches) > 0 {
		return matches, nil
	}

	return emptyMatches, nil
}

// FilterMockParamsContent filters mocks by HTTP query parameters content
func FilterMockParamsContent(mocks []models.MockHTTP, params url.Values) (results []models.MockHTTP, err error) {
	var emptyMatches []models.MockHTTP
	var matches []models.MockHTTP
	for _, mock := range mocks {
		mockParams := mock.Request.Params

		// if mockParams == nil {
		// 	emptyMatches = append(emptyMatches, mock)
		// 	continue
		// }

		if reflect.DeepEqual(mockParams, params) || (len(mockParams) == 0 && len(params) == 0) {
			matches = append(matches, mock)
		}
	}

	if len(matches) > 0 {
		return matches, nil
	}

	return emptyMatches, nil
}
