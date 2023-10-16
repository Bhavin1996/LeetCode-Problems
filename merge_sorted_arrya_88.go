package leetcode

/*Solution to this porblem is that you need to traverse from back i.e. from last index while untill n != 0
Also need to check the bigger element and put it 1st list from last index
corner cases will be one of list having no elements and duplicacy of the elements in a list
two approaches are mentioned below
*/

//brute force with space uses
/*
func merge(nums1 []int, m int, nums2 []int, n int) {
	var combine []int
	combine = append(combine, nums1[:]...)
	combine = append(combine, nums2[:]...)

	sort.Slice(combine, func(i, j int) bool {
		return combine[i] < combine[j]
	})
	var nonZeroNumbers []int
	for _, num := range combine {
		if num != 0 {
			nonZeroNumbers = append(nonZeroNumbers, num)
		}
	}
	for i, num := range nonZeroNumbers {
		nums1[i] = num
	}
}*/

// O(m + n) approach

func mergeArrays(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
