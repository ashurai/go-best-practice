package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("max go routines", runtime.NumGoroutine())
	fmt.Println("CPUS", runtime.NumCPU())
	go func() { foo() }()
	go func() { bar() }()
	fmt.Println("Hello, playground")
	wg.Wait()
	fmt.Println("max go routines", runtime.NumGoroutine())
	fmt.Println("CPUS", runtime.NumCPU())
}

func foo() {

	fmt.Println("I am foo")
	wg.Done()
}

func bar() {

	fmt.Println("i am bar")
	wg.Done()
}
