package main

import (
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2,3) = %d; expected %d", result, expected)
    }
} 