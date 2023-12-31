/*You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
[4,2,0,0,1,1,4,4,4,0,4,0]
*/

package leetcode

import "fmt"

func max1(arr []int) (b int) {
	a := arr[0]
	b = 0
	count := a
	for i := 1; i < len(arr); i++ {
		if arr[i] >= a {
			a = arr[i]
			b = i
			count = a
		} else if arr[i] == 0 {
			count -= 1
		}
		if count == 0 && arr[i] == 0 {
			a = 0
			b = i
			count = arr[i]
		}

	}
	return b + 1
}

func canJump(nums []int) bool {
	jumpDist := 0
	if (len(nums) == 2 && nums[0] != 0) || len(nums) == 1 {
		return true
	} else if len(nums) == 2 && nums[0] == 0 {
		return false
	}
	i := 0
	for i < len(nums) {
		jumpDist = i + nums[i]
		if nums[i] == 0 && i != len(nums)-1 {
			return false
		}
		if jumpDist >= len(nums)-1 {
			return true
		} else {
			index := max1(nums[i+1 : jumpDist+1])
			i += index
			fmt.Println(i)
		}
	}
	return true
}

func main2() {
	arr := []int{4, 2, 0, 0, 1, 1, 4, 4, 4, 0, 4, 0}
	if canJump(arr[:]) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

/* optimal solution

func canJump(nums []int) bool {
    // the best I can reach
    farthest := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if i + nums[i] > farthest {
            farthest = i + nums[i]
        }

        if nums[i] == 0 && i < n - 1 && i == farthest {
            return false
        }
    }

    return true
}*/
