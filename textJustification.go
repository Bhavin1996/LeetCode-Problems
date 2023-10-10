package leetcode

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var cur []string
	curLength := 0
	num_of_letters := 0
	index := 0

	for _, word := range words {
		cur = append(cur, word)
		curLength += len(word)
		num_of_letters++
		if curLength+num_of_letters > maxWidth {
			num_of_letters--
			spaces := (curLength - len(word)) / 2
			for i := 0; i < len(cur)-2; i++ {
				cur[i] = cur[i] + strings.Repeat(" ", spaces)
				fmt.Println(cur)
			}
			index++
			//cur = nil
			num_of_letters = 0
		}

	}
	return cur
}

func mai() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	res := fullJustify(words, 16)
	fmt.Println(res)
}
