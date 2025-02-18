package main

import (
    "fmt"
    "time"
)

// Example 1: Basic Defer and Panic
func basicDeferExample() {
    defer fmt.Println("1. This runs last")
    defer fmt.Println("2. This runs second to last")
    
    fmt.Println("3. This runs first")
    
    // Panic example (commented out)
    // panic("Something went wrong!")
}

// Example 2: Defer with Panic and Recover
func recoverExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    fmt.Println("Starting dangerous operation...")
    panic("Something went terribly wrong!")
    fmt.Println("This line never executes")
}


// func realWorldExample() {
// //     // Open a file
// //     file := openFile()
// //     defer file.Close()  // Will run evsen if panic occurs

// //     if err := processFile(); err != nil {
// //         panic("File processing failed!")  // Stop execution
// //         return                           // Never reaches this
// //     }
// }


func example() {
    // Order of execution:
    defer fmt.Println("1. Runs third (defer LIFO)")
    defer fmt.Println("2. Runs second")
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("3. Runs first - recovers from panic")
        }
    }()

    fmt.Println("4. Runs normally")
    panic("Panic!")                    // Execution stops here
    fmt.Println("Never runs")          // This line is skipped
}
// Example 3: Channel with Panic and Recover
func safeChannelSend(ch chan int, value int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from channel send: %v\n", r)
        }
    }()
    
    ch <- value
}

// Example 4: Complex Error Handling
type SafeChannel struct {
    ch    chan int
    done  chan bool
    error chan error
}

func NewSafeChannel() *SafeChannel {
    return &SafeChannel{
        ch:    make(chan int),
        done:  make(chan bool),
        error: make(chan error),
    }
}

func (sc *SafeChannel) SafeSend(value int) {
    defer func() {
        if r := recover(); r != nil {
            sc.error <- fmt.Errorf("send panic: %v", r)
        }
    }()
    
    sc.ch <- value
}

func (sc *SafeChannel) SafeReceive() (int, error) {
    defer func() {
        if r := recover(); r != nil {
            sc.error <- fmt.Errorf("receive panic: %v", r)
        }
    }()
    
    select {
    case val := <-sc.ch:
        return val, nil
    case err := <-sc.error:
        return 0, err
    case <-time.After(time.Second):
        return 0, fmt.Errorf("timeout")
    }
}

func main() {
    fmt.Println("=== Basic Defer Example ===")
    basicDeferExample()
    
    fmt.Println("\n=== Recover Example ===")
    recoverExample()

    fmt.Println("examplemof the recover")
    example()

    // fmt.Println("real world example")
    // realWorldExample()
    
    fmt.Println("\n=== Channel with Panic/Recover Example ===")
    ch := make(chan int)
    close(ch)
    safeChannelSend(ch, 42) // This would panic without recover
    
    fmt.Println("\n=== Safe Channel Example ===")
    sc := NewSafeChannel()
    
    // Start sender
    go func() {
        for i := 1; i <= 5; i++ {
            fmt.Printf("Sending: %d\n", i)
            sc.SafeSend(i)
            time.Sleep(200 * time.Millisecond)
        }
        close(sc.ch)
    }()
    
    // Receive values
    for i := 0; i < 6; i++ { // Try to receive one extra
        if val, err := sc.SafeReceive(); err != nil {
            fmt.Printf("Error receiving: %v\n", err)
            break
        } else {
            fmt.Printf("Received: %d\n", val)
        }
    }
} 