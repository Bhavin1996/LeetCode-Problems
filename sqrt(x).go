/*Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Solution :
Algorithm based on binary search. Use concept of binary search to find the square root of number
*/

package leetcode

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x+1

	for left < right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid
		} else if mid*mid == x {
			return mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
