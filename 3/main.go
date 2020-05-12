// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func main() {
	cases := map[string]int{
		"abba":     2,
		"hello":    3,
		"helo":     4,
		"abcabcbb": 3,
		"":         0,
		"bbbb":     1,
		"pwwkew":   3,
	}

	for str, v := range cases {
		if vv := lengthOfLongestSubstring(str); vv != v {
			panic(fmt.Sprintf("str=%s  value(%d)!=%d\n", str, vv, v))
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	var (
		head, tail int
		charIndex  = make(map[uint8]int)
		maxLen     int
	)

	for ; tail < len(s); tail++ {
		pos, ok := charIndex[s[tail]]
		if ok && pos >= 0 {
			if tail-head > maxLen {
				maxLen = tail - head
			}
			// 清理掉字串的位置记录
			for i := head; i < pos+1; i++ {
				charIndex[s[i]] = -1
			}
			head = pos + 1
		}
		charIndex[s[tail]] = tail
	}
	if tail-head > maxLen {
		maxLen = tail - head
	}
	return maxLen
}
