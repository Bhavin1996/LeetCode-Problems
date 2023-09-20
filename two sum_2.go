package leetcode

import "sort"

func twoSum2(nums []int, target int) []int {
	res := make([]int, 2)
	sum := 0
	sort.Ints(nums)
	j := len(nums) - 1
	i := 0
	for i < j {
		sum = nums[i] + nums[j]
		if (sum) < target {
			i += 1
		} else if (sum) > target {
			j -= 1
		} else if (sum) == target {
			res[0] = i + 1
			res[1] = j + 1
		}
	}
	return res
}
