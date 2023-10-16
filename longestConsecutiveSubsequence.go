/*
Algorithm:
1. Create an empty hash.
2. Insert all array elements to hash.
3. Do following for every element arr[i]
4. Check if this element is the starting point of a subsequence. To check this,
simply look for arr[i] 1 in the hash, if not found, then this is the first element a subsequence.
5. If this element is the first element,
then count the number of elements in the consecutive starting with this element.
Iterate from arr[i] + 1 till the last element that can be found.
6. If the count is more than the previous longest subsequence found, then update this.
*/

package leetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	lookup := make(map[int]int)
	maxCount := 0
	for _, num := range nums {
		lookup[num] = 1
	}

	for n := range lookup {
		// Below if_else block is to reduce unwanted iteration as no two sequence can have
		// value greater that half of size of array
		if maxCount == len(nums)-1 {
			return maxCount
		} else if maxCount > len(nums)-1/2 {
			return maxCount
		}
		if _, ok := lookup[n-1]; !ok { // This check if it is the starting point of sub-sequence
			count := 1
			for { // This loop gives us all subsquence for current n
				if _, ok := lookup[n+count]; ok {
					count += 1
					continue
				}
				maxCount = maxNum(count, maxCount) // Checks max len of subsequence
				break
			}
		}
	}

	return maxCount
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
