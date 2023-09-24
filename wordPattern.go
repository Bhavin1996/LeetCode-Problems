/* Problem :
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.



Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true

Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false

# Use HashMap
*/

package leetcode

import "strings"

//words := strings.Fields(s) stores by word not ch and eliminates space
// e. g. s = "dog cat cat dog"  so words = [dog cat man] words[0] = dog

func wordPattern(pattern string, s string) bool {
	m1 := make(map[string]int)
	m2 := make(map[byte]int)
	splited := strings.Split(s, " ")
	if len(splited) != len(pattern) {
		return false
	}
	cnt := 1
	for i := 0; i < len(splited); i++ {
		x := m2[pattern[i]]
		y := m1[splited[i]]
		if x != y {
			return false
		}
		if x == 0 {
			m2[pattern[i]] = cnt
			m1[splited[i]] = cnt
			cnt++
		}
	}
	return true
}
