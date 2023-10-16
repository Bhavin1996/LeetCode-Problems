package leetcode

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	start := intervals[0][0]
	end := 0
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			end = intervals[i+1][1]
		} else {
			res = append(res, []int{intervals[i+1][1]})
			start = intervals[i+1][0]
		}
	}
	return res
}
