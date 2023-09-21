/* This problem is solved by the Boyer Moore's Majority vote algorithm
reference link : https://youtu.be/X0G5jEcvroo?si=kLfICzd8tAZr0TlX
*/

package leetcode

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	if len(nums) == 1 || len(nums) == 2 {
		return nums[0]
	}
	for _, val := range nums {
		if count == 0 {
			candidate = val
			count = 1
		} else if val == candidate {
			count += 1
		} else if val != candidate {
			count -= 1
		}
	}
	count = 0
	for _, val := range nums {
		if candidate == val {
			count += 1
		}
	}
	if count > len(nums)/2 {
		return candidate
	} else {
		return 0
	}
}
