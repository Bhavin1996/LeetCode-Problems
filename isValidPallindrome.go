package leetcode

import (
	"regexp"
	"strings"
	"unicode"
)

func removeAlphanumeric(input string) string {
	regex := regexp.MustCompile("[[:alnum:][:punct:]]+")
	return regex.ReplaceAllString(input, "")
}

func isPalindrome(s string) bool {
	s = removeAlphanumeric(s)
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	if len(s) == 1 {
		return true
	}
	j := len(s) - 1
	flag := false
	for i := 0; i < j; i++ {
		if s[i] == s[j] && (unicode.IsLetter(rune(s[i])) && unicode.IsLetter(rune(s[j]))) {
			flag = true
		} else if !unicode.IsLetter(rune(s[i])) && !unicode.IsLetter(rune(s[j])) {
			j--
			flag = false
			continue
		} else if !unicode.IsLetter(rune(s[i])) && unicode.IsLetter(rune(s[j])) {
			flag = false
			continue
		} else if unicode.IsLetter(rune(s[i])) && !unicode.IsLetter(rune(s[j])) {
			j--
			flag = false
			continue
		}
	}
	return flag
}
