package leetcode

func lengthOfLongestSubstring(s string) int {
	lookup := make(map[rune]int, 26)
	count, maxCount := 0, 0
	start := 0

	for i, ch := range s {
		if pos, isPresent := lookup[ch]; isPresent && pos >= start {
			start = pos + 1
		}

		count = i - start + 1
		lookup[ch] = i

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
