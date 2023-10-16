package leetcode

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, value := range arr[1:] {
		if value <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, pivot), greater...)
}

func _main() {
	arr := []int{5, 2, 9, 3, 6, 1, 8, 4, 7}

	sorted := quickSort(arr)
	fmt.Println("Sorted array:", sorted)
}
