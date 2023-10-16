package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)
	if n == 0 {
		return res
	}
	start := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-1 == nums[i-1] {
			continue
		}
		if start == nums[i-1] {
			res = append(res, fmt.Sprintf("%d", start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i-1]))
		}
		start = nums[i]
	}

	if nums[n-1] == start {
		return append(res, fmt.Sprintf("%d", start))
	}
	return append(res, fmt.Sprintf("%d->%d", start, nums[n-1]))
}
