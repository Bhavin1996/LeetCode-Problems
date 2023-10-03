/*Given two integers left and right that represent the range [left, right],
return the bitwise AND of all numbers in this range, inclusive.

Example 1:

Input: left = 5, right = 7
Output: 4
Example 2:

Input: left = 0, right = 0
Output: 0
Example 3:

Input: left = 1, right = 2147483647
Output: 0
*/

package leetcode

func rangeBitwiseAnd(left int, right int) int {
	ans := left & right
	distance := right - left + 1

	for i := 0; i < 32; i++ {
		// Check if there's any zero bits within the range
		// at the corresponding position
		if (ans>>i)&1 == 1 && distance > (1<<i) {
			ans = ans ^ (1 << i)
		}
	}

	return ans
}
