package main

import (
	"fmt"
	"time"
)

type any interface {}

func acceptanything(a any) {
	fmt.Println("here i am accepting anythin which i want to accept here")
    fmt.Println(a)
}

type Speaker struct{}

func (s Speaker) Speak() any {

	for i := 0; i < 10; i++ {
		if i == 5 {
			return "hello world"
		}

		if i == 6 {
			return 1
		}

		if i == 7 {
			return true
		}

		if i == 8 {
			return "hello world!!!!"
		}

		if i == 9 {
			return "hello world!!!!!"
		}
	}
	return "hello world"
}

func main() {

	s:=Speaker{}
	fmt.Println(s.Speak())

	fmt.Println("Hello, World!")
	acceptanything(1)
	acceptanything("hello")
	acceptanything(true)
	acceptanything(1.2)

	time.Sleep(1200 * time.Millisecond)
}
