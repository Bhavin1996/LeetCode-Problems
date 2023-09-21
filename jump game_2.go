package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		return 1
	}
	end := len(nums) - 1
	maxJump := nums[0]
	numberOfJumps := 0
	count := 0
	jumpDist := nums[0]
	i := 1
	for i < end {
		max := i + nums[i]
		if max >= end {
			return numberOfJumps + 1
		}
		count += 1
		if max > maxJump {
			maxJump = max
			if count == jumpDist {
				jumpDist = nums[maxJump]
				numberOfJumps += 1
				i = maxJump
				maxJump = i + nums[i]
				count = 0
			}
		}
	}
}
