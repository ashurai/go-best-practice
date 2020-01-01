package main

import "testing"

func TestSum(t *testing.T){
	x := mySum(2,3,4)
	if x != 9 {
		t.Error("Expected 5, but got", x)
	}
}