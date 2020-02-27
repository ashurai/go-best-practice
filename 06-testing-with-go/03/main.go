package main

import (
	"fmt"

	"github.com/ashurai/go-best-practice/06-testing-with-go/03/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3, 5, 6))
	fmt.Println(acdc.Sum(2, 3, 5, 6, 7, 8))
}
