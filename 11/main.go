// https://leetcode-cn.com/problems/container-with-most-water/

package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	v := maxArea(arr)
	fmt.Printf("max array is %d\n", v)
}

// 双指针法
func maxArea(height []int) int {
	startItr, endItr := 0, len(height)-1
	maxVolume, tmpVolume := 0, 0
	for endItr >= startItr {
		if height[startItr] < height[endItr] {
			tmpVolume = height[startItr] * (endItr - startItr)
			startItr++
		} else {
			tmpVolume = height[endItr] * (endItr - startItr)
			endItr--
		}
		if tmpVolume > maxVolume {
			maxVolume = tmpVolume
		}
	}
	return maxVolume
}

// 暴力求解
//func maxArea(height []int) int {
//	tmpVolume := 0
//	maxVolume := 0
//	for maxStep := len(height) - 1; maxStep >= 1; maxStep-- {
//		startItr, endItr := 0, maxStep
//		for endItr < len(height) {
//			if height[startItr] < height[endItr] {
//				tmpVolume = height[startItr] * maxStep
//			} else {
//				tmpVolume = height[endItr] * maxStep
//			}
//			if tmpVolume > maxVolume {
//				maxVolume = tmpVolume
//			}
//			startItr++
//			endItr++
//		}
//	}
//	return maxVolume
//}
