package main

import (
    "fmt"
    "runtime/debug"
    "sync"
    "time"
)

// Custom error types
type ValidationError struct {
    Field string
    Issue string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("Validation error on %s: %s", e.Field, e.Issue)
}

type DatabaseError struct {
    Operation string
    Err       error
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("Database %s failed: %v", e.Operation, e.Err)
}

// Advanced recovery function with stack trace and error classification
func advancedRecover(operation string) func() {
    return func() {
        if r := recover(); r != nil {
            // Get stack trace
            stack := debug.Stack()
            
            // Classify the panic
            switch err := r.(type) {
            case *ValidationError:
                fmt.Printf("Validation error during %s: %v\n", operation, err)
            case *DatabaseError:
                fmt.Printf("Database error during %s: %v\n", operation, err)
            case error:
                fmt.Printf("Error during %s: %v\n", operation, err)
            default:
                fmt.Printf("Panic during %s: %v\n", operation, r)
            }
            
            fmt.Printf("Stack trace:\n%s\n", string(stack))
        }
    }
}

// Nested function with panic
func processUserData(userData map[string]string) {
    defer advancedRecover("user data processing")()
    
    if userData == nil {
        panic(&ValidationError{Field: "userData", Issue: "nil map"})
    }
    
    // Simulate nested function call
    validateUser(userData)
}

func validateUser(data map[string]string) {
    if data["username"] == "" {
        panic(&ValidationError{Field: "username", Issue: "empty username"})
    }
}

// Goroutine panic handling
func workerWithPanic(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    defer advancedRecover(fmt.Sprintf("worker %d", id))()
    
    // Simulate different panic scenarios
    switch id % 3 {
    case 0:
        panic(&ValidationError{Field: "worker", Issue: fmt.Sprintf("worker %d validation failed", id)})
    case 1:
        panic(&DatabaseError{Operation: "query", Err: fmt.Errorf("connection lost")})
    case 2:
        // Simulate panic with non-error value
        panic(fmt.Sprintf("unexpected error in worker %d", id))
    }
}

// Resource cleanup with panic handling
type Resource struct {
    name string
    data []byte
}

func (r *Resource) Close() {
    fmt.Printf("Cleaning up resource: %s\n", r.name)
    r.data = nil
}

func processResourceWithPanic(name string) (err error) {
    // Create resource
    resource := &Resource{name: name, data: make([]byte, 1000)}
    
    // Ensure cleanup happens even in case of panic
    defer func() {
        if r := recover(); r != nil {
            resource.Close()
            err = fmt.Errorf("process failed: %v", r)
        }
    }()
    
    // Simulate processing that might panic
    if len(name) < 5 {
        panic(fmt.Sprintf("invalid resource name: %s", name))
    }
    
    return nil
}

func main() {
    // fmt.Println("=== Testing Nested Panic Recovery ===")
    // // Test case 1: nil map
    // processUserData(nil)
    
    // // Test case 2: invalid data
    // processUserData(map[string]string{
    //     "email": "test@example.com",
    //     // missing username
    // })
    
    // fmt.Println("\n=== Testing Goroutine Panic Recovery ===")
    // var wg sync.WaitGroup
    // for i := 0; i < 3; i++ {
    //     wg.Add(1)
    //     go workerWithPanic(i, &wg)
    // }
    // wg.Wait()
    
    // fmt.Println("\n=== Testing Resource Cleanup ===")
    // // Test resource cleanup
    // if err := processResourceWithPanic("ok"); err != nil {
    //     fmt.Printf("Error: %v\n", err)
    // }
    // if err := processResourceWithPanic("bad"); err != nil {
    //     fmt.Printf("Error: %v\n", err)
    // }
    
    // // Demonstrate panic in deferred function
    // fmt.Println("\n=== Testing Panic in Defer ===")
    func() {
        defer advancedRecover("outer")()
        defer func() {
            defer advancedRecover("inner")()
            fmt.Println("inner defer")
            defer func() {
                fmt.Println("inner inner defer")
            }()
            panic("panic in defer")
        }()
        fmt.Println("Normal execution before defer panic")
    }()
    
    time.Sleep(time.Second) // Wait for goroutines to finish
    fmt.Println("\nProgram completed successfully")
} 