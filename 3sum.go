package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// Sort the input array in ascending order
	sort.Ints(nums)

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate values
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate values
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
