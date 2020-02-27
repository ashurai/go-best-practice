package acdc

import "fmt"

// Sum is just an exmple to show you writting documentation
func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
}

// EaxmpleSum is to help you how to write
// your exampel for documentation
func ExampleSum() {
	fmt.Println(Sum(2, 3, 4, 5))
	fmt.Println(Sum(2, 3))
}
