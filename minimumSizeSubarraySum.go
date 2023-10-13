/* Intuation : This problem is solved by the slinding window tchnique.
Where we one by one adding elemtent from array to a new stroing varable.
Take a variable as mini_value point to len(nums)

Approach :

Initialize two pointers, left and right, both initially pointing to the first element of the array.
Initialize a variable sum to keep track of the sum between these two pointers.

Initialize a variable minLen to a value larger than the maximum possible length
and a boolean variable found to false.

Start a loop that increments the right pointer until it reaches the end of the array.

In each iteration, add the element at the right pointer to the sum.

While the sum is greater than or equal to the target s,
calculate the current subarray length (right - left + 1).
If it's less than the current minLen, update minLen with this value.

Increment the left pointer to shrink the subarray by removing elements from the beginning.
Subtract the value of the element at the left pointer from the sum.

Repeat steps 3 to 6 until the right pointer reaches the end of the array.

After the loop ends, check if minLen has been updated (i.e., found is true).
If it has been updated, return minLen. Otherwise, return 0, indicating that no such subarray exists.
*/

package leetcode

func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	minLength := len(nums) + 1

	for right = 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++

		}
	}
	if minLength > len(nums) {
		return 0
	}
	return minLength
}
