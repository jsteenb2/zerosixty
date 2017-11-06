package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	atomic functions are the most performant way to serialize operations on shared memory
*/

var counter int64

func incCounter(step int) {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&counter, int64(step))
		if step > 1 {
			fmt.Println("--->", counter)
		} else {
			fmt.Println(counter)
		}

	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		incCounter(1)
		wg.Done()
	}()
	go func() {
		incCounter(10)
		wg.Done()
	}()

	wg.Wait()
}
