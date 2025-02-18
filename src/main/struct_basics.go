package main

import "fmt"

// Basic struct definition
type Person struct {
    FirstName string    // Public field (capitalized)
    LastName  string
    age       int      // Private field (lowercase)
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

// Pointer receiver method
func (p *Person) UpdateAge(newAge int) {
    p.age = newAge
}

func main() {
    // Different ways to create structs
    person1 := Person{FirstName: "John", LastName: "Doe", age: 30}
    person2 := new(Person)  // All fields zero-valued
	fmt.Println(person2)
	fmt.Println(person2.FirstName)
	fmt.Println(person2.LastName)
	fmt.Println(person2.age)
	fmt.Println(person2.email)

	var s string
	fmt.Println("printing string here")
	fmt.Println(s)

	var s1 int
		fmt.Println("printing int here")
	fmt.Println(s1)

	var s2 float64
	fmt.Println("printing float64 here")
	fmt.Println(s2)

	var s3 bool
	fmt.Println("printing bool here")
	fmt.Println(s3)

	var s4 complex64
	fmt.Println("printing complex64 here")
	fmt.Println(s4)

	var s5 []int
	fmt.Println("printing slice here")
	fmt.Println(s5)

	var s6 map[string]int
	fmt.Println("printing map here")
	fmt.Println(s6)

	var s7 chan int
	fmt.Println("printing channel here")
	fmt.Println(s7)

	var s8 func()
	fmt.Println("printing function here")
	fmt.Println(s8)

	var s9 interface{}
	fmt.Println("printing interface here")
	fmt.Println(s9)

	var s10 *int
	fmt.Println("printing pointer here")
	fmt.Println(s10)

	var s11 error
	fmt.Println("printing error here")
	fmt.Println(s11)

	var s12 []int
	fmt.Println("printing slice here")
	fmt.Println(s12)

	var s13 map[string]int
	fmt.Println("printing map here")
	fmt.Println(s13)



    person3 := &Person{FirstName: "Jane"}  // Partial initialization
    person4 := NewPerson("Bob", "Smith", 25)  // Using constructor

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
} 