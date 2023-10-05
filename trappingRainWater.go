// # Two pointer approach

package leetcode

func trap(height []int) int {
	l := len(height)
	if l == 0 || l == 1 || l == 2 {
		return 0
	}
	trappedRain := 0
	left, right := 0, l-1
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				trappedRain += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				trappedRain += (rightMax - height[right])
			}
			right--
		}
	}
	return trappedRain
}
