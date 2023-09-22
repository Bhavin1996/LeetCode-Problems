/*Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true

#solution :
# Create a hash map with magazine string as a base beacuse we are going to check if it has form ransomNote
# Now map structure will be with key as rune i.e. each charactor in string which is called rune in GO
# The value would be occuracnce of each rune in string.
# ch-'a' gives us integer values for the runes (eg. if ch == 'a' then a-a = 0 and so on)
# Now that this hash is created we go for one more loop with ransomNote string (loop range := len(ransomNote))
# We iterate over it in a way if that rune is present as key in HashMap we reduce the value associated with it
# So for any instance if value goes below zero means we can't form the required string phrase.
# If not just return true after done with loop.
*/

package leetcode

/*func canConstruct(ransomNote string, magazine string) bool {
	ransomMap := make(map[rune]int)
	magazineMap := make(map[rune]int)
	for _, ch := range ransomNote {
		ransomeMap[ch-'a']++
	}
	for _, ch := range magazine {
		magazineMap[ch-'a']++
	}
	for key, _ := range ransomeMap {
		if ransomeMap[key] > magazineMap[key] {
			return false
		}
	}
	return true
}*/

// # This one is more optimal with less space and one less loop

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, ch := range magazine {
		magazineMap[ch-'a']++
	}
	for _, ch := range ransomNote {
		magazineMap[ch-'a']--
		if magazineMap[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
