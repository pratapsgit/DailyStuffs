package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func firstTraverse(l int) {
	for i := 0; i < l; i++ {
		fmt.Printf("firstTraverse : %d\n", i)
	}
	wg.Done()
}

func secondTraverse(l int) {
	for i := 0; i < l; i++ {
		fmt.Printf("secondTraverse : %d\n", i)
	}
}

func main() {
	fmt.Println("Number of CPU's ", runtime.NumCPU())
	fmt.Println("Number of Go Routine's ", runtime.NumGoroutine())
	wg.Add(1)
	go firstTraverse(15)
	secondTraverse(15)
	fmt.Println("Number of CPU's ", runtime.NumCPU())
	fmt.Println("Number of Go Routine's ", runtime.NumGoroutine())
	wg.Wait()
}
