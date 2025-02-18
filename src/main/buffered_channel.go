package main

import (
	"fmt"
	"time"
)

func bufferExample(ch chan int) {
	fmt.Println("Sending to buffered channel...")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

func BufferedChannelExample() {
	fmt.Println("\n=== Buffered Channel Example ===")
	buffCh := make(chan int, 3) // Buffer size of 3
	go bufferExample(buffCh)
	time.Sleep(500 * time.Millisecond)
	for num := range buffCh {
		fmt.Printf("Received: %d\n", num)
	}
} 