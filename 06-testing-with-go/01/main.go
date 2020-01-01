package main

import "fmt"

func main(){
	fmt.Println("Adding 2+3 = ", mySum(2,3))
	fmt.Println("Adding 7+8 = ", mySum(7,8))
	fmt.Println("Adding 7+8+9+10+12 = ", mySum(7,8,9,10,12))
}

func mySum(si ...int) int {
	sum := 0
	for _, v := range si {
		sum += v
	}

	return sum
}