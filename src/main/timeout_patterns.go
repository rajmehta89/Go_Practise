package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// 1. Using time.After with select
func timeoutWithSelect(duration time.Duration) {
    ch := make(chan string)

    // Long running operation
    go func() {
        time.Sleep(2 * time.Second) // Simulate work
        ch <- "Operation completed"
    }()

    // Wait with timeout
    select {
    case result := <-ch:
        fmt.Println("Success:", result)
    case <-time.After(duration):
        fmt.Println("Operation timed out")
        // Cleanup code here
    }
}

// 2. Using context with timeout
func timeoutWithContext(duration time.Duration) {
    ctx, cancel := context.WithTimeout(context.Background(), duration)
    defer cancel() // Always cancel to release resources

    doneCh := make(chan string)

    go func() {
        // Simulate work that checks context
        for i := 0; i < 5; i++ {
            select {
            case <-ctx.Done():
                return
            default:
                time.Sleep(500 * time.Millisecond)
            }
        }
        doneCh <- "Work completed"
    }()

    select {
    case result := <-doneCh:
        fmt.Println("Context success:", result)
    case <-ctx.Done():
        fmt.Println("Context timeout:", ctx.Err())
    }
}

// 3. Using timer reset pattern
func timeoutWithTimer() {
    timer := time.NewTimer(1 * time.Second)
    defer timer.Stop()

    ch := make(chan string)
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Late result"
    }()

    select {
    case result := <-ch:
        if !timer.Stop() {
            <-timer.C // Drain the channel
        }
        fmt.Println("Timer got result:", result)
    case <-timer.C:
        fmt.Println("Timer expired")
    }
}

// 4. Timeout with recovery
func timeoutWithRecovery(duration time.Duration) {
    done := make(chan bool)
    result := make(chan string)

    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from: %v\n", r)
                done <- true
            }
        }()

        // Simulate work that might panic
        time.Sleep(2 * time.Second)
        if duration < 3*time.Second {
            panic("Operation took too long!")
        }
        result <- "Success"
    }()

    select {
    case <-done:
        fmt.Println("Operation failed but recovered")
    case res := <-result:
        fmt.Println("Operation succeeded:", res)
    case <-time.After(duration):
        fmt.Println("Operation timed out")
    }
}

// 5. Timeout with cleanup
func timeoutWithCleanup(duration time.Duration) {
    var wg sync.WaitGroup
    cleanup := make(chan struct{})
    result := make(chan string)

    // Start cleanup goroutine
    wg.Add(1)
    go func() {
        defer wg.Done()
        select {
        case <-cleanup:
            fmt.Println("Cleaning up resources...")
            time.Sleep(100 * time.Millisecond)
            fmt.Println("Cleanup completed")
        }
    }()

    // Main operation
    go func() {
        time.Sleep(2 * time.Second)
        result <- "Operation result"
    }()

    // Wait with timeout
    select {
    case res := <-result:
        fmt.Println("Got result:", res)
    case <-time.After(duration):
        fmt.Println("Timeout occurred, initiating cleanup")
        close(cleanup)
    }

    wg.Wait() // Wait for cleanup to complete
}

func main() {
    fmt.Println("=== Basic Timeout ===")
    timeoutWithSelect(1 * time.Second)

    fmt.Println("\n=== Context Timeout ===")
    timeoutWithContext(1 * time.Second)

    fmt.Println("\n=== Timer Timeout ===")
    timeoutWithTimer()

    fmt.Println("\n=== Timeout with Recovery ===")
    timeoutWithRecovery(1 * time.Second)

    fmt.Println("\n=== Timeout with Cleanup ===")
    timeoutWithCleanup(1 * time.Second)
} 