package leetcode

import "fmt"

func removeDuplicates2(nums []int) int {
	end := len(nums) - 1
	endElement := nums[end]
	count := 0
	index := 2
	altIndex := 0
	j := 0
	i := 2
	for j != end {
		if nums[i] == nums[i-1] && count < 1 {
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count < 2 {
			index = i
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count > 2 {
			altIndex = i
			i += 1
		} else if nums[i] != nums[i-1] {
			nums = append(nums[:index], nums[i:]...)
			i += 1
			j = i
			count = 0
			fmt.Println(nums)
		}
	}

	return int(len(nums))
}
