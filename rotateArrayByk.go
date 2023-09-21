/* simple inplementation of slicing based on index depending on the requirement*/
package leetcode

func rotate(nums []int, k int) {
	end := len(nums)
	modifiedK := end - k%end
	copy(nums[:], append(nums[modifiedK:], nums[:modifiedK]...))
}

/*
alternate solution
func rotate(nums []int, k int) {
    n := len(nums)
    k %= n
    reverse(nums, 0, n - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, n - 1)
}

func reverse(nums []int, start int, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}

*/
