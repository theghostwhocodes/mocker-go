package filters

import (
	"net/http"

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

// FilterMockHeaderContent filters mock by HTTP headers content
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

			if checkArrayEquality(values, headerValues) {
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

	return emptyHeaderMatches, nil
}
