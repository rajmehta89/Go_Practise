package main

import (
   
    "fmt"
)

// Custom error
type DivisionError struct {
    dividend float64
    divisor  float64
}

func (e *DivisionError) Error() string {
    return fmt.Sprintf("cannot divide %v by %v", e.dividend, e.divisor)
}

// Function that returns error
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, &DivisionError{x, y}
    }
    return x / y, nil
}

func main() {
    // Error handling example
    result, err := divide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Result: %v\n", result)
} 