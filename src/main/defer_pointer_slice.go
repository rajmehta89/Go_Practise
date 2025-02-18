package main

import "fmt"

func deferPointerExample() {
    // Pointer example
    c := &Counter{count: 0}
    
    // Using pointer - will see final value
    defer func(counter *Counter) {
        fmt.Printf("Pointer defer: count = %d\n", counter.count)
    }(c)
    
    c.count = 5
    fmt.Printf("Final count: %d\n", c.count)
}

func deferSliceExample() {
    // Slice example
    s := []int{0}
    
    // Using slice - will see final value
    defer func(slice []int) {
        fmt.Printf("Slice by pointer: %v\n", slice)
    }(s)
    
    // Using slice value - won't see changes
    defer func(slice []int) {
        fmt.Printf("Slice by value: %v\n", slice)
    }(append([]int{}, s...))  // Make a copy
    
    s[0] = 5        // Modifies underlying array
    s = append(s, 10)  // Creates new underlying array
    fmt.Printf("Final slice: %v\n", s)
}

func main() {
    fmt.Println("=== Pointer Behavior ===")
    deferPointerExample()
    
    fmt.Println("\n=== Slice Behavior ===")
    deferSliceExample()
} 