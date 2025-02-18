package main

import (
    "fmt"
    "sync"
    "time"
)

func deferInGoroutineWithDelay() {
    var wg sync.WaitGroup
    
    // Run multiple times to show different orderings
    for run := 1; run <= 3; run++ {
        fmt.Printf("\nRun %d:\n", run)
        wg.Add(3)
        
        for i := 0; i < 3; i++ {
            go func(id int) {
                defer wg.Done()
                defer fmt.Printf("Goroutine %d defer\n", id)
                
                // Random delay to show different orderings
                time.Sleep(time.Duration(id) * 100 * time.Millisecond)
                fmt.Printf("Goroutine %d executing\n", id)
            }(i)
        }
        
        wg.Wait()
        fmt.Println("-------------------")
    }
}

func main() {
    fmt.Println("Starting goroutine examples...")
    deferInGoroutineWithDelay()
} 