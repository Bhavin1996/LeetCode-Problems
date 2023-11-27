// package leetcode

// import "fmt"

// type TreeNode struct {
// 	value int
// 	left  *TreeNode
// 	right *TreeNode
// }

// //New Tree creation
// func NewTree(root int) *TreeNode {
// 	return &TreeNode{value: root}
// }

// //Insert method
// func (T *TreeNode) InsertNode(k int) {
// 	//Move right
// 	if T.value < k {
// 		if T.right == nil {
// 			T.right = &TreeNode{value: k}
// 		} else {
// 			T.right.InsertNode(k)
// 		}
// 	} else if T.value > k {
// 		//Move left
// 		if T.left == nil {
// 			T.left = &TreeNode{value: k}
// 		} else {
// 			T.left.InsertNode(k)
// 		}
// 	}
// }

// //search method

// func (T *TreeNode) SearchNode(k int) bool {

// 	if T == nil {
// 		//when searched node is not found
// 		return false
// 	}

// 	//move right
// 	if T.value < k {
// 		T.right.SearchNode(k)
// 	} else if T.value > k {
// 		//move left
// 		T.left.SearchNode(k)
// 	}

// 	return true
// }

// func mainTree() {
// 	fmt.Println("Enter the root value ...")
// 	var input int
// 	fmt.Scan(&input)
// 	res := NewTree(input)
// 	fmt.Println("Starting the tree operations")
// 	res.InsertNode(5)
// 	fmt.Println(res)
// }

package leetcode

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func printTopView(root *TreeNode) {
	if root == nil {
		return
	}

	// Create a map to store the nodes at each horizontal distance
	nodeMap := make(map[int]*TreeNode)

	// Helper function for top view traversal
	var topViewHelper func(node *TreeNode, distance int)
	topViewHelper = func(node *TreeNode, distance int) {
		if node == nil {
			return
		}

		// Only store the leftmost node for each horizontal distance
		if _, found := nodeMap[distance]; !found {
			nodeMap[distance] = node
			fmt.Print(node.Value, " ")
		}

		// Traverse left and right with adjusted distances
		topViewHelper(node.Left, distance-1)
		topViewHelper(node.Right, distance+1)
	}

	// Start the top view traversal
	topViewHelper(root, 0)
}

func printBottomView(root *TreeNode) {
	if root == nil {
		return
	}

	// Create a map to store the nodes at each horizontal distance
	nodeMap := make(map[int]*TreeNode)

	// Helper function for bottom view traversal
	var bottomViewHelper func(node *TreeNode, distance int)
	bottomViewHelper = func(node *TreeNode, distance int) {
		if node == nil {
			return
		}

		// Always update the node for each horizontal distance
		nodeMap[distance] = node

		// Traverse left and right with adjusted distances
		bottomViewHelper(node.Left, distance-1)
		bottomViewHelper(node.Right, distance+1)
	}

	// Start the bottom view traversal
	bottomViewHelper(root, 0)

	// Print the nodes in sorted order of horizontal distance
	for _, node := range nodeMap {
		fmt.Print(node.Value, " ")
	}
}

func main_tree() {
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
