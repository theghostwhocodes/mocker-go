package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/theghostwhocodes/mocker-go/internal/contentManagers"
)

func isValidJSON(content []byte) bool {
	var result map[string]interface{}
	return json.Unmarshal(content, &result) == nil
}

func getMapFromBytes(content []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(content, &result)
	return result
}

func hasRequest(content map[string]interface{}) bool {
	_, ok := content["request"]
	return ok
}

func hasResponse(content map[string]interface{}) bool {
	_, ok := content["response"]
	return ok
}

func sendErrorMessage(w http.ResponseWriter, message string) {
	errorMessage := fmt.Sprintf("{\n\t\"error\": \"%s\"\n}", message)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, errorMessage)
}

func manageSuccess(w http.ResponseWriter, r *http.Request, content []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, string(content))
	log.Printf("Serving %s\n", r.URL.Path[1:])
}

// HandlerFactory return a proper handler
func HandlerFactory(basePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := contentManagers.GetContent(basePath, r)
		isValid := isValidJSON(content)

		if err != nil {
			sendErrorMessage(w, err.Error())
			return
		}

		if !isValid {
			sendErrorMessage(w, "The mock file contains invalid JSON")
			return
		}

		jsonMap := getMapFromBytes(content)
		if !hasRequest(jsonMap) {
			sendErrorMessage(w, "Oops, probably your mock file doesn't contain 'request' section")
			return
		}
		if !hasResponse(jsonMap) {
			sendErrorMessage(w, "Oops, probably you mock file doesn't contain 'response' section")
			return
		}

		manageSuccess(w, r, content)
	}
}
