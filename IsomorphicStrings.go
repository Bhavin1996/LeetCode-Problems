/*
Problem Statement :
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character, but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true

Example 2:

Input: s = "foo", t = "bar"
Output: false

Example 3:

Input: s = "paper", t = "title"
Output: true


Problem breakdown

# Anyone facing issue in understanding the description of this question.
# You must understand the this question needs us to check if there is one directional map exist between given string and a pattern.
# In the given question, we can treat 's' as given string and 't' as pattern.
# So, essentially, we have to check if there is a one-directional map exist from 's' to 't'.

# Whatever character is there in 's' is not same as in 't'. For instance, s = 'PAPER', t = 'TITLE'.
# If we can see, character 'E' is common between string 's' and pattern 't'. Character 'E' from string 's' is mapped to character 'L' in pattern 't'.
# Meanwhile, character 'E' in the pattern 't' is being mapped to character 'R' from string 's'.
string (s)	pattern (t)
P	T
A	I
E	L
R	E

So, as long as, character from string 's', have only one map towards character in pattern 't', the strings are isomorphic.

Approach :

*/

package leetcode

func isIsomorphic(s string, t string) bool {
	sLen := len(s)
	mapS := [256]byte{}
	mapT := [256]byte{}

	for i := 0; i < sLen; i++ {
		if mapS[s[i]] != mapT[t[i]] {
			return false
		}
		mapS[s[i]] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)
	}
	return true
}

/*
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sLen := len(s)
	mapS := [256]byte{}
	mapT := [256]byte{}

	for i := 0; i < sLen; i++ {
		// Print the current characters and their mappings
		fmt.Printf("s[%d] = %c, t[%d] = %c\n", i, s[i], i, t[i])
		fmt.Printf("mapS[%c] = %d, mapT[%c] = %d\n", s[i], mapS[s[i]], t[i], mapT[t[i]])

		if mapS[s[i]] != mapT[t[i]] {
			fmt.Println("Mapping mismatch detected.")
			return false
		}

		// Update the mappings
		mapS[s[i]] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)

		// Print the updated mappings
		fmt.Printf("Updated mapS[%c] = %d, mapT[%c] = %d\n", s[i], mapS[s[i]], t[i], mapT[t[i]])
	}
	return true
}

func main() {
	fmt.Println("Hello, 世界")
	res := isIsomorphic("paper", "title")
	fmt.Println(res)
}
*/
