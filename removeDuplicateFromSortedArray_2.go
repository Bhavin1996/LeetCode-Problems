/*This problem is solved on basis of slicing techniques based on index and count of dupplicate.
When we have a count of more than one just store that index and use it to swap the main array
with element as soon as we find the mis-matched element. keep a check of corner cases. */

package leetcode

func RemoveDuplicates(nums []int) int {
	count := 0
	index := 0
	i := 1
	if len(nums) == 2 || len(nums) == 1 {
		return len(nums)
	}
	for {
		if nums[i] == nums[i-1] {
			count += 1
			if count == 2 {
				index = i
				if i == len(nums)-1 {
					nums = nums[:i]
					return len(nums)
				}
				i += 1
			} else if count > 2 {
				if i == len(nums)-1 {
					nums = nums[:index]
					return len(nums)
				}
				i += 1
			} else {
				if i == len(nums)-1 {
					nums = nums[:i+1]
					return len(nums)
				}
				i += 1
			}
		} else if nums[i] != nums[i-1] {
			if count == 0 {
				if i == len(nums)-1 {
					return len(nums)
				}
				i += 1
			} else if count == 1 {
				if i == len(nums)-1 {
					nums = nums[:i+1]
					return len(nums)
				}
				nums = append(nums[:i], nums[i:]...)
				i = i + 1
				count = 0
			} else {
				if i == len(nums)-1 {
					nums = append(nums[:index], nums[i:]...)
					return len(nums)
				}
				nums = append(nums[:index], nums[i:]...)
				count = 0
				i = index + 1
			}
		}
	}
}

/*

func removeDuplicates(nums []int) int {
	lasti := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lasti] {
			lasti++
			nums[lasti] = nums[i]
		} else if lasti == 0 || nums[lasti] != nums[lasti-1] {
			lasti++
			if lasti-1 != 0{
				nums[lasti] = nums[i]
			}
		}
	}

	return lasti + 1
}
*/
