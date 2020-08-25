package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	count := 0
	threads := 100

	wg.Add(threads)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			runtime.Gosched()
			count = v
			mu.Unlock()
			fmt.Println("Count is ", count)
			wg.Done()
		}()
	}

	wg.Wait()
}
