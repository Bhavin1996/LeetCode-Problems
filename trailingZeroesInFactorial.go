// https://www.youtube.com/watch?v=69jsFIMINpI

package leetcode

func trailingZeroes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	res := 0
	for i := 5; i <= n; i *= 5 {
		res = res + n/i
	}
	return res
}
