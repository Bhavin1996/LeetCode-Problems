/* Problem :
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]
solution :
# This is solved by two-ponter approach where the second pointer increments by 1 only on non-zero values
# After encountering such value the swap happens i.e non-zero value gets swaped with zero value at that ith index
# Time complexity O(n) and space complexity O(1)
*/

package leetcode

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		} else {
			continue
		}
	}
}
