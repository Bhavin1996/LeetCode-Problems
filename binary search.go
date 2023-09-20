package leetcode

import "fmt"

func binarySearch(start int, end int, arr []int) int {
	var val int
	fmt.Println("Enter the value that you need to search")
	fmt.Scan(&val)
	for start <= end {
		mid := (start + end) / 2
		if val == arr[mid] {
			return mid
		} else if val < arr[mid] {
			end = mid - 1
		} else if val > arr[mid] {
			start = mid + 1
		}
	}
	return -1
}

func main1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := binarySearch(0, len(arr)-1, arr[:])
	fmt.Println("The varaible was found in array at index :", res)
}
