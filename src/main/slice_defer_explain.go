package main

import "fmt"

func explainDeferSliceBehavior() {
    // Original slice
    s := []string{"original"}
    
    // Case 1: By Value - Makes a copy of the slice at defer time
    defer func(slice []string) {
        fmt.Printf("By Value: %v\n", slice)
    }(s)  // Passing 's' as an argument creates a copy
    
    // Case 2: By Closure - Captures reference to 's' variable
    defer func() {
        fmt.Printf("By Closure: %v\n", s)
    }()  // No arguments, uses 's' from outer scope
    
    // Modify the slice
    s[0] = "modified"        // Modifies existing element
    s = append(s, "appended") // Creates new underlying array
}

func main() {
    fmt.Println("=== Defer with Slice Behavior ===")
    explainDeferSliceBehavior()
} 