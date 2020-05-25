// https://leetcode-cn.com/problems/zigzag-conversion/submissions/

package main

import (
	"log"
)

func main() {
	v := NewZConvert(4, "LEETCODEISHIRING").Convert()
	if v != "LDREOEIIECIHNTSG" {
		log.Fatalf("%s != LDREOEIIECIHNTSG", v)
	}

	v = NewZConvert(3, "LEETCODEISHIRING").Convert()
	if v != "LCIRETOESIIGEDHN" {
		log.Fatalf("%s != LCIRETOESIIGEDHN", v)
	}
}

type ZConvert struct {
	s       string
	numRows int
}

func NewZConvert(numRows int, s string) *ZConvert {
	return &ZConvert{
		numRows: numRows,
		s:       s,
	}
}

func (c *ZConvert) Convert() string {
	if c.numRows == 1 {
		return c.s
	}

	buf := make([]byte, 0, len(c.s))
	anchors := c.GetBottomAnchor()
	for i := (c.numRows - 1); i >= 0; i-- {
		for _, anchor := range anchors {
			left := anchor - 1 - i
			right := anchor - 1 + i
			//fmt.Printf("anchor=%d, left=%d, right=%d\n", anchor, left, right)
			if left == right {
				right = -1 // 这个就是锚点, 只需要写一个就行了
			}
			if left >= 0 && left < len(c.s) {
				buf = append(buf, c.s[left])
				//fmt.Printf("left[%d]: %c\n", left, c.s[left])
			}
			if right-left != (2*c.numRows-2) && right >= 0 && right < len(c.s) {
				buf = append(buf, c.s[right])
				//fmt.Printf("righ[%d]: %c\n", right, c.s[right])
			}
		}
	}
	return string(buf)
}

func (c *ZConvert) GetBottomAnchor() []int {
	bottomAnchar := make([]int, 0, len(c.s))
	for i := 1; i <= len(c.s); i++ {
		anchor := c.numRows*i + (i-1)*(c.numRows-2)
		bottomAnchar = append(bottomAnchar, anchor)
		if anchor >= len(c.s) {
			break
		}
	}
	return bottomAnchar
}
