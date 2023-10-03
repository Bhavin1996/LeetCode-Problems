package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}
	return head
}
