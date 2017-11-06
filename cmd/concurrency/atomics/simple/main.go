package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	atomic functions are the most performant way to serialize operations on shared memory
*/

const (
	red   = "\033[0;31m"
	nc    = "\033[0m"
	green = "\033[0;32m"
)

var counter int64

func incCounter(step int) {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&counter, int64(step))
		if step == 1 {
			fmt.Println(green, counter, nc)
		} else {
			fmt.Println(red, counter, nc)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// keyword `go` is how you make something run concurrently
	go func() {
		incCounter(1)
		wg.Done()
	}()
	go func() {
		incCounter(-1)
		wg.Done()
	}()

	wg.Wait()
}
