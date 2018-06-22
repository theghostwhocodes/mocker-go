package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/theghostwhocodes/mocker-go/internal/handlers"
	"golang.org/x/sys/unix"
)

var version string
var showVersion bool
var dataPath string
var port int
var host string
var proxyFor string

func init() {
	const (
		dataPathDefault  = "./data"
		dataPathHelpText = "The data path"
		portDefault      = 8000
		portHelpText     = "The TCP port to listen"
		hostDefault      = "127.0.0.1"
		hostHelpText     = "The host to listen"
		proxyForDefault  = ""
		proxyForHelpText = "The real API endpoint"
	)

	flag.BoolVar(&showVersion, "version", false, "Mocker version")
	flag.StringVar(&dataPath, "d", dataPathDefault, dataPathHelpText+" (shorthand)")
	flag.StringVar(&dataPath, "data", dataPathDefault, dataPathHelpText)
	flag.IntVar(&port, "p", portDefault, portHelpText+" (shorthand)")
	flag.IntVar(&port, "port", portDefault, portHelpText)
	flag.StringVar(&host, "h", hostDefault, hostHelpText+" (shorthand)")
	flag.StringVar(&host, "host", hostDefault, hostHelpText)
	flag.StringVar(&proxyFor, "pf", proxyForDefault, proxyForHelpText+" (shorthand)")
	flag.StringVar(&proxyFor, "proxy-for", proxyForDefault, proxyForHelpText)
}

func cleanup() {
	fmt.Printf("Exiting mocker...\n")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("Mocker version %s", version)
		return
	}

	basePath, _ := filepath.Abs(dataPath)

	err := unix.Access(basePath, unix.R_OK+unix.W_OK)
	if err != nil {
		log.Printf("I can't access your data folder, please check folder permissions")
		os.Exit(1)
	}

	http.HandleFunc("/", handlers.HandlerFactory(basePath, proxyFor))
	log.Printf("Loading data from %s", basePath)
	log.Printf("Mocker listening on %s:%d...", host, port)

	if proxyFor != "" {
		log.Printf("Mocker acting as a proxy for %s...", proxyFor)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	address := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(address, nil))
}
