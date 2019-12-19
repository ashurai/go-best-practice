package main

import "fmt"

func main() {
	c := make(chan int)
	// sending
	go bar(c)

	// receiving
	foo(c) // it will receive the value and job will be completed

	// if below code will be availabel then need to add witgroup,
	// else will reach till foo and block
	// go foo(c)

	fmt.Println("i am completted, so exiting job!")
}

// receiving
func foo(c <-chan int) {
	fmt.Println("value: ", <-c)
}

// sending
func bar(c chan<- int) {
	c <- 42
}
