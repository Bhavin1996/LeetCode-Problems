package leetcode

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkList struct {
	head *node
	len  int
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

func (l linkList) String() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s", iterator))
	}
	return sb.String()
}
