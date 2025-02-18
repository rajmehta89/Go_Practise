package main

import (
    "fmt"
    "sync"
    "time"
)

type ImageTask struct {
    ID       int
    Size     int
    Priority int
}

func processImage(task ImageTask) string {
    // Simulate image processing
    time.Sleep(time.Duration(task.Size*50) * time.Millisecond)
    return fmt.Sprintf("Processed image %d (size: %d)", task.ID, task.Size)
}

func imageWorker(id int, tasks <-chan ImageTask, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for task := range tasks {
        fmt.Printf("Worker %d starting image %d\n", id, task.ID)
        result := processImage(task)
        results <- result
    }
}

func main() {
    tasks := make(chan ImageTask, 10)
    results := make(chan string, 10)
    
    // Start worker pool
    numWorkers := 3
    var wg sync.WaitGroup
    
    // Fan out to workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go imageWorker(i, tasks, results, &wg)
    }
    
    // Send image processing tasks
    go func() {
        for i := 1; i <= 6; i++ {
            task := ImageTask{
                ID:       i,
                Size:     i * 100,
                Priority: i % 3,
            }
            tasks <- task
        }
        close(tasks)
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
} 