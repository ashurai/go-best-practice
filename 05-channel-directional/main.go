package main

import "fmt"

func main() {
	ch := make(<-chan int)

	c := make(chan<- int)

	fmt.Println("receive only")
	fmt.Printf("%T\t\n", ch)
	fmt.Println("send only")
	//fmt.Println(<-ch)
	fmt.Printf("%T\t", c)
}
