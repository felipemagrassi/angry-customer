package internal

import (
	"fmt"
	"time"
)

type Report struct {
	totalTime           time.Duration
	totalRequests       uint64
	requestErrors       uint64
	requestDistribution map[int]uint64
}

func NewReport() *Report {
	return &Report{
		requestDistribution: make(map[int]uint64),
	}
}

func (r *Report) AddRequest(statusCode int) {
	r.totalRequests++
	codeCount, ok := r.requestDistribution[statusCode]
	if !ok {
		codeCount = 0
		r.requestDistribution[statusCode] = codeCount
	}
	r.requestDistribution[statusCode] = codeCount + 1
}

func (r *Report) AddRequestError() {
	r.totalRequests++
	r.requestErrors++
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
	fmt.Println("Successful requests: ", r.totalRequests-r.requestErrors)
	fmt.Println("Requests with errors: ", r.requestErrors)
	fmt.Println("200:", r.requestDistribution[200])
	for code, count := range r.requestDistribution {
		if code != 200 {
			fmt.Printf("%d: %d\n", code, count)
		}
	}
}
