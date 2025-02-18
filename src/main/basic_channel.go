package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func receiver(ch chan string, name string) {
	for msg := range ch {
		fmt.Printf("Receiver %s got: %s\n", name, msg)
	}
}

func BasicChannelExample() {
	fmt.Println("\n=== Basic Channel Communication Example ===")
	ch := make(chan string)
	go sender(ch)
	go receiver(ch, "A")
	go receiver(ch, "B")
	time.Sleep(1 * time.Second)
} 