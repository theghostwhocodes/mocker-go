package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
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
	fmt.Fprintln(w, err)
}

func manageSuccess(w http.ResponseWriter, r *http.Request, content []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, string(content))
	log.Printf("Serving %s\n", r.URL.Path[1:])
}

func handlerFactory(basePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := getContent(basePath, r)

		if err != nil {
			manageError(err, w)
		} else {
			manageSuccess(w, r, content)
		}
	}
}

func main() {
	dataPath := flag.String("path", "./data", "The data path")
	port := flag.Int("port", 8000, "The TCP port to listen")
	host := flag.String("host", "127.0.0.1", "The host to listen")
	flag.Parse()

	basePath, _ := filepath.Abs(*dataPath)

	http.HandleFunc("/", handlerFactory(basePath))

	log.Printf("Loading data from %s", basePath)
	log.Printf("Mocker listening on %s:%d...", *host, *port)

	address := fmt.Sprintf("%s:%d", *host, *port)
	http.ListenAndServe(address, nil)
}
