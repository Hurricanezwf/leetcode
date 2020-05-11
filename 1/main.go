package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", twoSum([]int{23, 234, 63, 23}, 86))
}

/*
 * 不能排序
 *
 * 思路:
 * 构造一个链式的map
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int][]int) // value ==> idx array
	for idx, v := range nums {
		idxArray := m[v]
		if idxArray == nil {
			idxArray = make([]int, 0, 0)
		}
		m[v] = append(idxArray, idx)
	}

	for idx1, v1 := range nums {
		for _, idx2 := range m[target-v1] {
			if idx2 != idx1 {
				return []int{idx1, idx2}
			}
		}
	}
	return []int{}
}

// 这是执行最快的解
func twoSumTheFastest(nums []int, target int) []int {
	m := make(map[int]int)
	for idx1, v1 := range nums {
		if idx2, ok := m[target-v1]; ok {
			return []int{idx2, idx1}
		}
		m[v1] = idx1
	}
	return []int{}
}
