package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type testResults struct {
	Iterations    int     `json:"iterations"`
	JSONInSeconds float64 `json:"json_in_seconds"`
	GZIPInSeconds float64 `json:"gzip_in_seconds"`
}

func dataCompressionTestHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	rawIterations := queryValues.Get("iterations")
	log.Printf("%v iterations requested", rawIterations)

	parsedIterations, _ := strconv.Atoi(rawIterations)
	if parsedIterations <= 0 {
		parsedIterations = iterations
	}

	// Without GZIP
	startTime := time.Now()

	err := performRequests(serverURLRoot+"raw_data", parsedIterations)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	jsonInSeconds := time.Now().Sub(startTime).Seconds()

	log.Printf("%v iterations WITHOUT gzip completed in %v seconds", parsedIterations, jsonInSeconds)

	///////////////
	// With GZIP //
	///////////////

	startTime = time.Now()

	err = performRequests(serverURLRoot+"gzipped_data", parsedIterations)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	gzipInSeconds := time.Now().Sub(startTime).Seconds()

	log.Printf("%v iterations WITH gzip completed in %v seconds", parsedIterations, gzipInSeconds)

	results := testResults{
		Iterations:    parsedIterations,
		JSONInSeconds: jsonInSeconds,
		GZIPInSeconds: gzipInSeconds,
	}

	resultsBytes, err := json.Marshal(results)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(resultsBytes)
	w.WriteHeader(http.StatusOK)

	return
}

func gzippedDataTestHandler(w http.ResponseWriter, r *http.Request) {
	// With GZIP
	startTime := time.Now()
	performRequests(serverURLRoot+"gzipped_data", iterations)
	timeInSeconds := time.Now().Sub(startTime).Seconds()

	log.Printf("%v iterations WITH gzip completed in %v seconds", iterations, timeInSeconds)
}

func performRequests(serverURL string, requestedIterations int) error {
	client := &http.Client{}

	request, err := http.NewRequest("GET", serverURL, nil)
	if err != nil {
		return fmt.Errorf("unable to construct http request due to %v", err)
	}
	// Force no caching
	request.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	for i := 0; i < requestedIterations; i++ {
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("unable to complete http get due to %v", err)
		}
		if response.StatusCode != http.StatusOK {
			return errors.New("unable to get a 200 OK response from the server")
		}
	}

	return nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Healthcheck OK!")
	w.Write([]byte(`OK`))
}
