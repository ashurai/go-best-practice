package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(chan<- int)
	cs := make(<-chan int)

	fmt.Printf("C\t %T\n", c)
	fmt.Printf("CR\t %T\n", cr)
	fmt.Printf("CD\t %T\n", cs)

	// General to specific type assigment of channels
	cr = c
	cs = c

	fmt.Printf("C\t%T\n", c)
	fmt.Printf("CR\t%T\n", cr)
	fmt.Printf("CD\t%T\n", cs)
}
