package leetcode

func singleNumberII(nums []int) int {
	var ones int = 0
	var twos int = 0
	var num int
	for i := 0; i < len(nums); i++ {
		num = nums[i]
		ones ^= (num & ^twos)
		twos ^= (num & ^ones)
	}

	return ones
}
