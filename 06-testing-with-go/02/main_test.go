package main

import "testing"

func TestSum(t *testing.T){
	type test struct{
		data []int
		answer int
	}
	testL := []test{
		test{[]int{14,15}, 29},
		test{[]int{10,13}, 23},
		test{[]int{9,15}, 24},
		test{[]int{14,13}, 27},
	}

	for _, v := range testL{
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected,",v.answer ,"but got", x)
		}
	}
}