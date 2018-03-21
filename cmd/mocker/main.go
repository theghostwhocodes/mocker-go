package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/theghostwhocodes/mocker-go/internal/handlers"
)

var version string
var showVersion bool
var dataPath string
var port int
var host string

func init() {
	const (
		dataPathDefault  = "./data"
		dataPathHelpText = "The data path"
		portDefault      = 8000
		portHelpText     = "The TCP port to listen"
		hostDefault      = "127.0.0.1"
		hostHelpText     = "The host to listen"
	)

	flag.BoolVar(&showVersion, "version", false, "Mocker version")
	flag.StringVar(&dataPath, "d", dataPathDefault, dataPathHelpText+" (shorthand)")
	flag.StringVar(&dataPath, "data", dataPathDefault, dataPathHelpText)
	flag.IntVar(&port, "p", portDefault, portHelpText+" (shorthand)")
	flag.IntVar(&port, "port", portDefault, portHelpText)
	flag.StringVar(&host, "h", hostDefault, hostHelpText+" (shorthand)")
	flag.StringVar(&host, "host", hostDefault, hostHelpText)
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("Mocker version %s", version)
		return
	}

	basePath, _ := filepath.Abs(dataPath)

	http.HandleFunc("/", handlers.HandlerFactory(basePath))
	log.Printf("Loading data from %s", basePath)
	log.Printf("Mocker listening on %s:%d...", host, port)

	address := fmt.Sprintf("%s:%d", host, port)
	http.ListenAndServe(address, nil)
}
