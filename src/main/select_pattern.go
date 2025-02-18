package main

import (
	"fmt"
	"time"
)

func selectExample(ch1, ch2 chan string, done chan bool) {
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
			done <- true
			return
		}
	}
}

func SelectPatternExample() {
	fmt.Println("\n=== Select Pattern Example ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			ch1 <- fmt.Sprintf("Message from ch1: %d", i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Millisecond)
			ch2 <- fmt.Sprintf("Message from ch2: %d", i)
		}
	}()

	go selectExample(ch1, ch2, done)
	<-done
} 