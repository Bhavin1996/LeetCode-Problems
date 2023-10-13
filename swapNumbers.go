// # Using XOR to swap two numbers

package leetcode

func swap(a int, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
