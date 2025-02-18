package main

import (
    "fmt"
    "sync"
    "time"
)

type Task struct {
    ID       int
    Duration time.Duration
}

// Each worker has its own output channel
func worker(id int, tasks <-chan Task, output chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    defer close(output) // Each worker closes its own output channel

    for task := range tasks {
        startTime := time.Now()
        fmt.Printf("Worker %d started task %d at %s\n", 
            id, task.ID, startTime.Format("15:04:05.000"))
        
        time.Sleep(task.Duration)
        output <- fmt.Sprintf("Worker %d completed task %d", id, task.ID)
    }
}

func main() {
    tasks := []Task{
        {ID: 1, Duration: 500 * time.Millisecond},
        {ID: 2, Duration: 300 * time.Millisecond},
        {ID: 3, Duration: 400 * time.Millisecond},
        {ID: 4, Duration: 200 * time.Millisecond},
    }

    // Input channel
    taskChan := make(chan Task, len(tasks))

    // Create separate output channel for each worker
    numWorkers := 3
    outputs := make([]chan string, numWorkers)
    var wg sync.WaitGroup

    // Start workers with their own output channels
    for i := 0; i < numWorkers; i++ {
        outputs[i] = make(chan string, len(tasks)/numWorkers+1)
        wg.Add(1)
        go worker(i+1, taskChan, outputs[i], &wg)
    }

    // Send tasks
    go func() {
        for _, task := range tasks {
            taskChan <- task
        }
        close(taskChan)
    }()

    // Fan-in: Collect results from all output channels
    go func() {
        wg.Wait() // Wait for all workers to finish
    }()

    // Read from all output channels
    for i, ch := range outputs {
        fmt.Printf("\nResults from Worker %d:\n", i+1)
        for result := range ch {
            fmt.Printf("%s\n", result)
        }
    }
} 