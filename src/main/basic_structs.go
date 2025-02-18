package main

import "fmt"

// Basic struct definition
type Person struct {
    FirstName string // Public field (capitalized)
    LastName  string
    age       int    // Private field (lowercase)
    email     string
}

// Constructor function
func NewPerson(firstName, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        age:       age,
    }
}

// Value receiver method
func (p Person) FullName() string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Pointer receiver method (can modify struct)
func (p *Person) Birthday() {
    p.age++
}

// Getter for private field
func (p Person) Age() int {
    return p.age
}

func main() {
    // Create using constructor
    person1 := NewPerson("Alice", "Smith", 30)
    
    // Create using struct literal
    person2 := Person{
        FirstName: "Bob",
        LastName:  "Jones",
        age:       25,
    }

    fmt.Printf("Person 1: %s, Age: %d\n", person1.FullName(), person1.Age())
    person1.Birthday()
    fmt.Printf("After birthday: %d\n", person1.Age())
	fmt.Println(person1.age)
    
    fmt.Printf("Person 2: %s, Age: %d\n", person2.FullName(), person2.Age())
} 