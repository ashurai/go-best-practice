package main

import "fmt"

func main() {
	c := make(chan int)
	// sending
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receiving
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("i am completted, so exiting job!")
}
