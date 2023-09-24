/* Problem statement :

You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

    0 <= j <= nums[i] and
    i + j < n

Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].



Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [2,3,0,1,4]
Output: 2
solution :
# This is a next variation of jump game problem
# In this we need to calculate the minimum number of jumps to reach n-1 index
# To achieve this, the approach that needs to be followed is as below
# The maxJump We can achieve is ith index + value at that index
# Now we jump to next marker if value of i is equal to maxJump, put check for conditions and we have our answer
*/

package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		return 1
	}
	currJump, maxJump, numberOfJumps := 0, 0, 0
	for i := 0; i < len(nums); i++ {

		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}

		if i == currJump {
			currJump, numberOfJumps = maxJump, numberOfJumps+1
			if currJump >= len(nums)-1 {
				return numberOfJumps
			}
		}
	}
	return numberOfJumps
}
