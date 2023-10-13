package leetcode

func lengthOfLongestSubstring(s string) int {
	lookup := make(map[rune]int)
	count, maxCount := 0, 0
	for _, ch := range s {
		count++
		if _, isPresent := lookup[ch]; isPresent {
			if count > maxCount {
				maxCount = count - 1
			}
			count = 0
		} else {
			lookup[ch-'a']++
		}
	}
	return maxCount
}
