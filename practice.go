package leetcode

import "fmt"

func RemoveDuplicates1(nums []int) int {
	val := 2
	nums = append(nums[:val], nums[val+1:]...)
	fmt.Println(nums)
	return len(nums)
}
