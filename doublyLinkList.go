package leetcode

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type doublyLinkList struct {
	Head *Node
	Tail *Node
}

func (dd *doublyLinkList) addNode(val int) {
	newNode := new(Node)
	newNode.prev = dd.Tail
	newNode.value = val
	if dd.Tail != nil {
		dd.Tail.next = newNode
	}
}
