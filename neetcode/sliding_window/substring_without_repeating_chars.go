package sliding_window

/*
Given a string s, find the length of the longest substring without duplicate characters.
A substring is a contiguous sequence of characters within a string.

Example 1:
Input: s = "zxyzxyz"
Output: 3
Explanation: The string "xyz" is the longest without duplicate characters.

Example 2:
Input: s = "xxxx"
Output: 1

Constraints:

	0 <= s.length <= 1000
	s may consist of printable ASCII characters
*/
func LengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res := 0, 0

	for r := 0; r < len(s); r++ {
		if idx, found := mp[s[r]]; found {
			l = max(idx+1, l)
		}
		mp[s[r]] = r
		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
