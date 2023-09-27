/*Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

    Starting with any positive integer, replace the number by the sum of the squares of its digits.
    Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
    Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.



Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

Example 2:

Input: n = 2
Output: false
*/

package leetcode

func isHappy(n int) bool {
	slow, fast := n, n
	slow = getSum(slow)
	fast = getSum(getSum(fast))

	for slow != fast {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
	}

	return slow == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		remainder := n % 10
		sum += remainder * remainder
		n /= 10
	}

	return sum
}

/*

// nextNumber transforms a given number `n` to a new number by summing the squares of its individual digits
// time commplexity is O(d), where d is the number of digits in n. Or O(log(n)) because d <= log(n) + 1
func nextNumber(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10;
        sum += digit * digit;
        n = n / 10;
    }
    return sum
}


func isHappy(n int) bool {

    // bitmap stores information on which number (< 100) has been already visited.
    // 1 number uint64 can stores 64 bits, so to store 100 bits, we need an array of two number
    bitmap := []uint64{0, 0}

    for n != 1 {
        if n < 64 {
            // if n < 64, its bit information is stored on the first number in the array, thus bitmap[0]
            if bitmap[0] & (1 << n) > 0 {
                // if the bit at location n is already set, which mean the location n has been visited previosly => there is a loop => return false.
                return false
            }
            //first time visited, set the bit to 1
            bitmap[0] = bitmap[0] | (1 << n)
        } else if n < 100 {

            // same logic as the above if statement, but the information is stored in the second number.

            if bitmap[1] & (1 << (n % 64)) > 0 {
                return false
            }
            bitmap[0] = bitmap[0] | (1 << (n % 64))
        }
        n = nextNumber(n)

    }

    return true

}*/
