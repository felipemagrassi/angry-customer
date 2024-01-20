package internal

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
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

	if numberOfThreads > numberOfRequests {
		fmt.Println("Number of threads greater than number of requests, setting number of threads to number of requests")
		numberOfThreads = numberOfRequests
	}

	stresser(url, numberOfRequests, numberOfThreads)
	return nil
}

func stresser(url string, numberOfRequests, numberOfThreads uint64) {
	report := NewReport()
	fmt.Println(report.requestDistribution)
	fmt.Printf("Starting stresser with %d requests and %d threads\n", numberOfRequests, numberOfThreads)

	threadControl := make(chan struct{}, numberOfThreads)
	start := time.Now()

	for i := uint64(0); i < numberOfRequests; i++ {
		threadControl <- struct{}{}
		printProgressBar(i, numberOfRequests, len(threadControl))
		wg.Add(1)
		go makeRequest(url, threadControl, report)
	}

	wg.Wait()
	finish := time.Now()

	report.SetTotalTime(finish.Sub(start))
	report.Print()
}

func makeRequest(url string, threadControl <-chan struct{}, report *Report) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		report.AddRequestError()
		<-threadControl
		return
	}

	time.Sleep(200 * time.Millisecond)
	go report.AddRequest(res.StatusCode)

	<-threadControl
}

func printProgressBar(current, total uint64, threadLen int) {
	fmt.Printf("\r%d%% [%d/%d] Threads: %d", (current*100)/total, current, total, threadLen)
}
