/* # Floyd's Cycle-Finding Algorithm,
also known as the "tortoise and hare" algorithm, is used to detect cycles or loops in a sequence,
such as a linked list or an array.
The algorithm is named after Robert W. Floyd, who is credited with its discovery.

Here's a high-level description of the algorithm:

1. Initialize two pointers: one slow pointer and one fast pointer. Initially,
both pointers point to the first element of the sequence.

2. Move the slow pointer one step at a time (i.e., advance it by one element),
and move the fast pointer two steps at a time (i.e., advance it by two elements). This simulates a race between the two pointers.

3. Continue moving both pointers until one of the following conditions is met:
   - The fast pointer reaches the end of the sequence (i.e., it becomes `nil` or exceeds the bounds of the array or linked list).
     In this case, there is no cycle in the sequence.
   - The fast pointer meets (or overtakes) the slow pointer. This indicates the presence of a cycle in the sequence.

4. If the fast pointer met the slow pointer, there is a cycle in the sequence.
You can then determine the starting point of the cycle and its length if needed.

The key insight behind this algorithm is that if there is a cycle in the sequence,
the fast pointer will eventually catch up to the slow pointer as they move around the cycle.
This property allows the algorithm to detect cycles efficiently with a time complexity of O(n), where n is the length of the sequence.

Floyd's Cycle-Finding Algorithm is widely used in computer science and has applications in various areas,
including detecting cycles in linked lists, graph algorithms, and performance optimization in algorithms and data structures.*/

package leetcode

import (
	"fmt"
)

// Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

// Function to detect a cycle in a linked list.
func hasCycle(head *listNode) bool {
	// Initialize two pointers, slow and fast.
	slow := head
	fast := head

	// Traverse the linked list using two pointers.
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer one step at a time.
		fast = fast.Next.Next // Move fast pointer two steps at a time.

		// If there is a cycle, the slow and fast pointers will meet.
		if slow == fast {
			return true
		}
	}

	// If there is no cycle, the fast pointer will reach the end of the list.
	return false
}

func main1() {
	// Create a linked list with a cycle (for testing).
	head := &listNode{Val: 3}
	node1 := &listNode{Val: 2}
	node2 := &listNode{Val: 0}
	node3 := &listNode{Val: -4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1 // Cycle: -4 -> 2

	// Check if the linked list has a cycle.
	result := hasCycle(head)
	fmt.Println("Has Cycle:", result) // Should print true
}
