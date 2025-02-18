package main

import (
    "fmt"
    "time"
)

// 1. Basic Deadlock - Blocking without goroutine
func basicDeadlock() {
    fmt.Println("\n=== Basic Deadlock ===")
    ch := make(chan int)
    
    // This causes deadlock - main thread blocks forever
    // ch <- 1  // Send blocks waiting for receiver
    // OR
    // <-ch     // Receive blocks waiting for sender
    
    // Correct way: Use goroutine
    go func() {
        ch <- 1
    }()
    fmt.Printf("Received: %d\n", <-ch)
}

// 2. Circular Deadlock
func circularDeadlock() {
    fmt.Println("\n=== Circular Deadlock ===")
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    // This would deadlock if uncommented:
    /*
    go func() {
        val := <-ch1    // Waits for ch1
        ch2 <- val      // Then sends to ch2
    }()
    
    val := <-ch2        // Waits for ch2
    ch1 <- val         // Then sends to ch1
    */
    
    // Correct way: Break the circular dependency
    go func() {
        ch1 <- 1        // Send first
        fmt.Println(<-ch2)  // Then receive
    }()
    
    val := <-ch1        // Receive first
    ch2 <- val * 2      // Then send
}

// 3. Buffer Overflow Deadlock
func bufferDeadlock() {
    fmt.Println("\n=== Buffer Deadlock ===")
    ch := make(chan int, 2)
    
    // Fill buffer
    ch <- 1
    ch <- 2
    
    // This would deadlock if not in goroutine
    go func() {
        // Buffer is full, this will block
        ch <- 3
        fmt.Println("Sent 3")
    }()
    
    time.Sleep(time.Second)
    // Must receive to make space
    fmt.Printf("Received: %d\n", <-ch)
    time.Sleep(time.Second)
}

// 4. Multiple Channel Deadlock
func multiChannelDeadlock() {
    fmt.Println("\n=== Multiple Channel Deadlock ===")
    ch1 := make(chan int)
    ch2 := make(chan int)
    done := make(chan bool)
    
    // Potential deadlock if channels aren't managed properly
    go func() {
        select {
        case val := <-ch1:
            ch2 <- val * 2
        case <-done:
            return
        }
    }()
    
    // Safe way: Use timeout or done channel
    select {
    case ch1 <- 42:
        val := <-ch2
        fmt.Printf("Got result: %d\n", val)
    case <-time.After(time.Second):
        done <- true
        fmt.Println("Timed out")
    }
}

// 5. Goroutine Leak (not exactly deadlock but related)
func goroutineLeak() {
    fmt.Println("\n=== Goroutine Leak ===")
    ch := make(chan int)
    
    // This goroutine will be stuck forever
    go func() {
        val := <-ch    // Will never receive if no sender
        fmt.Println(val)
    }()
    
    // Correct way: Always provide a way to exit
    done := make(chan bool)
    go func() {
        select {
        case val := <-ch:
            fmt.Println(val)
        case <-done:
            fmt.Println("Goroutine exiting")
            return
        }
    }()
    
    time.Sleep(time.Second)
    done <- true
}

func main() {
    basicDeadlock()
    circularDeadlock()
    bufferDeadlock()
    multiChannelDeadlock()
    goroutineLeak()
} 