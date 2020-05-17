// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

package main

import "fmt"

func main() {
	var nums1, nums2 []int
	var v float64

	//
	nums1 = []int{1, 3}
	nums2 = []int{2}
	v = findMedianSortedArrays(nums1, nums2)
	if v != float64(2) {
		panic("panic")
	}

	//
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	v = findMedianSortedArrays(nums1, nums2)
	if v != float64(2.5) {
		panic("panic")
	}

	//
	nums1 = []int{2, 4}
	nums2 = []int{1, 3, 5, 6, 7, 8}
	v = findMedianSortedArrays(nums1, nums2)
	if v != float64(4.5) {
		panic("panic")
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		a, b        []int
		target      = [2]int{-1, -1}
		mergedArray = make([]int, 0, len(a)+len(b))
		mergeRestA  = true
		searchStart = 0
		median      float64
		ok          bool
	)

	// 先找我们的目标
	if (len(nums1)+len(nums2))%2 == 0 {
		target[0] = (len(nums1) + len(nums2)) / 2
		target[1] = target[0] + 1
	} else {
		target[0] = (len(nums1) + len(nums2) + 1) / 2
	}

	// 约定a比b长
	if len(nums1) >= len(nums2) {
		a, b = nums1, nums2
	} else {
		a, b = nums2, nums1
	}

	// 存在空数组
	if len(b) == 0 {
		median, _ = TryToFindMedian(a, target)
		return median
	}

	//
	for idx, bVal := range b {
		pos := SearchGreaterValueInA(a[searchStart:], bVal)
		fmt.Printf("search: pos=%d, bVal=%d, a=%v\n", pos, bVal, a[searchStart:])
		// 从某一段开始, bVal 总是比 a 大
		if pos < 0 {
			mergedArray = append(mergedArray, a[searchStart:]...)
			mergedArray = append(mergedArray, b[idx:]...)
			fmt.Printf("pos < 0  array: %v\n", mergedArray)
			if median, ok = TryToFindMedian(mergedArray, target); ok {
				return median
			}
			mergeRestA = false
			break
		}
		// bVal 总是比 a 小
		if pos == 0 {
			mergedArray = append(mergedArray, bVal)
			fmt.Printf("pos = 0  array: %v\n", mergedArray)
			if median, ok = TryToFindMedian(mergedArray, target); ok {
				return median
			}
			continue
		}
		// a中pos之前的都比bVal小
		mergedArray = append(mergedArray, a[searchStart:searchStart+pos]...)
		mergedArray = append(mergedArray, bVal)
		searchStart += pos
		fmt.Printf("pos(%d) > 0  array: %v\n", pos, mergedArray)
		if median, ok = TryToFindMedian(mergedArray, target); ok {
			return median
		}
	}
	if mergeRestA {
		mergedArray = append(mergedArray, a[searchStart:]...)
	}
	median, _ = TryToFindMedian(mergedArray, target)
	return median
}

func TryToFindMedian(nums []int, target [2]int) (float64, bool) {
	fmt.Printf("array: %v\n", nums)
	if len(nums) < target[0] || (target[1] > 0 && len(nums) < target[1]) {
		return 0, false
	}
	if target[1] < 0 {
		return float64(nums[target[0]-1]), true
	}
	return float64(nums[target[0]-1]+nums[target[1]-1]) / float64(2), true
}

// 如果pos返回-1, 则表示bVal比a数组中的任意一个数字都要大于等于
func SearchGreaterValueInA(a []int, bVal int) (pos int) {
	// 极端情况处理, 免除下面的边界情况处理
	if a[len(a)-1] <= bVal {
		return -1
	}
	if a[0] > bVal {
		return 0
	}
	// 一般情况, 找第一个比bVal大的数的索引
	var startItr = 0
	var endItr = len(a) - 1
	for {
		half := halfAndCeil(startItr + endItr)
		if a[half] > bVal {
			if a[half-1] <= bVal {
				return half
			} else {
				endItr = half - 1
			}
		}
		if a[half] <= bVal {
			if a[half+1] > bVal {
				return half + 1
			} else {
				startItr = half + 1
			}
		}
	}
	panic(fmt.Sprintf("no value found that is greater than %d in array a!", bVal))
}

func halfAndCeil(v int) int {
	if v%2 == 0 {
		return v / 2
	}
	return (v + 1) / 2
}
