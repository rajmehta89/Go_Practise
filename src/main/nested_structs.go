package main

import "fmt"

// Address struct
type Address struct {
    Street  string
    City    string
    Country string
    Postal  string
}

// Contact information
type Contact struct {
    Email     string
    Phone     string
    Emergency struct {
        Name  string
        Phone string
    }
}

// Employee with nested structs
type Employee struct {
    Person            // Embedded struct
    HomeAddress    Address
    WorkAddress    Address
    ContactInfo    Contact
    Position      string
    Salary        float64
}

func (e Employee) DisplayInfo() {
    fmt.Printf("Employee: %s\n", e.FullName())
    fmt.Printf("Position: %s\n", e.Position)
    fmt.Printf("Home: %s, %s\n", e.HomeAddress.Street, e.HomeAddress.City)
    fmt.Printf("Work: %s, %s\n", e.WorkAddress.Street, e.WorkAddress.City)
}

func NestedStructExample() {
    emp := Employee{
        Person: Person{
            FirstName: "John",
            LastName:  "Doe",
            age:       30,
        },
        HomeAddress: Address{
            Street:  "123 Home St",
            City:    "Hometown",
            Country: "USA",
        },
        WorkAddress: Address{
            Street:  "456 Office Blvd",
            City:    "Worktown",
            Country: "USA",
        },
        Position: "Developer",
        Salary:   75000,
    }

    emp.ContactInfo.Emergency.Name = "Jane Doe"
    emp.ContactInfo.Emergency.Phone = "555-0123"

    emp.DisplayInfo()
} 