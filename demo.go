package leetcode

import "fmt"

type TreeNode struct {
	value int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTopView(root *TreeNode) {
	if root == nil {
		return
	}

	nodeMap := make(map[int]*TreeNode)
	var topView func(node *TreeNode, distance int)
	topView = func(node *TreeNode, distance int) {
		if node == nil {
			return
		}

		if _, found := nodeMap[distance]; !found {
			nodeMap[distance] = node
			fmt.Println(node.value)
		}

		topView(node.Left, distance-1)
		topView(node.Right, distance+1)
	}

	topView(root, 0)

}

func main() {
	// Example usage:
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 4,
			},
			Right: &TreeNode{
				Value: 5,
			},
		},
		Right: &TreeNode{
			Value: 3,
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}

	fmt.Println("Top View:")
	printTopView(root)
	fmt.Println("\nBottom View:")
	printBottomView(root)
}
