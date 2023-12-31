package leetcode

import (
	"fmt"
	"strings"
)

type node struct {
	val  int
	next *node
}

type linkList struct {
	head *node
	len  int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (l linkList) String() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s ", iterator))
	}
	return sb.String()
}

func (l *linkList) add(value int) {
	newNode := new(node)
	newNode.val = value
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l linkList) get(value int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.val == value {
			return iterator
		}
	}
	return nil
}

func (l *linkList) remove(value int) {
	var prev *node
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.val == value {
			if l.head == iterator {
				l.head = iterator.next
				return
			} else {
				prev.next = iterator.next
				iterator.next = nil
				return
			}
		}
		prev = iterator
	}
}

func (l *linkList) reverse() {
	var next, prev *node
	curr := l.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
func main_LL() {
	ll := linkList{}
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	fmt.Println(ll)
	fmt.Println(ll.get(2))
	ll.remove(2)
	fmt.Println(ll)
}
