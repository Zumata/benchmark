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

type combinedResults struct {
	Iterations  int         `json:"iterations"`
	JSONResults *runResults `json:"json_results"`
	GZIPResults *runResults `json:"gzip_results"`
}

type runResults struct {
	AverageInSeconds float64 `json:"average_in_seconds"`
	FastestInSeconds float64 `json:"fastest_in_seconds"`
	SlowestInSeconds float64 `json:"slowest_in_seconds"`
}

func dataCompressionTestHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	rawIterations := queryValues.Get("iterations")
	log.Printf("%v iterations requested", rawIterations)

	parsedIterations, _ := strconv.Atoi(rawIterations)
	if parsedIterations <= 0 {
		parsedIterations = iterations
	}

	//////////////////
	// Without GZIP //
	//////////////////

	jsonRunResults, err := performRequests(serverURLRoot+"raw_data", parsedIterations)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	///////////////
	// With GZIP //
	///////////////

	gzipRunResults, err := performRequests(serverURLRoot+"gzipped_data", parsedIterations)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	results := combinedResults{
		Iterations:  parsedIterations,
		JSONResults: jsonRunResults,
		GZIPResults: gzipRunResults,
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

func performRequests(serverURL string, requestedIterations int) (*runResults, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", serverURL, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to construct http request due to %v", err)
	}
	// Force no caching
	request.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	currentRunResults := &runResults{
		FastestInSeconds: 1000000000,  // Just some large number
		SlowestInSeconds: -1000000000, // Just some small number
	}

	totalRunTime := float64(0)
	for i := 0; i < requestedIterations; i++ {
		startTime := time.Now()

		response, err := client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("unable to complete http get due to %v", err)
		}
		if response.StatusCode != http.StatusOK {
			return nil, errors.New("unable to get a 200 OK response from the server")
		}

		timeInSeconds := time.Now().Sub(startTime).Seconds()

		// Populate timings if needed
		if timeInSeconds < currentRunResults.FastestInSeconds {
			currentRunResults.FastestInSeconds = timeInSeconds
		}
		if timeInSeconds > currentRunResults.SlowestInSeconds {
			currentRunResults.SlowestInSeconds = timeInSeconds
		}

		totalRunTime += timeInSeconds
	}

	currentRunResults.AverageInSeconds = totalRunTime / float64(requestedIterations)

	return currentRunResults, nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Healthcheck OK!")
	w.Write([]byte(`OK`))
}
