package main

import (
	"log"
)

func main() {
	//TestMaxHuiWenBanJing()
	TestLongestPalindrome()
}

type MockString struct {
	s string
}

func NewMockString(s string) *MockString {
	return &MockString{
		s: s,
	}
}

func (ms *MockString) MockStringLen() int {
	return len(ms.s)*2 + 1
}

func (ms *MockString) LongestPalindrome() string {
	if len(ms.s) <= 1 {
		return ms.s
	}

	longestPalindromeLen := 0
	longestPalindromeIdx := 0
	for idx := 0; idx < ms.MockStringLen(); idx++ {
		if v := ms.MaxHuiWenBanJing(idx); v > longestPalindromeLen {
			longestPalindromeLen = v
			longestPalindromeIdx = idx
		}
	}
	return ms.GetPalindromeString(longestPalindromeIdx, longestPalindromeLen)
}

func (ms *MockString) GetPalindromeString(middleIdx int, maxHuiWenBanJing int) string {
	// 取所有奇数位的字符
	startIdx := middleIdx - maxHuiWenBanJing
	endIdx := middleIdx + maxHuiWenBanJing
	if !IsOdd(startIdx) {
		startIdx++
		endIdx--
	}
	return ms.s[(startIdx-1)/2 : (endIdx-1)/2+1]
}

func (ms *MockString) MaxHuiWenBanJing(charPos int) int {
	front := charPos - 1
	back := charPos + 1
	maxHuiWenBanJing := 0
	for front >= 0 && back < ms.MockStringLen() {
		if IsOdd(front) && ms.s[(front-1)/2] != ms.s[(back-1)/2] {
			break
		}
		maxHuiWenBanJing++
		front--
		back++
	}
	return maxHuiWenBanJing
}

func IsOdd(v int) bool {
	return v%2 != 0
}

func TestMaxHuiWenBanJing() {
	// #a#b#b#a#
	s := NewMockString("abba")
	if v := s.MaxHuiWenBanJing(0); v != 0 {
		panic("0")
	}
	if v := s.MaxHuiWenBanJing(1); v != 1 {
		panic("1")
	}
	if v := s.MaxHuiWenBanJing(2); v != 0 {
		panic("2")
	}
	if v := s.MaxHuiWenBanJing(3); v != 1 {
		log.Panicf("#3  %d != 1", v)
	}
	if v := s.MaxHuiWenBanJing(4); v != 4 {
		log.Panicf("#4  %d != 4", v)
	}
	log.Printf("OK\n")
}

func TestLongestPalindrome() {
	if v := NewMockString("abba").LongestPalindrome(); v != "abba" {
		log.Fatalf("#1  %d != abba", v)
	}
	if v := NewMockString("").LongestPalindrome(); v != "" {
		log.Fatalf("#2  %d != ", v)
	}
	if v := NewMockString("s").LongestPalindrome(); v != "s" {
		log.Fatalf("#3  %d != s", v)
	}
	if v := NewMockString("aaaa").LongestPalindrome(); v != "aaaa" {
		log.Fatalf("#4  %d != aaaa", v)
	}
	if v := NewMockString("aa").LongestPalindrome(); v != "aa" {
		log.Fatalf("#5  %d != aa", v)
	}
	if v := NewMockString("cbbd").LongestPalindrome(); v != "bb" {
		log.Fatalf("#6  %d != bb", v)
	}
	if v := NewMockString("babad").LongestPalindrome(); v != "bab" {
		log.Fatalf("#7  %d != bab", v)
	}

	log.Printf("OK\n")
}
