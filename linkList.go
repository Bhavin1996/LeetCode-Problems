package leetcode

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}
type linkList struct {
	head *node
	len  int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
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
	newNode.value = value
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
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l *linkList) remove(value int) {
	var previous *node
	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
				return
			} else {
				previous.next = current.next
				return
			}
		}
		previous = current
	}
}

func mergeList(list1 *linkList, list2 *linkList) {
	ls := list1.head
	ls1 := list2.head
	var head *node
	if ls.value < ls1.value {
		head = ls
		ls = ls.next
	} else {
		head = ls1
		ls1 = ls1.next
	}

	temp := head
	for ls.next != nil && ls1.next != nil {
		if ls.value < ls1.value {
			temp.next = ls
			ls = ls.next
		} else {
			temp.next = ls1
			ls1 = ls1.next
		}
		temp = temp.next
	}
	if ls.next == nil {
		temp.next = ls1
	}
	if ls1.next == nil {
		temp.next = ls
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
