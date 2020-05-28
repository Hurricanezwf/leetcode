package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	if v := reverse(123); v != 321 {
		log.Fatalf("%d != 321", v)
	}
	if v := reverse(-123); v != -321 {
		log.Fatalf("%d != -321", v)
	}
	if v := reverse(120); v != 21 {
		log.Fatalf("%d != 21", v)
	}
	if v := reverse(7463847413); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	fmt.Printf("OK\n")
}

func reverse(x int) int {
	str := []byte(strconv.Itoa(x))
	left, right := 0, len(str)-1
	if str[left] == '-' {
		left++
	}
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	reverseVal, err := strconv.ParseInt(string(str), 10, 64)
	if err != nil {
		panic(err)
	}
	if reverseVal < int64(math.MinInt32) || reverseVal > int64(math.MaxInt32) {
		return 0
	}
	return int(reverseVal)
}
