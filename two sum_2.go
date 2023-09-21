/* This problem is solved by the two pointer approach in which, which is as follows that we assign the one pointer to the first index of array
and second pointer to the last element.
# Keep note that this only works on the sorted elements
# Now after assigning we do addition of both elements based on two pointers we have initialised
# If the sum is smaller than target increase the first element
# and if not then decrease the second pointer.
# Do this until u get target value and return the index of elements at that point

*/

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
