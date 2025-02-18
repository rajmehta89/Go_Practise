package main

import "fmt"

type Resource struct {
    name string
}

func (r *Resource) Close() {
    fmt.Printf("Closing resource: %s\n", r.name)
}

// Example 1: Basic defer order
func basicDeferOrder() {
    fmt.Println("\n=== Basic Defer Order ===")
    defer fmt.Println("1. First defer")
    defer fmt.Println("2. Second defer")
    defer fmt.Println("3. Third defer")
    fmt.Println("4. Normal execution")
}

// Example 2: Defer in loop
func deferInLoop() {
    fmt.Println("\n=== Defer in Loop ===")
    for i := 0; i < 3; i++ {
        defer fmt.Printf("Defer in loop: %d\n", i)
		fmt.Println("i is", i)
    }
    fmt.Println("Loop finished")
}

// Example 3: Defer with method calls
func deferWithMethods() {
    fmt.Println("\n=== Defer with Methods ===")
    r1 := &Resource{name: "DB"}
    r2 := &Resource{name: "File"}
    
    defer r1.Close()  // Will close last
    defer r2.Close()  // Will close first
    
    fmt.Println("Working with resources...")
}

// Example 4: Defer evaluation
func deferEvaluation() {
    fmt.Println("\n=== Defer Evaluation ===")
    var s []int = []int{1,2,34}
	s[0] = 10
    defer fmt.Printf("Deferred value of i: %d\n", s) // Value captured when defer is called
    s[0]=100
    fmt.Printf("Final value of i: %d\n", s[0])
}

// Example 5: Defer in nested functions
func deferNested() {
    fmt.Println("\n=== Defer in Nested Functions ===")
    defer fmt.Println("1. Outer defer")
    
    func() {
        defer fmt.Println("2. Inner defer")
        fmt.Println("3. Inner execution")
    }()
    
    fmt.Println("4. Outer execution")
}

// Example 6: Defer with panic
func deferWithPanic() {
    fmt.Println("\n=== Defer with Panic ===")
    defer fmt.Println("1. This runs even after panic")
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("2. Recovered from: %v\n", r)
        }
    }()
    
    fmt.Println("3. About to panic")
    panic("Something went wrong!")
    fmt.Println("4. This never runs")
}

func main() {
    basicDeferOrder()
    deferInLoop()
    deferWithMethods()
    deferEvaluation()
    deferNested()
    
    // Run panic example in separate function to continue execution
    func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Main recovered from panic")
            }
        }()
        deferWithPanic()
    }()
    
    fmt.Println("\nProgram completed")
} 