package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	mutex   sync.Mutex
	counter int
)

func main() {
	numRoutines := 10

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		// if i is used as id, through closure, then you may get some undesired behavior
		// comment and uncomment like lines to witness
		go func(id int) {
			//go func() {
			mutex.Lock() // using the {} to id where the actions when lock is in place
			{
				value := counter

				// yields goroutine, but the lock stops other goroutines from changing the counter value
				// the mutex locks the goroutine as the context of the scheduler
				runtime.Gosched()

				value++
				fmt.Printf("id: %d\tcounter: %d\tvalue: %d\n", id, counter, value)
				//fmt.Printf("id: %d\tcounter: %d\tvalue: %d\n", i, counter, value)
				counter = value
			}
			mutex.Unlock()
			wg.Done()
			//}()
		}(i)
	}
	wg.Wait()
}
