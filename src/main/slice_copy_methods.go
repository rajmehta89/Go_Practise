package main

import (
	"fmt"
	"reflect"
)

func printSliceInfo(name string, slice []int) {
	fmt.Printf("%s: value=%v, addr=%p, len=%d, cap=%d\n", 
		name, slice, &slice[0], len(slice), cap(slice))
}

func demonstrateSliceCopyMethods() {
	// Original slice
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("\n=== Original Slice ===")
	printSliceInfo("Original", original)

	// 1. Using make + copy
	fmt.Println("\n=== Method 1: make + copy ===")
	copy1 := make([]int, len(original))
	copy(copy1, original)
	printSliceInfo("Copy1", copy1)

	// 2. Using append with empty slice
	fmt.Println("\n=== Method 2: append to empty slice ===")
	copy2 := append([]int{}, original...)
	printSliceInfo("Copy2", copy2)

	// 3. Using append + make
	fmt.Println("\n=== Method 3: append + make ===")
	copy3 := append(make([]int, 0, len(original)), original...)
	printSliceInfo("Copy3", copy3)

	// 4. Using slice of slice + full capacity
	fmt.Println("\n=== Method 4: full slice + make ===")
	copy4 := make([]int, len(original))
	copy4 = original[:len(original):len(original)] // Full slice with limited capacity
	printSliceInfo("Copy4", copy4)

	// 5. Using reflect
	fmt.Println("\n=== Method 5: reflect ===")
	copy5 := reflect.ValueOf(original).Interface().([]int)
	printSliceInfo("Copy5", copy5)

	// 6. Manual copy
	fmt.Println("\n=== Method 6: manual copy ===")
	copy6 := make([]int, len(original))
	for i, v := range original {
		copy6[i] = v
	}
	printSliceInfo("Copy6", copy6)

	// Demonstrate independence
	fmt.Println("\n=== Testing Independence ===")
	original[0] = 100
	fmt.Printf("After modifying original[0] to 100:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copy1: %v\n", copy1)
	fmt.Printf("Copy2: %v\n", copy2)
	fmt.Printf("Copy3: %v\n", copy3)
	fmt.Printf("Copy4: %v\n", copy4)
	fmt.Printf("Copy5: %v\n", copy5)
	fmt.Printf("Copy6: %v\n", copy6)

	// Deep copy of slice of slices
	fmt.Println("\n=== Deep Copy of Nested Slices ===")
	nested := [][]int{{1, 2}, {3, 4}}
	
	// Shallow copy
	shallowCopy := append([][]int{}, nested...)
	
	// Deep copy
	deepCopy := make([][]int, len(nested))
	for i, innerSlice := range nested {
		deepCopy[i] = append([]int{}, innerSlice...)
	}

	// Test nested modification
	nested[0][0] = 999
	fmt.Printf("After modifying nested[0][0]:\n")
	fmt.Printf("Original: %v\n", nested)
	fmt.Printf("Shallow: %v\n", shallowCopy)
	fmt.Printf("Deep: %v\n", deepCopy)
}

func main() {
	demonstrateSliceCopyMethods()
}