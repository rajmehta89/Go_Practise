package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Starting program...")

	for i := 0; i < 10; i++ {
		go func() {
			data := make([]byte, 10*1024*1024) // Allocate 10MB
			_ = data                           // Use data to prevent optimization
		}()
	}

	time.Sleep(2 * time.Second) // Allow time for GC to run

	fmt.Println("Manually triggering GC...")
	runtime.GC() // Force garbage collection

	fmt.Println("Memory stats:")
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Heap Alloc: %d KB\n", memStats.HeapAlloc/1024)
	fmt.Printf("GC Count: %d\n", memStats.NumGC)
}

//func getThreadID() int {
//	return runtime.NumGoroutine()
//}
//
//func sy() {
//	fmt.Println(":pritn the stack trace here")
//
//	fmt.Println("to get teh runtime procks here in the go lang for that", runtime.GOMAXPROCS(10))
//
//	//buff := make([]byte, 1024)
//	//n := string(runtime.Stack(buff, false))
//
//	//fmt.Println(string(n[10:]))
//	fmt.Println("thread id in this goroutine is", getThreadID())
//	fmt.Print("hello raj")
//}
//
//func system() {
//	fmt.Println("go routine is started here")
//	//time.Sleep(2 * time.Second)
//	fmt.Println("go routine is over here")
//}
//
//func main() {
//	fmt.Println("number of cpu", runtime.NumCPU())
//	fmt.Println("number of goroutines before", runtime.NumGoroutine())
//	go sy()
//	go system()
//	fmt.Println("after the say() method go routines are", runtime.NumGoroutine())
//	time.Sleep(2 * time.Second)
//
//}
