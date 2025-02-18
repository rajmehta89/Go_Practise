package main

import (
    "fmt"
    "sync"
)

// Simple worker that processes numbers
func worker(id int, numbers <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for num := range numbers {
        // Process the number (multiply by 2)
        result := num * 2
        fmt.Printf("Worker %d processed %d -> %d\n", id, num, result)
        results <- result
    }
}

func main() {
    numbers := make(chan int, 10)
    results := make(chan int, 10)
    
    // Start 3 workers
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, numbers, results, &wg)
    }
    
    // Send numbers to process
    go func() {
        for i := 1; i <= 10; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    // Wait for workers and close results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    var processedNumbers []int
    for result := range results {
        processedNumbers = append(processedNumbers, result)
    }
    
    fmt.Printf("Processed numbers: %v\n", processedNumbers)
} 