package main

import (
	"fmt"
	"time"
)

// Channel communication
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch)
	go consumer(ch, done)

	<-done
} 