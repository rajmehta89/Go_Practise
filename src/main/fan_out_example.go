package main

import (
    "fmt"
    "sync"
    "time"
)

// Simulating work with different processing times
func processTask(id int, task int) int {
    // Simulate different processing times
    time.Sleep(time.Duration(task*100) * time.Millisecond)
    return task * 2
}

// Worker function
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task)
        result := processTask(id, task)
        results <- result
        fmt.Printf("Worker %d completed task %d\n", id, task)
    }
}

func main() {
    numWorkers := 3
    numTasks := 10

    // Create channels
    tasks := make(chan int, numTasks)
    results := make(chan int, numTasks)

    // Start multiple workers (Fan-out)
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, results, &wg)
    }

    // Send tasks
    go func() {
        for i := 1; i <= numTasks; i++ {
            tasks <- i
        }
        close(tasks)
    }()

    // Wait for all workers to finish and close results
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect results (Fan-in)
    for result := range results {
        fmt.Printf("Got result: %d\n", result)
    }
} 