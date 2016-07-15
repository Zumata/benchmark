package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	iterations int = 50
)

var (
	serverURLRoot string
)

type appHandlerLight struct {
	h func(http.ResponseWriter, *http.Request)
}

func (ah appHandlerLight) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.h(w, r)
}

func initRouter() error {
	router := mux.NewRouter()
	router.Handle("/data_compression_test", appHandlerLight{dataCompressionTestHandler}).Methods("GET")
	router.Handle("/healthcheck", appHandlerLight{healthCheckHandler}).Methods("GET")

	s := &http.Server{
		Addr:           ":3334",
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	return s.ListenAndServe()
}

func main() {
	serverHost, ok := os.LookupEnv("GZIP_TEST_SERVICE_HOST")
	if !ok {
		log.Fatalf("server host env was not set with GZIP_TEST_SERVICE_HOST")
	}
	serverPort, ok := os.LookupEnv("GZIP_TEST_SERVICE_PORT")
	if !ok {
		log.Fatalf("server port env was not set with GZIP_TEST_SERVICE_PORT")
	}

	serverURLRoot = "http://" + serverHost + ":" + serverPort + "/"

	err := initRouter()
	if err != nil {
		log.Fatalf("unable to start server due to %v", err)
	}
}
