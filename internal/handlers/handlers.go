package handlers

import (
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

func manageError(err error, w http.ResponseWriter) {
	log.Println(err)
	errorMessage := fmt.Sprintf("{\n\t\"error\": \"%s\"\n}", err)
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

		if err != nil {
			manageError(err, w)
		} else {
			manageSuccess(w, r, content)
		}
	}
}
