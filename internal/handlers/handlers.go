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

func manageSuccess(w http.ResponseWriter, r *http.Request, content []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, string(content))
	log.Printf("Serving %s\n", r.URL.Path[1:])
}

// HandlerFactory return a proper handler
func HandlerFactory(basePath string, proxyFor string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := new(http.Response)
		var body []byte
		var err error

		if proxyFor != "" {
			fmt.Printf("ProxyFor\n")
			urlPath := r.URL.Path[1:]
			url := fmt.Sprintf("%s/%s", proxyFor, urlPath)
			httpVerb := r.Method
			httpHeaders := r.Header
			httpBody := r.Body
			fmt.Printf("Calling %s using verb %s...\n", url, httpVerb)
			// fmt.Printf("Headers %v\n", httpHeaders)
			// fmt.Printf("Body %v", httpBody)
			client := &http.Client{}
			req, err := http.NewRequest(httpVerb, url, httpBody)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}
			for key, value := range httpHeaders {
				req.Header.Add(key, value[0])
			}
			res, err = client.Do(req)
			fmt.Printf("A Response status code %d\n", res.StatusCode)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}

			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				sendErrorMessage(w, err.Error())
				return
			}
		}

		fmt.Printf("Response status code %d\n", res.StatusCode)
		if res.StatusCode == 404 || res.StatusCode == 405 || err != nil || proxyFor == "" {
			fmt.Printf("Normal\n")
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
		}

		manageSuccess(w, r, body)
	}
}
