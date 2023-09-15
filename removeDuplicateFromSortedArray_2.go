package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	index := 0
	i := 1
	if len(nums) == 2 {
		return len(nums)
	}
	for {
		if index < len(nums)-1 {
			return int(len(nums))
		}
		if nums[i] == nums[i-1] && count < 1 {
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count < 2 {
			index = i
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count >= 2 {
			i += 1
		} else if nums[i] != nums[i-1] {
			if count == 1 {
				nums = append(nums[:i], nums[i:]...)
				fmt.Println(nums)
				if index < len(nums)-1 {
					return int(len(nums))
				}
				count = 0
				i += 1
			} else if count == 2 {
				nums = append(nums[:index], nums[i:]...)
				fmt.Println(nums)
				count = 0
			} else if count > 2 {
				nums = append(nums[:index], nums[i:]...)
				fmt.Println(nums)
				if index < len(nums)-1 {
					return int(len(nums))
				}
				i = index + 1
				count = 0
			}
		}
	}
}
