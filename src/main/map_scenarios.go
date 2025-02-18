package main

import (
    "fmt"
    "sync"
    "time"
)

// Thread-safe map wrapper
type SafeMap struct {
    sync.RWMutex
    data map[string]int
}

// SafeMap methods
func (sm *SafeMap) Set(key string, value int) {
    sm.Lock()
    defer sm.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.RLock()
    defer sm.RUnlock()
    val, exists := sm.data[key]
    return val, exists
}

func (sm *SafeMap) Delete(key string) {
    sm.Lock()
    defer sm.Unlock()
    delete(sm.data, key)
}

// Worker that processes map operations
func mapWorker(id int, sm *SafeMap, ops chan string, results chan string, done chan bool) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Worker %d recovered from: %v\n", id, r)
        }
    }()

    for {
        select {
        case op, ok := <-ops:
            if !ok {
                return
            }
            // Process operation
            sm.Set(fmt.Sprintf("worker_%d_%s", id, op), id)
            results <- fmt.Sprintf("Worker %d processed %s", id, op)
            
        case <-done:
            fmt.Printf("Worker %d shutting down\n", id)
            return
            
        case <-time.After(2 * time.Second):
            fmt.Printf("Worker %d timed out\n", id)
            return
        }
    }
}

// Demonstrates different map scenarios
func demonstrateMapScenarios() {
    // Initialize safe map
    sm := &SafeMap{
        data: make(map[string]int),
    }

    // Channels for coordination
    ops := make(chan string, 10)
    results := make(chan string, 10)
    done := make(chan bool)

    // Start multiple workers
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            mapWorker(id, sm, ops, results, done)
        }(i)
    }

    // Send operations
    go func() {
        operations := []string{"add", "update", "delete"}
        for _, op := range operations {
            ops <- op
            time.Sleep(100 * time.Millisecond)
        }
        close(ops)
    }()

    // Collect results with timeout
    go func() {
        for {
            select {
            case result, ok := <-results:
                if !ok {
                    return
                }
                fmt.Println(result)
            
            case <-time.After(1* time.Millisecond):
                fmt.Println("Result collection timed out")
                close(done)
                return
            }
        }
    }()

    // Wait for all workers
    wg.Wait()
    close(results)

    // Print final map state
    fmt.Println("\nFinal map state:")
    sm.Lock()
    defer sm.Unlock()
    for k, v := range sm.data {
        fmt.Printf("%s: %d\n", k, v)
    }
}

// Demonstrates select with multiple channels
func demonstrateSelect() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    done := make(chan bool)

    // Sender 1
    go func() {
        for i := 0; i < 3; i++ {
            ch1 <- fmt.Sprintf("Message from ch1: %d", i)
            time.Sleep(100 * time.Millisecond)
        }
    }()

    // Sender 2
    go func() {
        for i := 0; i < 3; i++ {
            ch2 <- fmt.Sprintf("Message from ch2: %d", i)
            time.Sleep(150 * time.Millisecond)
        }
    }()

    // Timeout after 1 second
    go func() {
        time.Sleep(1 * time.Second)
        done <- true
    }()

    // Select with multiple cases
    for {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        
        case msg2 := <-ch2:
            fmt.Println(msg2)
        
        case <-done:
            fmt.Println("Operation timed out")
            return
        
        default:
            // Non-blocking case
            fmt.Println("No messages available")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    fmt.Println("=== Map with Goroutines Example ===")
    demonstrateMapScenarios()

    fmt.Println("\n=== Select Pattern Example ===")
    demonstrateSelect()
} 