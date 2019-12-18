package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))

			runtime.Gosched()

			wg.Done()
		}()
		fmt.Println("No of GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("No of GoRoutines", runtime.NumGoroutine())
	fmt.Println("Counter", counter)

}
