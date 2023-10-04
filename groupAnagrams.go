/*
All the characters of the english alphabet have their own ordinal number assigned to them.
Since anagrams have the same characters, we can leverage this property to create a unique signature to tag an anagram group.
Approach

Started with the most basic summation of the ordinals

tan
t ~ int x in ordinal
a ~ int y in ordinal
n ~ int z in ordinal

Then calculated the sum of all the ordinals. Well this wasn't enough to keep all the signatures distinct.

Tried the same with the products of the ordinals. Still wasn't enough.

Then went with the combination of the two methods in order to create a unique fingerprint that identifies an anagram.
*/

package leetcode

func groupAnagrams(strs []string) (result [][]string) {
	anagramGroups := make(map[int][]string)
	for _, d := range strs {
		prod := 1
		sum := 0
		for _, sd := range d {
			prod *= int(sd)
			sum += int(sd)
		}

		signature := prod * sum
		anagramGroups[signature] = append(anagramGroups[signature], d)
	}

	for _, v := range anagramGroups {
		result = append(result, v)
	}

	return
}

/*
type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
*/
