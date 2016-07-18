package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type appHandlerLight struct {
	h func(http.ResponseWriter, *http.Request)
}

func (ah appHandlerLight) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.h(w, r)
}

func initRouter() error {
	router := mux.NewRouter()
	router.Handle("/raw_data", appHandlerLight{rawDataHandler}).Methods("GET")
	router.Handle("/gzipped_data", appHandlerLight{gzippedDataHandler}).Methods("GET")
	router.Handle("/healthcheck", appHandlerLight{healthCheckHandler}).Methods("GET")

	s := &http.Server{
		Addr:           ":3333",
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	return s.ListenAndServe()
}

func main() {
	err := initRouter()
	if err != nil {
		log.Fatalf("unable to start server due to %v", err)
	}
}
