package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
)

func getContent(basePath string, r *http.Request) ([]byte, error) {
	method := r.Method
	fileName := fmt.Sprintf("%s.%s.json", r.URL.Path[1:], strings.ToUpper(method))
	content, err := ioutil.ReadFile(
		path.Join(
			basePath,
			fileName,
		),
	)
	return content, err
}

func isValidJSON(content []byte) bool {
	var result map[string]interface{}
	return json.Unmarshal(content, &result) == nil
}

func sendErrorMessage(message string, w http.ResponseWriter) {
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
		content, err := getContent(basePath, r)
		isValid := isValidJSON(content)

		if err != nil {
			sendErrorMessage(err.Error(), w)
			return
		}

		if !isValid {
			sendErrorMessage("The mock file contains invalid JSON", w)
			return
		}

		manageSuccess(w, r, content)
	}
}
