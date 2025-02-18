package main

import "fmt"

// Interface definition
type Animal interface {
    Speak() string
    Move() string
}

// Implementing structs
type Dog struct {
    Name string
}

type Cat struct {
    Name string
}

// Dog methods
func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running"
}

// Cat methods
func (c Cat) Speak() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Prowling"
}

func main() {
    animals := []Animal{
        Dog{Name: "Rover"},
        Cat{Name: "Whiskers"},
    }

    for _, animal := range animals {
        fmt.Printf("Animal says: %s and is %s\n", 
            animal.Speak(), 
            animal.Move())
    }
}

type Processor interface {
    Process() string
    GetName() string
}

type TextProcessor struct {
    Name string
    Text string
}

func (tp TextProcessor) Process() string {
    return fmt.Sprintf("Processing text: %s", tp.Text)
}

func (tp TextProcessor) GetName() string {
    return tp.Name
}

type NumberProcessor struct {
    Name   string
    Number int
}

func (np NumberProcessor) Process() string {
    return fmt.Sprintf("Processing number: %d", np.Number)
}

func (np NumberProcessor) GetName() string {
    return np.Name
}

func InterfaceExample() {
    processors := []Processor{
        TextProcessor{Name: "Text Proc", Text: "Hello"},
        NumberProcessor{Name: "Num Proc", Number: 42},
    }

    for _, p := range processors {
        fmt.Printf("Processor %s: %s\n", p.GetName(), p.Process())
    }
} 