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
	red    = "\033[0;31m"
	nc     = "\033[0m"
	green  = "\033[0;32m"
	cyan   = "\033[0;36m"
	blue   = "\033[0;34m"
	purple = "\033[0;35m"
)

var counter int64

func incCounter(step int) {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&counter, int64(step))
		if step > 0 {
			fmt.Println(green, counter, nc)
		} else if step == 0 {
			fmt.Println(blue, counter, nc)
		} else {
			fmt.Println(red, counter, nc)
		}

	}
}

func main() {
	numRoutines := 7
	var wg sync.WaitGroup
	wg.Add(numRoutines)

	// wg.Done() not called enough to stop wg.Wait() from completing
	for i := -3; i < numRoutines-4; i++ {
		go func(step int) {
			incCounter(step)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
