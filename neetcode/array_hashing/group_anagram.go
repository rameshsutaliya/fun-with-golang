package array_hashing

import "sort"

/*
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.
An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
Example 1:
Input: strs = ["act","pots","tops","cat","stop","hat"]
Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:
Input: strs = ["x"]
Output: [["x"]]
Example 3:
Input: strs = [""]
Output: [[""]]
Constraints:
1 <= strs.length <= 1000.
0 <= strs[i].length <= 100
strs[i] is made up of lowercase English letters.
*/

func GroupAnagram(list []string) [][]string {
	sol := make([][]string, 0)
	anagramMap := make(map[string][]string)
	for _, str := range list {
		slicedStr := []rune(str)
		sort.Slice(slicedStr, func(i, j int) bool {
			return slicedStr[i] < slicedStr[j]
		})
		sortedStr := string(slicedStr)
		anagrams, ok := anagramMap[sortedStr]
		if !ok {
			anagrams = make([]string, 0)
		}
		anagrams = append(anagrams, str)
		anagramMap[sortedStr] = anagrams
	}

	for _, value := range anagramMap {
		sol = append(sol, value)
	}

	return sol
}
