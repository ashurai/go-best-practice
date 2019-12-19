package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// sending
	go send(eve, odd, quit)

	//receiving
	receive(eve, odd, quit)

	fmt.Println("i am completted, so exiting job!")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel", v)
		case v := <-o:
			fmt.Println("from the odd channel", v)
		case v := <-q:
			fmt.Println("from the quit channel", v)
			return
		}
	}
}
