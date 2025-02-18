package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Write file
    data := []byte("Hello, Go!")
    err := ioutil.WriteFile("test.txt", data, 0644)
    if err != nil {
        fmt.Printf("Write error: %v\n", err)
        return
    }

    // Read file
    content, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Printf("Read error: %v\n", err)
        return
    }
    fmt.Printf("File content: %s\n", content)

    // Clean up
    os.Remove("test.txt")
} 