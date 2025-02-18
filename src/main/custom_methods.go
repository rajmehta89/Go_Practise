package main

type Counter struct {
    value int
    max   int
}

// Value receiver (doesn't modify struct)
func (c Counter) GetValue() int {
    return c.value
}

// Pointer receiver (modifies struct)
func (c *Counter) Increment() bool {
    if c.value < c.max {
        c.value++
        return true
    }
    return false
}

// Factory function
func NewCounter(max int) *Counter {
    return &Counter{
        value: 0,
        max:   max,
    }
} 