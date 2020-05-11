// https://leetcode-cn.com/problems/add-two-numbers/

package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
}

func (l *ListNode) Append(simpleV int) (next *ListNode) {
	l.Next = NewListNode(simpleV)
	return l.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var (
		jinwei, yushu int
		head          = NewListNode(-1)
		tail          = head
		l1Itr         = l1
		l2Itr         = l2
		finalItr      *ListNode
	)

	for l1Itr != nil && l2Itr != nil {
		jinwei, yushu = splitValue(l1Itr.Val + l2Itr.Val + jinwei)
		tail = tail.Append(yushu)
		l1Itr, l2Itr = l1Itr.Next, l2Itr.Next
	}
	// 链表长度相等
	if l1Itr == nil && l2Itr == nil {
		if jinwei > 0 {
			tail = tail.Append(jinwei)
		}
		return head.Next
	}

	if l1Itr != nil {
		finalItr = l1Itr
	} else {
		finalItr = l2Itr
	}
	for finalItr != nil {
		jinwei, yushu = splitValue(finalItr.Val + jinwei)
		tail = tail.Append(yushu)
		finalItr = finalItr.Next
	}
	if jinwei > 0 {
		tail = tail.Append(jinwei)
	}
	return head.Next
}

func splitValue(v int) (jinwei, yushu int) {
	if v <= 9 {
		return 0, v
	}
	return 1, v % 10
}
