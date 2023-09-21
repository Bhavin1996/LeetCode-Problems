/*This problem is solved using hash-map, where basically we check
if the target_Sum - curr_elemetn is present or not
if it is then return the index of those
else add them in map with structure as map[array_element][current_index of ele]
*/

package leetcode

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
