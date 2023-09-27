package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	head = l1
	prev := head
	sum := 0
	for l1 != nil && l2 != nil {

		sum = l1.Val + l2.Val
		fmt.Println(sum)
		fmt.Println(prev.Val)
		l1.Val = sum % 10
		prev.Val += sum / 10
		fmt.Println(prev.Val)
		prev = l1
		l1 = l1.Next
		l2 = l2.Next
		fmt.Println(l1, l2)
	}
	return head
}
