package main

func main() {
	if v := isPalindrome(121); !v {
		panic(121)
	}
	if v := isPalindrome(-121); v {
		panic(-121)
	}
	if v := isPalindrome(10); v {
		panic(10)
	}
	if v := isPalindrome(0); !v {
		panic(0)
	}
	if v := isPalindrome(-1); v {
		panic(-1)
	}
}

// 解法1: 转成字符串处理
//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	if x >= 0 && x <= 9 {
//		return true
//	}
//
//	s := strconv.Itoa(x)
//	left, right := 0, len(s)-1
//	for left <= right {
//		if s[left] != s[right] {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

// 解法2: 整数逆序后对比大小
func isPalindrome(x int) bool {
	if x > 9 {
		var xcopy = x
		var xreverse = 0
		for xcopy > 0 {
			xreverse = 10*xreverse + xcopy%10
			xcopy /= 10
		}
		return x == xreverse
	}
	return x >= 0
}
