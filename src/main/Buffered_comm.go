package main

import (
	"fmt"
	"sync"
)

func sender(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("hello from channel %v", i)
	}
	close(ch)
}

func reciver(ch chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	msg := <-ch
	msg1 := <-ch
	fmt.Println("the message is this", msg, msg1)
}

func reciver2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Println("the message from reciver 2 is", msg)
}

func reciver3(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	msg1 := <-ch
	msg2 := <-ch
	msg3 := <-ch
	msg4 := <-ch
	fmt.Println("the message from reciver 3 is\n", msg, "\n", msg1, "\n", msg2, "\n", msg3, "\n", msg4)
}

func main() {
	fmt.Print("working of the buffred channel here like how it is working here\n")
	ch := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go sender(ch, &wg)
	wg.Add(1)
	go reciver(ch, &wg)
	wg.Add(1)
	go reciver2(ch, &wg)
	wg.Add(1)
	go reciver3(ch, &wg)
	wg.Wait()
}
