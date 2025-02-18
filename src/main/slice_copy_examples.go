package main

import "fmt"

func demonstrateSliceCopies() {
    // Original slice
    original := []int{1, 2, 3}
    fmt.Printf("Original: %v, addr: %p\n", original, &original[0])

    // Case 1: Creates a new copy
    copySlice := make([]int, len(original))
    copy(copySlice, original)
    fmt.Printf("Copy: %v, addr: %p\n", copySlice, &copySlice[0])

    // Case 2: Same underlying array
    sliceView := original[:]
    fmt.Printf("View: %v, addr: %p\n", sliceView, &sliceView[0])

    // Case 3: append() might create new array
    smallSlice := make([]int, 0, 2)
    fmt.Printf("\nSmall slice initial addr: %p\n", &smallSlice)
    
    smallSlice = append(smallSlice, 1)
    fmt.Printf("After append 1: %v, addr: %p\n", smallSlice, &smallSlice[0])
    
    smallSlice = append(smallSlice, 2, 3) // Exceeds capacity, creates new array
    fmt.Printf("After append 2,3: %v, addr: %p\n", smallSlice, &smallSlice[0])

    // Case 4: Function parameters
    modifySlice(original)
    fmt.Printf("\nAfter function call: %v\n", original)

    // Case 5: Using append(...) to copy
    newCopy := append([]int{}, original...)
    fmt.Printf("\nNew copy addr: %p\n", &newCopy[0])
    fmt.Printf("Original addr: %p\n", &original[0])
}

func modifySlice(s []int) {
    fmt.Printf("In function addr: %p\n", &s[0])
    s[0] = 100 // Modifies original array
    s = append(s, 4) // Creates new array, but original unchanged
}

// Case 6: Different ways to create slices
func differentSliceCreations() {
    fmt.Println("\n=== Different Slice Creations ===")
    
    // Direct creation - single underlying array
    s1 := []int{1, 2, 3}
    s2 := s1
    fmt.Printf("s1 addr: %p\n", &s1[0])
    fmt.Printf("s2 addr: %p\n", &s2[0])

    // make() - new underlying array
    s3 := make([]int, 3)
    copy(s3, s1)
    fmt.Printf("s3 addr: %p\n", &s3[0])

    // append(...) - new underlying array
    s4 := append([]int{}, s1...)
    fmt.Printf("s4 addr: %p\n", &s4[0])
}

func main() {
    demonstrateSliceCopies()
    differentSliceCreations()

    // Practical example
    fmt.Println("\n=== Practical Example ===")
    data := []int{1, 2, 3}
    
    // These share the same underlying array
    fmt.Printf("Original data addr: %p\n", &data[0])
    
    defer func(slice []int) {
        fmt.Printf("Defer view addr: %p, value: %v\n", &slice[0], slice)
    }(data)
    
    defer func(slice []int) {
        fmt.Printf("Defer copy addr: %p, value: %v\n", &slice[0], slice)
    }(append([]int{}, data...))
    
    data[0] = 100
    data = append(data, 4) // Creates new array
    fmt.Printf("Modified data addr: %p, value: %v\n", &data[0], data)
} 