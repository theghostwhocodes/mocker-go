package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/theghostwhocodes/mocker-go/internal/contentmanagers"
)

func getMapFromBytes(content []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(content, &result)
	return result
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
		jsonMaps, err := contentmanagers.GetScannedMockContent(basePath, r)
		jsonMap := jsonMaps[0]

		body, _ := json.Marshal(jsonMap.Response.Body)
		if err != nil {
			sendErrorMessage(w, err.Error())
		}
		manageSuccess(w, r, body)
	}
}
