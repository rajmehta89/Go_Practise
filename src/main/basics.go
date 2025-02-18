package main

import "fmt"

// Person struct with basic info
type Person struct {
    Name    string
    Age     int
    Address Address
}

type Address struct {
    Street  string
    City    string
    Country string
}

// Value receiver method
func (p Person) GetInfo() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Pointer receiver method
func (p *Person) Birthday() {
    p.Age++
}

func BasicExample() {
    person := Person{
        Name: "Alice",
        Age:  25,
        Address: Address{
            Street:  "123 Main St",
            City:    "Tech City",
            Country: "Golang Land",
        },
    }

    fmt.Println(person.GetInfo())
    person.Birthday()
    fmt.Println(person.GetInfo())
} 