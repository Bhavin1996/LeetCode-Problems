/*
In this code:

	    We handle special cases where numRows is 1 or numRows is greater than or equal to the length of the input string,
		in which case the result is the input string itself.

	    We create a list of strings (rows) to represent each row in the zigzag pattern.

	    We iterate through the characters in the input string,
		placing each character in the appropriate row according to the zigzag pattern.

	    We update the current row and direction as we move through the pattern.

	    Finally, we concatenate the rows to get the zigzag pattern result.

For the input string "PAYPALISHIRING" and numRows = 3, the output will be "PAHNAPLSIIGYIR," as expected.
You can call the convert function with different inputs to test the solution.
*/
package leetcode

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Create numRows strings to represent the rows in the zigzag pattern
	rows := make([]string, numRows)

	// Initialize variables for tracking the current row and direction
	currentRow := 0
	direction := 1 // 1 for downward, -1 for upward

	// Iterate through the characters in the input string
	for _, char := range s {
		rows[currentRow] += string(char)

		// Change direction when reaching the top or bottom row
		if currentRow == 0 {
			direction = 1
		} else if currentRow == numRows-1 {
			direction = -1
		}

		currentRow += direction
	}

	// Concatenate the rows to form the zigzag pattern result
	result := strings.Join(rows, "")

	return result
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	converted := convert(s, numRows)
	fmt.Println(converted)
}

/*
Intuition

First, I decided to draw the scheme this way.
There we can see that the difference between the indexes of neighboring letters is defined by a simple dependency.
Approach

In this case I decided to use strings.Builder because of its efficiency. To start, we need to run for rows of this structure from 0 to numRows - 1. Inside this cycle we start new and append s[i] to res , wrere i is updated index by formulas.Until i < n we update it i += 2 * (numRows - r - 1) first and i += 2 * r second.
We can see that in the cases of the first and last rows, we are typing the same letter twice, so we create special index updates for those cases.
Complexity

    Time complexity:
    O(n)O(n)O(n)

    Space complexity:
    O(n)O(n)O(n
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I




package leetcode

func convert(s string, numRows int) string {
	n := len(s)

	if numRows == 1 || n < numRows {
		return s
	}

	var i, tmp1, tmp2 int
	var res strings.Builder

	res.Grow(n)

	for r := 0; r < numRows; r++ {
		tmp1 = 2 * (numRows - r - 1)
		tmp2 = 2 * r

		if tmp2 == 0 {
			for i = r; i < n; i += tmp1 {
				res.WriteByte(s[i])
			}
		} else if tmp1 == 0 {
			for i = r; i < n; i += tmp2 {
				res.WriteByte(s[i])
			}
		} else {
			res.WriteByte(s[r])
			for i = r + tmp1; i < n; i += tmp1 {
				res.WriteByte(s[i])
				i += tmp2
				if i < n {
					res.WriteByte(s[i])
				} else {
					break
				}
			}
		}

	}

	return res.String()
}
*/
