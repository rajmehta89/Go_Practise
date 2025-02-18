package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	// Get current goroutine ID and thread ID
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	id := string(buf[:n])
	
	for i := 0; i < 5; i++ {
		fmt.Printf("Time: %v - Goroutine ID: %v - Thread: %d - Message: %s\n", 
			time.Now().Format("15:04:05.000"), 
			id[10:20], // Extract just the goroutine ID number
			getThreadID(),
			s)
		time.Sleep(100 * time.Millisecond)
	}
}

func getThreadID() int {
	return runtime.NumGoroutine()
}

// Example 1: Basic Channel Communication
func sender(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Important: Close channel when done
}

func receiver(ch chan string, name string) {
	for msg := range ch { // Range over channel until closed
		fmt.Printf("Receiver %s got: %s\n", name, msg)
	}
}

// Example 2: Buffered Channel
func bufferExample(ch chan int) {
	fmt.Println("Sending to buffered channel...")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

// Example 3: Worker Pool Pattern
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(200 * time.Millisecond) // Simulate work
		results <- job * 2
	}
}

// Example 4: Select Pattern
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

func main() {
	// Run all examples
	fmt.Println("=== Basic Struct Example ===")
	BasicStructExample()

	fmt.Println("\n=== Nested Struct Example ===")
	NestedStructExample()

	fmt.Println("\n=== Struct Methods Example ===")
	StructMethodsExample()

	fmt.Println("\n=== Struct Tags Example ===")
	StructTagsExample()
}
