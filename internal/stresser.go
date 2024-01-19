package internal

import (
	"errors"
	"net/http"
)

func RunStresser(url string, numberOfRequests, numberOfThreads uint64) error {
	if url == "" {
		return errors.New("url is required")
	}

	if numberOfRequests < 1 {
		return errors.New("number of requests must be greater than 0")
	}

	if numberOfThreads < 1 {
		return errors.New("number of threads must be greater than 0")
	}

	stresser(url, numberOfRequests, numberOfThreads)
	return nil
}

func stresser(url string, numberOfRequests, numberOfThreads uint64) {
}

func makeRequest(url string) {
	http.Get(url)
}
