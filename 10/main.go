package main

import (
	"fmt"
	"time"

	"github.com/Hurricanezwf/pkg/assert"
)

func main() {
	funcs := []assert.AssertFunc{
		assert.True(isMatch("", "") == true, "1"),
		assert.True(isMatch("", ".") == false, "2"),
		assert.True(isMatch("", "*") == true, "3"),
		assert.True(isMatch("", "a") == false, "4"),
		assert.True(isMatch("ab", "") == false, "5"),
		assert.True(isMatch("a", "a*") == true, "6"),
		assert.True(isMatch("a", "a*.") == true, "7"),
		assert.True(isMatch("abc", "a.") == false, "8"),
		assert.True(isMatch("azc", "a.*") == true, "9"),
		assert.True(isMatch("azc", "az*.") == true, "10"),
		assert.True(isMatch("az", "az*.") == true, "11"),
		assert.True(isMatch("azzz", "az*") == true, "12"),
		assert.True(isMatch("aa", "a") == false, "13"),
		assert.True(isMatch("aa", "a*") == true, "14"),
		assert.True(isMatch("ab", ".*") == true, "15"),
		assert.True(isMatch("aab", "c*a*b") == true, "16"),
		assert.True(isMatch("mississippi", "mis*is*p*.") == false, "17"),
		assert.True(isMatch("a", ".*") == true, "18"),
		assert.True(isMatch("abcde", "a.*e") == true, "19"),
		assert.True(isMatch("ad", "a.*.d") == false, "20"),
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%d: OK\n", time.Now().Unix())
}

func isMatch(s, p string) bool {
	if s == "" {
		return isMatchWhenStringIsEmpty(p)
	}
	if p == "" {
		return isMatchWhenRegrexIsEmpty(s)
	}
	return isMatchWhenBothNotEmpty(s, p)
}

// 此时, 表达式要么为空, 要么就都是 *
func isMatchWhenStringIsEmpty(p string) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			return false
		}
	}
	return true
}

// 此时, 只有字符串为空时, 才能匹配
func isMatchWhenRegrexIsEmpty(s string) bool {
	return s == ""
}

//
func isMatchWhenBothNotEmpty(s, p string) bool {
	var (
		_s        = NewString(s)
		_p        = NewString(p)
		pLastChar = byte(0x00)
	)
	for !_p.EndOfString() {
		// 关键是识别p中的字符
		switch _p.CurByte() {
		case '.':
			pLastChar = _p.CurByte()
			_s.Next()
			_p.Next()
			if 0x00 == _s.CurByte() && 0x00 == _p.CurByte() {
				// 两个都到头了
				return true
			}
			if _s.EndOfString() {
				// string到头了, regrex没到头
				if _p.CurByte() != '*' {
					return false
				}
				_p.Next()
				continue
			}
		case '*':
			// 遇到 *, 要求跳过N个上一个字符
			switch pLastChar {
			case '.':
				pLastChar = '.'
				_s.Next()
			default:
				// pLastChar 不动
				_s.SkipByte(pLastChar)
			}
			_s.Next()
			_p.Next()
			if 0x00 == _s.CurByte() && 0x00 == _p.CurByte() {
				// 两个都到头了
				return true
			}
			if _s.EndOfString() {
				// string 到头了, regrex没到头
				if _p.CurByte() != '*' && _p.CurByte() != '.' {
					return false
				}
				continue
			}
		default:
			pLastChar = _p.CurByte()
			if _p.ReadNextByte() == '*' {
				// 表达式是 c* 的情况
				_s.SkipByte(_p.CurByte())
				_p.Next().Next()
				continue
			}
			if _p.CurByte() != _s.CurByte() {
				// 表达式字符与字符串字符不一致, 则后移字符串
				_s.Next()
				if _s.EndOfString() {
					return false
				}
				continue
			}
			// 字符相等
			_p.Next()
			_s.Next()
		}
	}
	return _s.EndOfString()
}

//
type String struct {
	s   string
	idx int
}

func NewString(s string) *String {
	if s == "" {
		panic("string can't be empty")
	}
	return &String{
		s: s,
	}
}

func (s *String) CurIdx() int {
	return s.idx
}

func (s *String) CurByte() byte {
	if s.EndOfString() {
		return 0x00
	}
	return s.s[s.idx]
}

func (s *String) Next() *String {
	s.idx++
	return s
}

func (s *String) ReadNextByte() byte {
	idx := s.idx + 1
	if idx >= len(s.s) {
		return 0x00
	}
	return s.s[idx]
}

func (s *String) SkipByte(c byte) {
	for i := s.idx; i < len(s.s); i++ {
		if s.s[i] == c {
			s.idx++
			continue
		}
		break
	}
}

func (s *String) EndOfString() bool {
	return s.idx >= len(s.s)
}
