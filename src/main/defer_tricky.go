package main

import (
    "fmt"
    "sync"
)

// Scenario 1: Defer with return values
func deferWithReturn() (result int) {
    defer func() {
        result++  // This modifies the named return value
        fmt.Printf("Defer modified result to: %d\n", result)
    }()
    return 1  // Returns 2 because defer modifies it
}

// Scenario 2: Defer with pointer vs value
type Counter struct {
    count int
}

func deferWithPointer() {
    c := &Counter{count: 0}
    
    // Using pointer
    defer func(counter *Counter) {
        counter.count++
        fmt.Printf("Pointer defer: count = %d\n", counter.count)
    }(c)
    
    // Using value
    defer func(counter Counter) {
        counter.count++
        fmt.Printf("Value defer: count = %d\n", counter.count)
    }(*c)
    
    c.count = 5
    fmt.Printf("Final count before defers: %d\n", c.count)
}

// Scenario 3: Defer in goroutines
func deferInGoroutine() {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            defer fmt.Printf("Goroutine %d defer\n", id)
            fmt.Printf("Goroutine %d executing\n", id)
        }(i)
    }
    
    wg.Wait()
}

// Scenario 4: Defer with method value vs method expression
type Resource struct {
    name string
}

func (r *Resource) CloseValue() {
    fmt.Printf("Closing %s (value)\n", r.name)
}

func (r *Resource) CloseExpr() {
    fmt.Printf("Closing %s (expr)\n", r.name)
}

func deferWithMethods() {
    r := &Resource{name: "resource"}
    
    // Method value - captures receiver when defer is called
    defer r.CloseValue()
    
    // Method expression - captures receiver when method is called
    defer (*Resource).CloseExpr(r)
    
    r.name = "modified resource"
}

// Scenario 5: Defer with slice modification
func deferWithSlice() {
    s := []string{"original"}
    
    defer func(slice []string) {
        fmt.Printf("Slice in defer (by value): %v\n", slice)
    }(s)
    
    defer func() {
        fmt.Printf("Slice in defer (by closure): %v\n", s)
    }()
    
    s[0] = "modified"
    s = append(s, "appended")
}

func main() {
    fmt.Println("\n=== Defer with Return ===")
    fmt.Printf("Returned value: %d\n", deferWithReturn())
    
    fmt.Println("\n=== Defer with Pointer vs Value ===")
    deferWithPointer()
    
    fmt.Println("\n=== Defer in Goroutines ===")
    deferInGoroutine()
    
    fmt.Println("\n=== Defer with Methods ===")
    deferWithMethods()
    
    fmt.Println("\n=== Defer with Slice ===")
    deferWithSlice()
} 