/*
	This is a standard solution to missing number problem where a number is missing and we need to find it

but the catch here is that it has to be a O(1) solution below A.P method does the work

package Leetcode

import (
	"fmt"
)

func MissingNumber(ls []int, size int) {
	expectedSum := (size * (size + 1)) / 2
	actualSum := 0
	for _, num := range ls {
		actualSum += num
	}
	missingNumberIs := expectedSum - actualSum
	fmt.Printf("Missing number is %v", missingNumberIs)
}

func Missing() {
	var x int
	_, _ = fmt.Scan(&x) // Read the number of elements (size of the slice)
	ls := make([]int, x)
	for i := 0; i < x; i++ {
		_, _ = fmt.Scan(&ls[i]) // Read the value directly into the slice element
	}
	MissingNumber(ls, x)
}*/
