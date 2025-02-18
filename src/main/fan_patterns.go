package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func modifyMap(m map[int]int) {
	mu.Lock()
	m[1] = 100 // Concurrent modification (unsafe)
	mu.Unlock()
}

func main() {

	m := make(map[int]int)
	go modifyMap(m)
	go modifyMap(m)

	time.Sleep(time.Second)
	fmt.Println(m) // May cause a runtime error
}
