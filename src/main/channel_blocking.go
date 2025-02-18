package main

import (
    "fmt"
    "time"
)

// 1. Blocking on Unbuffered Channel
func unbufferedChannelBlocking() {
    fmt.Println("\n=== Unbuffered Channel Blocking ===")
    ch := make(chan int) // Unbuffered channel

    // Sender blocks until receiver is ready
    go func() {
        fmt.Println("Sender: About to send")
        ch <- 42 // This blocks until receiver is ready
        fmt.Println("Sender: Data sent")
    }()

    // Simulate delay in receiver
    time.Sleep(2 * time.Second)
    fmt.Println("Receiver: About to receive")
    data := <-ch
    fmt.Printf("Receiver: Got %d\n", data)
}

// 2. Blocking on Buffered Channel
func bufferedChannelBlocking() {
    fmt.Println("\n=== Buffered Channel Blocking ===")
    ch := make(chan int, 2) // Buffer size 2

    // Fill buffer
    ch <- 1
    fmt.Println("Sent 1")
    ch <- 2
    fmt.Println("Sent 2")

    // This will block because buffer is full
    go func() {
        fmt.Println("Sender: Trying to send 3")
        ch <- 3 // This blocks until space is available
        fmt.Println("Sender: Sent 3")
    }()

    time.Sleep(1 * time.Second)
    fmt.Println("Receiver: Starting to receive")
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
}

// 3. Blocking with Multiple Receivers
func multipleReceiversBlocking() {
    fmt.Println("\n=== Multiple Receivers Blocking ===")
    ch := make(chan int)

    // Multiple receivers - all block until data is available
    for i := 0; i < 3; i++ {
        go func(id int) {
            fmt.Printf("Receiver %d: Waiting for data\n", id)
            value := <-ch
            fmt.Printf("Receiver %d: Got %d\n", id, value)
        }(i)
    }

    time.Sleep(1 * time.Second)
    fmt.Println("Sender: Sending single value")
    ch <- 42 // Only one receiver gets this
    time.Sleep(1 * time.Second)
}

// 4. Blocking with Select and Default
func selectWithDefault() {
    fmt.Println("\n=== Select with Default (Non-blocking) ===")
    ch := make(chan int)

    // Non-blocking send
    select {
    case ch <- 42:
        fmt.Println("Sent data")
    default:
        fmt.Println("Send would block, skipping")
    }

    // Non-blocking receive
    select {
    case data := <-ch:
        fmt.Printf("Received: %d\n", data)
    default:
        fmt.Println("Receive would block, skipping")
    }
}

// 5. Blocking with Timeout
func blockingWithTimeout() {
    fmt.Println("\n=== Blocking with Timeout ===")
    ch := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- 42
    }()

    select {
    case data := <-ch:
        fmt.Printf("Received: %d\n", data)
    case <-time.After(1 * time.Second):
        fmt.Println("Timed out waiting for data")
    }
}

// 6. Deadlock Example
func deadlockExample() {
    fmt.Println("\n=== Deadlock Example ===")
    ch := make(chan int)

    // This would deadlock if uncommented:
    // ch <- 1  // Blocks forever - deadlock!
    
    // Safe way with goroutine:
    go func() {
        ch <- 1
    }()
    fmt.Printf("Received: %d\n", <-ch)
}

func main() {
    unbufferedChannelBlocking()
    bufferedChannelBlocking()
    multipleReceiversBlocking()
    selectWithDefault()
    blockingWithTimeout()
    deadlockExample()
} 