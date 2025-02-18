package main

import (
    "fmt"
    "sync"
    "time"
)

// Task represents work to be done
type Task struct {
    ID       int
    Type     string    // "cpu", "io", or "network"
    Duration time.Duration
}

// Result represents processed task result
type Result struct {
    TaskID    int
    WorkerID  int
    StartTime time.Time
    EndTime   time.Time
    Output    string
}

// Worker processes tasks in parallel
func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()

    for task := range tasks {
        // Record start time
        startTime := time.Now()
        
        fmt.Printf("Worker %d started  task %d at %s\n", 
            id, task.ID, startTime.Format("15:04:05.000"))
        
        // Simulate processing
        time.Sleep(task.Duration)
        
        // Record completion
        endTime := time.Now()
        
        results <- Result{
            TaskID:    task.ID,
            WorkerID:  id,
            StartTime: startTime,
            EndTime:   endTime,
            Output:    fmt.Sprintf("%s Result for task %d", task.Type, task.ID),
        }
        
        fmt.Printf("Worker %d finished task %d at %s\n", 
            id, task.ID, endTime.Format("15:04:05.000"))
    }
}

func main() {
    // Create tasks with different durations
    tasks := []Task{
        {ID: 1, Type: "cpu", Duration: 500 * time.Millisecond},
        {ID: 2, Type: "io", Duration: 300 * time.Millisecond},
        {ID: 3, Type: "network", Duration: 400 * time.Millisecond},
        {ID: 4, Type: "cpu", Duration: 500 * time.Millisecond},
        {ID: 5, Type: "io", Duration: 300 * time.Millisecond},
        {ID: 6, Type: "network", Duration: 400 * time.Millisecond},
    }

    taskChan := make(chan Task, len(tasks))
    resultChan := make(chan Result, len(tasks))

    startTime := time.Now()
    fmt.Printf("Starting processing at: %s\n", startTime.Format("15:04:05.000"))

    // Start multiple workers
    numWorkers := 3
    var wg sync.WaitGroup

    // Launch workers in parallel
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, taskChan, resultChan, &wg)  // Workers run concurrently
    }

    // Send tasks
    go func() {
        for _, task := range tasks {
            taskChan <- task
        }
        close(taskChan)
    }()

    // Wait for workers
    go func() {
        wg.Wait()
        close(resultChan)
    }()

    // Collect and analyze results
    var results []Result
    for result := range resultChan {
        results = append(results, result)
    }

    // Show parallel execution evidence
    fmt.Println("\nParallel Execution Analysis:")
    fmt.Printf("Total tasks: %d\n", len(tasks))
    fmt.Printf("Number of workers: %d\n", numWorkers)
    
    // Calculate total sequential time vs actual time
    totalSequentialTime := time.Duration(0)
    for _, task := range tasks {
        totalSequentialTime += task.Duration
    }
    
    actualTime := time.Since(startTime)
    
    fmt.Printf("\nSequential would take: %v\n", totalSequentialTime)
    fmt.Printf("Actual parallel time: %v\n", actualTime)
    fmt.Printf("Speed improvement: %.2fx\n", 
        float64(totalSequentialTime)/float64(actualTime))

    // Show overlapping executions
    fmt.Println("\nTask Execution Timeline:")
    for _, r := range results {
        fmt.Printf("Worker %d: Task %d - Start: %s, End: %s\n",
            r.WorkerID, r.TaskID,
            r.StartTime.Format("15:04:05.000"),
            r.EndTime.Format("15:04:05.000"))
    }
} 