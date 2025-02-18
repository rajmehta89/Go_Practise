package main

// Address struct
type Address struct {
    Street  string
    City    string
    Country string
}

// Contact details
type Contact struct {
    Email  string
    Phone  string
}

// Employee with nested structs
type Employee struct {
    Person            // Embedded struct (inheritance-like)
    HomeAddress    Address    // Nested struct
    WorkAddress    Address
    ContactInfo    Contact
    Position      string
    Salary        float64
}

// Usage example
func StructComposition() {
    emp := Employee{
        Person: Person{
            FirstName: "John",
            LastName:  "Doe",
        },
        HomeAddress: Address{
            Street: "123 Home St",
            City:   "Hometown",
        },
        Position: "Developer",
    }
    
    // Access fields
    fmt.Println(emp.FirstName)        // From embedded Person
    fmt.Println(emp.HomeAddress.City) // From nested Address
} 