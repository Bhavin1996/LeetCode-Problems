// # Explanation to approach https://www.youtube.com/watch?v=gREVHiZjXeQ

package leetcode

func productExceptSelf(nums []int) []int {
	resArray := make([]int, len(nums))
	if len(nums) < 1 {
		return resArray
	}

	product := 1
	for _, num := range nums {
		product *= num
		resArray = append(resArray, product)
	}
	product = 1

	for i := len(nums) - 1; i > 0; i-- {
		resArray[i] = product * resArray[i-1]
		product *= nums[i]
	}
	resArray[0] = product
	return resArray
}
