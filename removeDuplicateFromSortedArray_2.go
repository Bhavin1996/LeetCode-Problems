package leetcode

func removeDuplicates2(nums []int) int {

	count := 0
	index := 0
	i := 1
	if len(nums) == 2 {
		return len(nums)
	}
	for {
		if nums[i] == nums[i-1] && count < 1 {
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count < 2 {
			index = i
			count += 1
			i++
		} else if nums[i] == nums[i-1] && count > 2 {
			i += 1
		} else if nums[i] != nums[i-1] {
			if count == 1 {
				nums = append(nums[:i], nums[i:]...)
				if index < len(nums)-1 {
					return int(len(nums))
				}
				count = 0
				i += 1
				continue
			}
			if count == 2 {
				nums = append(nums[:index], nums[i:]...)
				if index < len(nums)-1 {
					return int(len(nums))
				}
				count = 0
				continue
			}
			if count > 2 {
				nums = append(nums[:index], nums[i:]...)
				if index < len(nums)-1 {
					return int(len(nums))
				}
				i = index + 1
				count = 0
			}
		}
	}
}
