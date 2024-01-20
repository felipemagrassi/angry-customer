package internal

import (
	"fmt"
	"sync"
	"time"
)

type Report struct {
	URL                 string
	totalTime           time.Duration
	totalRequests       uint64
	imcompleteRequests  uint64
	requestDistribution map[int]uint64
	mutex               sync.Mutex
}

func NewReport() *Report {
	return &Report{
		requestDistribution: make(map[int]uint64),
	}
}

func (r *Report) AddRequest(statusCode int) {
	r.mutex.Lock()

	r.totalRequests++
	codeCount, ok := r.requestDistribution[statusCode]
	if !ok {
		codeCount = 0
		r.requestDistribution[statusCode] = codeCount
	}
	r.requestDistribution[statusCode] = codeCount + 1

	r.mutex.Unlock()
}

func (r *Report) AddRequestError() {
	r.mutex.Lock()

	r.totalRequests++
	r.imcompleteRequests++

	r.mutex.Unlock()
}

func (r *Report) SetTotalTime(totalTime time.Duration) {
	r.totalTime = totalTime
}

func (r *Report) Print() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Report:")
	fmt.Println("Total time: ", r.totalTime)
	fmt.Println("Total requests: ", r.totalRequests)
	fmt.Println("Successful requests: ", r.totalRequests-r.imcompleteRequests)
	fmt.Println("Imcomplete requests: ", r.imcompleteRequests)
	fmt.Println("Requests with status 200:", r.requestDistribution[200])
	fmt.Println("Status code distribution:")
	for code, count := range r.requestDistribution {
		formated_percent := fmt.Sprintf("%d: %f%%\n", code, float64(count)/float64(r.totalRequests)*100)
		fmt.Print(formated_percent)
	}
	fmt.Print("Imcomplete Requests: ", float64(r.imcompleteRequests)/float64(r.totalRequests)*100, "%\n")
}
