//https://www.youtube.com/watch?v=bTauscvOymA

package leetcode

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
