package main

import (
	"fmt"
	"log"
)

func main() {
	//v := isHuiWen("abxba", 0, 4)
	//fmt.Printf("isHuiWen=%t\n", v)

	//m := indexCharDict("abxba")
	//fmt.Printf("%v\n", m)

	if str := longestPalindrome("abxba"); str != "abxba" {
		log.Panicf("%s != abxba", str)
	}
	if str := longestPalindrome("babad"); str != "bab" && str != ",ba" {
		log.Panicf("%s != bab", str)
	}
	if str := longestPalindrome("cbbd"); str != "bb" {
		log.Panicf("%s != bb", str)
	}
	if str := longestPalindrome("a"); str != "a" {
		log.Panicf("%s != a", str)
	}
	if str := longestPalindrome(""); str != "" {
		log.Panicf("%s != ", str)
	}
	if str := longestPalindrome("ac"); str != "a" {
		log.Panicf("%s != a", str)
	}
	if str := longestPalindrome("ccc"); str != "ccc" {
		log.Panicf("%s != ccc", str)
	}

	log.Printf("OK\n")
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	fmt.Printf("s=%s\n", s)
	var rt string
	var dict = indexCharDict(s)
	for char, charPosList := range dict {
		for i := 0; i < len(charPosList); i++ {
			for j := i + 1; j < len(charPosList); j++ {
				if isHuiWen(s, charPosList[i], charPosList[j]) && charPosList[j]+1-charPosList[i] > len(rt) {
					rt = s[charPosList[i] : charPosList[j]+1]
					fmt.Printf("%c [%d:%d] -- %v -- %s\n", char, charPosList[i], charPosList[j], true, rt)
					continue
				}
				fmt.Printf("%c [%d:%d] -- %v -- %s\n", char, charPosList[i], charPosList[j], false, rt)
			}
		}
	}
	// 无重复字符, 取首字符
	if rt == "" {
		rt = s[0:1]
	}
	return rt
}

func indexCharDict(s string) map[uint8][]int {
	var m = make(map[uint8][]int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		m[c] = append(m[c], i)
	}
	return m
}

// isHuiWen判断s的[startIdx,endIdx]之间的字串是否是回文
func isHuiWen(s string, startIdx, endIdx int) bool {
	var startItr = startIdx
	var endItr = endIdx
	for startItr < endItr {
		if s[startItr] != s[endItr] {
			return false
		}
		startItr++
		endItr--
	}
	return true
}
