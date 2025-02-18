package main

import (
    "fmt"
    "time"
)

func sender(ch chan int, name string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Sender %s trying to send: %d\n", name, i)
        ch <- i  // Might block if buffer is full
        fmt.Printf("Sender %s sent: %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func receiver(ch chan int, name string) {
    for {
        fmt.Printf("Receiver %s waiting...\n", name)
        value := <-ch  // Might block if buffer is empty
        fmt.Printf("Receiver %s got: %d\n", name, value)
        time.Sleep(300 * time.Millisecond)
    }
}

func BufferedDemoExample() {
    fmt.Println("\n=== Buffered Channel Blocking Demo ===")
    
    // Case 1: Buffer larger than sends (no blocking)
    fmt.Println("\nCase 1: Large Buffer (size 5)")
    ch1 := make(chan int, 5)
    go sender(ch1, "A")
    time.Sleep(1 * time.Second)
    go receiver(ch1, "X")
    time.Sleep(2 * time.Second)

    // Case 2: Small buffer (will cause blocking)
    fmt.Println("\nCase 2: Small Buffer (size 2)")
    ch2 := make(chan int, 2)
    go sender(ch2, "B")
    time.Sleep(500 * time.Millisecond)
    go receiver(ch2, "Y")
    time.Sleep(2 * time.Second)
} 