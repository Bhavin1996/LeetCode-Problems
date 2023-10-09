package leetcode

import "strings"

func addSpaces(inputString string, n int) string {
	result := inputString + strings.Repeat(" ", n)
	return result
}

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var cur []string
	num_of_letters := 0
	index := 0

	for _, word := range words {
		cur = append(cur, word)
		num_of_letters++
		if len(cur) > maxWidth {
			num_of_letters--
			l := len(cur)
			spaces := (len(cur) - len(cur[l-1]) - 1 - num_of_letters) / 2
			for i := 0; i < len(cur)-2; i++ {
				res[index] += addSpaces(words[i], spaces)
			}
			index++
			res[index] = word
			cur = nil
			cur = append(cur, word)
			num_of_letters = 0
		}

	}

	return res
}
