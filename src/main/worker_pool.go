package main

import (
	"fmt"
	"sync"
	"time"
)

// Job types and their processors
const (
	ProcessData = iota  // 0
	FileOperation      // 1
	Calculation       // 2
)

// Process different types of jobs
func processJob(jobID int) int {
	// Determine job type by modulo
	switch jobID % 3 {
	case ProcessData:
		// Simulate data processing
		time.Sleep(200 * time.Millisecond)
		return jobID * 100

	case FileOperation:
		// Simulate file operation
		time.Sleep(300 * time.Millisecond)
		return jobID + 50

	case Calculation:
		// Simulate calculation
		time.Sleep(150 * time.Millisecond)
		return jobID * jobID
	}
	return 0
}

// Worker function to process jobs
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()  // Mark this worker as done when finished

	// Process jobs until channel is closed
	for job := range jobs {
		// Log start of job
		fmt.Printf("Worker %d starting job %d\n", id, job)
		
		// Process the job
		result := processJob(job)
		
		// Safely send result
		select {
		case results <- result:
			fmt.Printf("Worker %d completed job %d with result %d\n", id, job, result)
		default:
			fmt.Printf("Worker %d couldn't send result for job %d\n", id, job)
			return
		}
	}
}

func main() {
	fmt.Println("\n=== Worker Pool Example ===")
	
	const numJobs = 6
	const numWorkers = 3
	
	// Create buffered channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// WaitGroup to track workers
	var wg sync.WaitGroup
	
	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)  // No more jobs to send
	
	// Wait for workers in separate goroutine
	go func() {
		wg.Wait()
		close(results)  // Close results when all workers done
		fmt.Println("All workers finished")
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Got result: %d\n", result)
	}
} 