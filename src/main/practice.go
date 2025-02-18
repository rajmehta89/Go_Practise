package main

import (
    "fmt"
    "time"
)

// Example 1: Basic Defer and Panic
func deferPanicExample() {
    // Multiple defers - executed in LIFO order
    defer fmt.Println("1. First defer")
    defer fmt.Println("2. Second defer")
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("3. Recovered from: %v\n", r)
        }
    }()

    fmt.Println("4. Before panic")
    panic("Something went wrong!")
    fmt.Println("This won't be printed")
}

// Example 2: Channel with Error Handling
func safeSender(ch chan<- int, val int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Sender recovered from: %v\n", r)
        }
    }()
    
    ch <- val
}

// Example 3: Resource Cleanup with Defer
type Resource struct {
    name string
}

func (r *Resource) Close() {
    fmt.Printf("Closing resource: %s\n", r.name)
}

func resourceExample() {
    r1 := &Resource{"Database"}
    r2 := &Resource{"File"}
    
    defer r1.Close()
    defer r2.Close()
    
    fmt.Println("Working with resources...")
    panic("Resource error!")
}

// Example 4: Channel Timeout Pattern
func timeoutExample(ch chan string) {
    defer close(ch)
    
    for i := 1; i <= 5; i++ {
        select {
        case ch <- fmt.Sprintf("Message %d", i):
            time.Sleep(100 * time.Millisecond)
        case <-time.After(300 * time.Millisecond):
            panic("Send timeout!")
        }
    }
}

func main() {
    fmt.Println("=== Defer and Panic Example ===")
    deferPanicExample()
    
    fmt.Println("\n=== Channel Example ===")
    ch := make(chan int, 1)
    safeSender(ch, 42)
    close(ch)
    safeSender(ch, 43) // This would panic but is recovered
    
    fmt.Println("\n=== Resource Example ===")
    func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Resource error recovered: %v\n", r)
            }
        }()
        resourceExample()
    }()
    
    fmt.Println("\n=== Timeout Example ===")
    msgCh := make(chan string)
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Timeout recovered: %v\n", r)
            }
        }()
        timeoutExample(msgCh)
    }()
    
    // Receive messages with timeout
    for i := 0; i < 6; i++ {
        select {
        case msg, ok := <-msgCh:
            if !ok {
                fmt.Println("Channel closed")
                break
            }
            fmt.Printf("Received: %s\n", msg)
        case <-time.After(500 * time.Millisecond):
            fmt.Println("Receive timeout")
            return
        }
    }
} 