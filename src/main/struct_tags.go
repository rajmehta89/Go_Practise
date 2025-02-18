package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Username  string `json:"username" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Age       int    `json:"age" validate:"gte=0,lte=130"`
    Password  string `json:"-"`  // Won't be included in JSON
    Role      string `json:"role,omitempty"`
}

// Convert struct to JSON
func (u *User) ToJSON() (string, error) {
    bytes, err := json.Marshal(u)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// Convert JSON to struct
func FromJSON(jsonStr string) (*User, error) {
    var user User
    err := json.Unmarshal([]byte(jsonStr), &user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func StructTagsExample() {
    // Create a user
    originalUser := User{
        Username: "johndoe",
        Email:    "john@example.com",
        Age:      30,
        Password: "secret123",
        Role:     "user",
    }

    // Convert to JSON
    jsonStr, err := originalUser.ToJSON()
    if err != nil {
        fmt.Printf("Error marshaling: %v\n", err)
        return
    }
    fmt.Printf("JSON: %s\n", jsonStr)

    // Convert back to struct
    newUser, err := FromJSON(jsonStr)
    if err != nil {
        fmt.Printf("Error unmarshaling: %v\n", err)
        return
    }

    // Compare values
    fmt.Printf("\nOriginal User: %+v\n", originalUser)
    fmt.Printf("New User from JSON: %+v\n", newUser)
    
    // Note that Password field is not in JSON
    fmt.Printf("\nPassword field comparison:\n")
    fmt.Printf("Original password: %s\n", originalUser.Password)
    fmt.Printf("New password (empty): %s\n", newUser.Password)
} 