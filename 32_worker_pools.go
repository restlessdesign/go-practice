package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Processing job ", job, " on worker ", id)

		// Simulate long computation
		time.Sleep(time.Second)

		results <- job * 10
	}
}

func createWorkers(worker_count int, jobs <-chan int, results chan<- int) {
	for i := 1; i <= worker_count; i++ {
		go worker(i, jobs, results)
	}
}

func processJobs(jobs chan<- int, total_jobs int) {
	for j := 1; j <= total_jobs; j++ {
		jobs <- j
	}
}

func processResults(results <-chan int, total_jobs int) {
	for k := 1; k <= total_jobs; k++ {
		<-results
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	total_jobs := 9

	createWorkers(3, jobs, results)
	processJobs(jobs, total_jobs)
	close(jobs)
	processResults(results, total_jobs)
}
