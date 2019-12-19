package main

import "fmt"

func main() {
	ch := make(chan<- int)

	ch <- 42

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch) // dosen't work
	//fmt.Println(<-ch)
}
