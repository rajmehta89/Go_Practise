package main

import (
    "fmt"
    "sync"
    "time"
)

// Example 1: Fan-out/Fan-in Pattern
func worker(id int, jobs <-chan int, results chan<- string) {
    defer func() {
        if r := recover(); r != nil {
            results <- fmt.Sprintf("Worker %d error: %v", id, r)
        }
    }()
    
    for job := range jobs {
        if job%3 == 0 {
            panic(fmt.Sprintf("Worker %d: Can't process job %d", id, job))
        }
        results <- fmt.Sprintf("Worker %d processed job %d", id, job)
        time.Sleep(100 * time.Millisecond)
    }
}

// Example 2: Pipeline Pattern
type Data struct {
    value int
    err   error
}

func generator(done chan bool) <-chan Data {
    out := make(chan Data)
    go func() {
        defer close(out)
        for i := 1; i <= 5; i++ {
            select {
            case out <- Data{i, nil}:
            case <-done:
                return
            }
        }
    }()
    return out
}

func processor(done chan bool, in <-chan Data) <-chan Data {
    out := make(chan Data)
    go func() {
        defer close(out)
        for data := range in {
            if data.err != nil {
                out <- data
                continue
            }
            
            select {
            case out <- Data{data.value * 2, nil}:
            case <-done:
                return
            }
        }
    }()
    return out
}

func main() {
    fmt.Println("=== Fan-out/Fan-in Example ===")
    jobs := make(chan int, 10)
    results := make(chan string, 10)
    
    // Start workers
    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            worker(id, jobs, results)
        }(w)
    }
    
    // Send jobs
    go func() {
        for i := 1; i <= 9; i++ {
            jobs <- i
        }
        close(jobs)
    }()
    
    // Wait and close results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
    
    fmt.Println("\n=== Pipeline Example ===")
    done := make(chan bool)
    defer close(done)
    
    // Create pipeline
    input := generator(done)
    output := processor(done, input)
    
    // Collect results with timeout
    timeout := time.After(2 * time.Second)
    for {
        select {
        case data, ok := <-output:
            if !ok {
                return
            }
            if data.err != nil {
                fmt.Printf("Error: %v\n", data.err)
            } else {
                fmt.Printf("Result: %d\n", data.value)
            }
        case <-timeout:
            fmt.Println("Pipeline timeout")
            return
        }
    }
} 