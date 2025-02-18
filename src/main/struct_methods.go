package main

import (
    "fmt"
    "math"
)

// Interface definition
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle implements Shape
type Circle struct {
    Radius float64
}

// Rectangle methods
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle methods
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Generic shape processor
func ProcessShape(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func StructMethodsExample() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 5}

    fmt.Println("Rectangle:")
    ProcessShape(rect)

    fmt.Println("\nCircle:")
    ProcessShape(circle)
} 