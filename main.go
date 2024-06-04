package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	url           string
	totalRequests int
	concurrency   int
)

var rootCmd = &cobra.Command{
	Use:   "stress-tester",
	Short: "A simple HTTP stress tester",
	Run: func(cmd *cobra.Command, args []string) {
		runStressTest(url, totalRequests, concurrency)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8080", "URL to test (required)")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().IntVarP(&totalRequests, "requests", "r", 1000, "Total number of requests (required)")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent requests (required)")
	rootCmd.MarkFlagRequired("concurrency")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runStressTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	requestsPerGoroutine := totalRequests / concurrency
	extraRequests := totalRequests % concurrency

	results := make(chan int, totalRequests)
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int, requests int) {
			defer wg.Done()
			for j := 0; j < requests; j++ {
				resp, err := http.Get(url)
				if err != nil {
					log.Printf("Request failed: %v", err)
					results <- 0
					continue
				}
				results <- resp.StatusCode
				ioutil.ReadAll(resp.Body)
				resp.Body.Close()
			}
		}(i, requestsPerGoroutine+boolToInt(i < extraRequests))
	}

	wg.Wait()
	close(results)

	duration := time.Since(start)
	reportResults(duration, results)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func reportResults(duration time.Duration, results chan int) {
	total := 0
	status200 := 0
	statusCounts := make(map[int]int)

	for status := range results {
		total++
		if status == 200 {
			status200++
		}
		statusCounts[status]++
	}

	fmt.Printf("Total time taken: %v\n", duration)
	fmt.Printf("Total requests: %d\n", total)
	fmt.Printf("200 OK: %d\n", status200)
	for status, count := range statusCounts {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
