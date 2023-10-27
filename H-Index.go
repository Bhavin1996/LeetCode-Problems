package leetcode

//https://en.wikipedia.org/wiki/H-index

import "sort"

func hIndex(citations []int) int {
	hIndex := 0
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := 0; i < len(citations); i++ {
		if i+1 <= citations[i] {
			hIndex = i + 1
		} else {
			return hIndex
		}
	}
	return hIndex
}
