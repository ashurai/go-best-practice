package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 42

	ch <- 43

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
