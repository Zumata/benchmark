package main

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func rawDataHandler(w http.ResponseWriter, r *http.Request) {
	gzipping := requestingGzip(r)
	log.Printf("Raw data handler hit, requesting gzip: %v", gzipping == true)

	if !gzipping {
		// Do nothing
	}

	jsonBytes, err := json.Marshal(hotelPackages)
	if err != nil {
		log.Printf("unable to json marshal hotel packages due to %v", err)

		w.Write([]byte("Error Occurred"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
	w.WriteHeader(http.StatusOK)

	return
}

func gzippedDataHandler(w http.ResponseWriter, r *http.Request) {
	gzipping := requestingGzip(r)
	log.Printf("Gzipped data handler hit, requesting gzip: %v", gzipping == true)

	if !gzipping {
		w.Write([]byte("Bad Request"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	err := json.NewEncoder(gzipWriter).Encode(hotelPackages)
	if err != nil {
		log.Printf("unable to gzip hotel packages due to %v", err)

		w.Write([]byte("Error Occurred"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(http.StatusOK)

	return
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Healthcheck OK!")
	w.Write([]byte(`OK`))
}

func requestingGzip(r *http.Request) bool {
	acceptedEncodings := r.Header.Get("Accept-Encoding")

	// Simplistic check for gzip
	return strings.Contains(acceptedEncodings, "gzip")
}
