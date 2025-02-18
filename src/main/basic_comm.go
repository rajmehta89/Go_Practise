package main

import (
	"fmt"
	"time"
)

func senderr(ch chan string) {

	fmt.Printf("in the sender side here")

	for i := 1; i < 5; i++ {
		ch <- fmt.Sprint("message i: %d", i)
		time.Sleep(100 * time.Microsecond)
	}
	close(ch)

}

func reciver(ch chan string, name string) {

	fmt.Println("in the reciver here")
	for msg := range ch {
		fmt.Printf("Receiver %s got %s\n", name, msg)
	}
	fmt.Println("reciver ended here")
}

func main() {
	ch := make(chan string)
	go reciver(ch, "A")
	go reciver(ch, "B")
	go senderr(ch)

	time.Sleep(1 * time.Second)
}
