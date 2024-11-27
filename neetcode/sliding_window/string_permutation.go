package sliding_window

/*
You are given two strings s1 and s2.
Return true if s2 contains a permutation of s1, or false otherwise. That means if a permutation of s1 exists as a substring of s2, then return true.
Both strings only contain lowercase letters.

Example 1:
Input: s1 = "abc", s2 = "lecabee"
Output: true
Explanation: The substring "cab" is a permutation of "abc" and is present in "lecabee".

Example 2:
Input: s1 = "abc", s2 = "lecaabee"
Output: false

Constraints:
    1 <= s1.length, s2.length <= 1000
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	charCount := make([]int, 26)
	s1Len := len(s1)
	for i := 0; i < s1Len; i++ {
		charCount[s1[i]-'a'] += 1
	}
	tempCount := make([]int, 26)
	for i := 0; i < s1Len; i++ {
		tempCount[s2[i]-'a'] += 1
	}
	if isEqual(charCount, tempCount) {
		return true
	}
	for i := s1Len; i < len(s2); i++ {
		tempCount[s2[i-s1Len]-'a'] -= 1
		tempCount[s2[i]-'a'] += 1
		if isEqual(charCount, tempCount) {
			return true
		}
	}

	return false
}

func isEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
