package main

import (
    "fmt"
    "sync"
    "time"
)

// Example 1: Fan-out pattern with error handling
func worker(id int, jobs <-chan int, results chan<- int, errors chan<- error) {
    defer func() {
        if r := recover(); r != nil {
            errors <- fmt.Errorf("worker %d panicked: %v", id, r)
        }
    }()
    
    for job := range jobs {
        if job%2 == 0 {
            results <- job * 2
        } else {
            panic(fmt.Sprintf("odd number not allowed: %d", job))
        }
    }
}

// Example 2: Pipeline with error handling
type Result struct {
    value int
    err   error
}

func generator(nums ...int) <-chan Result {
    out := make(chan Result)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- Result{value: n, err: nil}
        }
    }()
    return out
}

func multiply(in <-chan Result) <-chan Result {
    out := make(chan Result)
    go func() {
        defer close(out)
        for res := range in {
            if res.err != nil {
                out <- res
                continue
            }
            if res.value == 0 {
                out <- Result{0, fmt.Errorf("cannot multiply zero")}
                continue
            }
            out <- Result{value: res.value * 2, err: nil}
        }
    }()
    return out
}

// 1. Basic Channel Communication
func basicChannelExample() {
    fmt.Println("\n=== Basic Channel ===")
    ch := make(chan string)
    
    go func() {
        ch <- "Hello"  // Send
        close(ch)      // Close when done
    }()
    
    msg := <-ch       // Receive
    fmt.Println(msg)
}

// 2. Buffered Channel
func bufferedChannelExample() {
    fmt.Println("\n=== Buffered Channel ===")
    ch := make(chan int, 3)  // Buffer size 3
    
    // Can send 3 without blocking
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Start receiving in another goroutine
    go func() {
        for i := 0; i < 3; i++ {
            fmt.Printf("Received: %d\n", <-ch)
        }
    }()
    
    time.Sleep(100 * time.Millisecond)
}

// 3. Channel Direction
func sender(ch chan<- string) {   // Send-only channel
    ch <- "Message"
    close(ch)
}

func receiver(ch <-chan string) { // Receive-only channel
    msg := <-ch
    fmt.Printf("Received: %s\n", msg)
}

func channelDirectionExample() {
    fmt.Println("\n=== Channel Direction ===")
    ch := make(chan string)
    go sender(ch)
    receiver(ch)
}

// 4. Range over Channel
func rangeChannelExample() {
    fmt.Println("\n=== Range over Channel ===")
    ch := make(chan int)
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
        close(ch)  // Must close for range to work
    }()
    
    for num := range ch {
        fmt.Printf("Got: %d\n", num)
    }
}

// 5. Select with Multiple Channels
func selectExample() {
    fmt.Println("\n=== Select with Multiple Channels ===")
    ch1 := make(chan string)
    ch2 := make(chan string)
    done := make(chan bool)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "Message from ch1"
    }()
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "Message from ch2"
    }()
    
    go func() {
        time.Sleep(300 * time.Millisecond)
        done <- true
    }()
    
    for {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        case <-done:
            fmt.Println("Done")
            return
        default:
            fmt.Println("No message available")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

// 6. Fan-out Pattern
func fanOutExample() {
    fmt.Println("\n=== Fan-out Pattern ===")
    input := make(chan int)
    
    // Create multiple workers
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for num := range input {
                fmt.Printf("Worker %d processed %d\n", id, num)
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // Send work
    for i := 1; i <= 6; i++ {
        input <- i
    }
    close(input)
    
    wg.Wait()
}

// 7. Fan-in Pattern
func fanInExample() {
    fmt.Println("\n=== Fan-in Pattern ===")
    ch1 := make(chan string)
    ch2 := make(chan string)
    merged := make(chan string)
    
    // Merge channels
    go func() {
        var wg sync.WaitGroup
        wg.Add(2)
        
        output := func(c <-chan string) {
            defer wg.Done()
            for msg := range c {
                merged <- msg
            }
        }
        
        go output(ch1)
        go output(ch2)
        
        wg.Wait()
        close(merged)
    }()
    
    // Send data
    go func() {
        ch1 <- "1"
        ch1 <- "2"
        close(ch1)
    }()
    
    go func() {
        ch2 <- "A"
        ch2 <- "B"
        close(ch2)
    }()
    
    // Receive merged data
    for msg := range merged {
        fmt.Printf("Received: %s\n", msg)
    }
}

// 8. Pipeline Pattern
func pipelineExample() {
    fmt.Println("\n=== Pipeline Pattern ===")
    
    // Stage 1: Generate numbers
    nums := make(chan int)
    go func() {
        for i := 1; i <= 3; i++ {
            nums <- i
        }
        close(nums)
    }()
    
    // Stage 2: Square numbers
    squared := make(chan int)
    go func() {
        for n := range nums {
            squared <- n * n
        }
        close(squared)
    }()
    
    // Stage 3: Print results
    for n := range squared {
        fmt.Printf("Result: %d\n", n)
    }
}

func main() {
    fmt.Println("=== Fan-out Pattern Example ===")
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    errors := make(chan error, 5)
    
    // Start workers
    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            worker(id, jobs, results, errors)
        }(w)
    }
    
    // Send jobs
    go func() {
        for i := 0; i < 5; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Wait for workers in separate goroutine
    go func() {
        wg.Wait()
        close(results)
        close(errors)
    }()
    
    // Collect results and errors
    for {
        select {
        case res, ok := <-results:
            if !ok {
                results = nil
                continue
            }
            fmt.Printf("Result: %d\n", res)
        case err, ok := <-errors:
            if !ok {
                errors = nil
                continue
            }
            fmt.Printf("Error: %v\n", err)
        }
        if results == nil && errors == nil {
            break
        }
    }
    
    fmt.Println("\n=== Pipeline Pattern Example ===")
    // Create pipeline
    numbers := generator(0, 1, 2, 3, 4)
    doubled := multiply(numbers)
    
    // Collect results
    for res := range doubled {
        if res.err != nil {
            fmt.Printf("Pipeline error: %v\n", res.err)
        } else {
            fmt.Printf("Pipeline result: %d\n", res.value)
        }
    }

    basicChannelExample()
    bufferedChannelExample()
    channelDirectionExample()
    rangeChannelExample()
    selectExample()
    fanOutExample()
    fanInExample()
    pipelineExample()
} 