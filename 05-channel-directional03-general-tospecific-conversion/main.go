package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(chan<- int)
	cs := make(<-chan int)

	fmt.Printf("C\t %T\n", c)
	fmt.Printf("CR\t %T\n", cr)
	fmt.Printf("CD\t %T\n", cs)

	// General to specific type conversion
	cr = c
	cs = c

	fmt.Printf("C\t%T\n", c)
	fmt.Printf("C\t%T\n", (<-chan int)(c))
	fmt.Printf("C\t%T\n", (chan<- int)(c))
}
