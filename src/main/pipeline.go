package main

import (
    "fmt"
    "time"
)

// Stage 1: Generate numbers
func generateNumbers() <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        // Generate numbers 1 to 5
        for i := 1; i <= 5; i++ {
            fmt.Printf("Generating: %d\n", i)
            out <- i
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

// Stage 2: Double the numbers
func doubleNumbers(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            result := num * 2
            fmt.Printf("Doubling: %d -> %d\n", num, result)
            out <- result
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

// Stage 3: Filter even numbers
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            if num%2 == 0 {
                fmt.Printf("Filtering: keeping %d\n", num)
                out <- num
            } else {
                fmt.Printf("Filtering: dropping %d\n", num)
            }
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return out
}

func main() {
    fmt.Println("Starting Pipeline Example")
    
    // Create the pipeline
    numbers := generateNumbers()           // Stage 1
    doubled := doubleNumbers(numbers)      // Stage 2
    filtered := filterEven(doubled)        // Stage 3
    
    // Collect results
    for result := range filtered {
        fmt.Printf("Got result: %d\n", result)
    }
    
    fmt.Println("Pipeline Complete")
} 