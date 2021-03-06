package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func manageSuccess(w http.ResponseWriter, r *http.Request, content []byte, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	fmt.Fprintln(w, string(content))
	log.Printf("Serving %s\n", r.URL.Path[1:])
}

// HandlerFactory return a proper handler
func HandlerFactory(basePath string, proxyFor string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := new(http.Response)
		var body []byte
		var err error
		var httpStatusCode int

		if proxyFor != "" {
			res, err = contentmanagers.ProxyFor(r, proxyFor)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}

			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}
			w.Header().Set("Mocker-Proxy", "true")
			httpStatusCode = res.StatusCode
		}

		if res.StatusCode == 404 || res.StatusCode == 405 || err != nil || proxyFor == "" {
			jsonMaps, err := contentmanagers.GetScannedMockContent(basePath, r)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}

			if len(jsonMaps) < 1 {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			jsonMap := jsonMaps[0]

			body, _ = json.Marshal(jsonMap.Response.Body)
			httpStatusCode = jsonMap.Response.Status
		}

		w.Header().Set("Mocker-Stubbed", "true")
		manageSuccess(w, r, body, httpStatusCode)
	}
}
