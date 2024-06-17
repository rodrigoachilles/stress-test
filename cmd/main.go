package main

import (
	"flag"
	"fmt"
	"github.com/rodrigoachilles/stress-test/configs/logger"
	"net/http"
	"os"
	"time"
)

type Result struct {
	StatusCode int
}

type ResultOutput struct {
	Elapsed                                                                   time.Duration
	TotalRequests                                                             int
	StatusCode1xx, StatusCode2xx, StatusCode3xx, StatusCode4xx, StatusCode5xx []int
}

func main() {
	var (
		url         string
		requests    int
		concurrency int
	)

	flag.StringVar(&url, "url", "", "URL of the web service to be tested")
	flag.IntVar(&requests, "requests", 0, "Total number of requests to be made")
	flag.IntVar(&concurrency, "concurrency", 1, "Number of simultaneous calls")

	flag.Parse()

	if url == "" || requests <= 0 || concurrency <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	startTime := time.Now()
	results := make(chan *Result)

	go stressTest(url, requests, concurrency, results)

	resultOutput := ResultOutput{}
	var totalRequests int

	for i := 0; i < requests; i++ {
		result := <-results
		totalRequests++
		addResultToReport(result, &resultOutput)
	}
	resultOutput.Elapsed = time.Since(startTime)
	resultOutput.TotalRequests = totalRequests

	report(resultOutput)
}

func addResultToReport(result *Result, resultOutput *ResultOutput) {
	firstDigit := result.StatusCode / 100

	switch firstDigit {
	case 1:
		resultOutput.StatusCode1xx = append(
			resultOutput.StatusCode1xx,
			result.StatusCode,
		)
	case 2:
		resultOutput.StatusCode2xx = append(
			resultOutput.StatusCode2xx,
			result.StatusCode,
		)
	case 3:
		resultOutput.StatusCode3xx = append(
			resultOutput.StatusCode3xx,
			result.StatusCode,
		)
	case 4:
		resultOutput.StatusCode4xx = append(
			resultOutput.StatusCode4xx,
			result.StatusCode,
		)
	case 5:
		resultOutput.StatusCode5xx = append(
			resultOutput.StatusCode5xx,
			result.StatusCode,
		)
	}
}

func stressTest(url string, requests, concurrency int, results chan *Result) {
	requestsPerWorker := requests / concurrency
	httpClient := &http.Client{}

	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < requestsPerWorker; j++ {
				resp, _ := httpClient.Get(url)
				results <- &Result{StatusCode: resp.StatusCode}
			}
		}()
	}

	remainder := requests % concurrency
	for j := 0; j < remainder; j++ {
		resp, _ := httpClient.Get(url)
		results <- &Result{StatusCode: resp.StatusCode}
	}
}

func report(resultOutput ResultOutput) {
	logger.Info("=== Stress Test Results ===")
	logger.Info(fmt.Sprintf("Total time: %s", resultOutput.Elapsed))
	logger.Info(fmt.Sprintf("Total requests: %d", resultOutput.TotalRequests))
	reportStatusCode("1xx", resultOutput.StatusCode1xx)
	reportStatusCode("2xx", resultOutput.StatusCode2xx)
	reportStatusCode("3xx", resultOutput.StatusCode3xx)
	reportStatusCode("4xx", resultOutput.StatusCode4xx)
	reportStatusCode("5xx", resultOutput.StatusCode5xx)
}

func reportStatusCode(statusCodeTitle string, statusCodes []int) {
	logger.Info(fmt.Sprintf("HTTP Status Code [%s]: %d", statusCodeTitle, len(statusCodes)))
	for status, count := range groupStatusCode(statusCodes) {
		logger.Info(fmt.Sprintf(" * [%d] : %d", status, count))
	}
}

func groupStatusCode(statusCodes []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range statusCodes {
		counts[value]++
	}
	return counts
}
