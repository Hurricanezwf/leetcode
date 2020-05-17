package main

import (
	"fmt"
	"testing"
)

func Test_SearchGreaterValueInA(t *testing.T) {
	var a []int

	a = []int{2, 2, 7, 9, 11}
	fmt.Printf("4=%d\n", SearchGreaterValueInA(a, 10))

	a = []int{2, 2}
	fmt.Printf("-1=%d\n", SearchGreaterValueInA(a, 10))
}
