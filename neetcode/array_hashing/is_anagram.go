package array_hashing

/*
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
Example 1:
Input: s = "racecar", t = "carrace"
Output: true
Example 2:
Input: s = "jar", t = "jam"
Output: false
Constraints:
`s` and `t` consist of lowercase English letters.

*/

func IsAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	charCount := make([]int, 26)
	for idx := 0; idx < len(first); idx++ {
		charCount[first[idx]-'a']++
		charCount[second[idx]-'a']--
	}

	for _, num := range charCount {
		if num != 0 {
			return false
		}
	}

	return true
}
