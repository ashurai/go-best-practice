package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Go Routines \t", runtime.NumGoroutine())
	fmt.Println("i am completted, so exiting job!")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	//wg.Add(2)
	for v := range c1 {
		wg.Add(1)
		fmt.Println("before Go Routines \t", runtime.NumGoroutine())
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
		fmt.Println("after Go Routines \t", runtime.NumGoroutine())
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
