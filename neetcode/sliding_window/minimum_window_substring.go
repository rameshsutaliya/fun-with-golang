package sliding_window

/*
Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".
You may assume that the correct output is always unique.

Example 1:
Input: s = "OUZODYXAZV", t = "XYZ"
Output: "YXAZ"
Explanation: "YXAZ" is the shortest substring that includes "X", "Y", and "Z" from string t.

Example 2:
Input: s = "xyz", t = "xyz"
Output: "xyz"

Example 3:
Input: s = "x", t = "xy"
Output: ""

Constraints:
    1 <= s.length <= 1000
    1 <= t.length <= 1000
    s and t consist of uppercase and lowercase English letters.
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tCharCountMap := make(map[byte]int)
	sCharCountMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := sCharCountMap[s[i]]; ok {
			sCharCountMap[s[i]] = val + 1
		} else {
			sCharCountMap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if val, ok := tCharCountMap[t[i]]; ok {
			tCharCountMap[t[i]] = val + 1
		} else {
			tCharCountMap[t[i]] = 1
		}
	}

	if !firstContainsSecond(tCharCountMap, sCharCountMap) {
		return ""
	}

	sol := s
	left, right := 0, len(t)
	subStringMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if val, ok := subStringMap[s[i]]; ok {
			subStringMap[s[i]] = val + 1
		} else {
			subStringMap[s[i]] = 1
		}
	}

	for right < len(s) || left <= right {
		if firstContainsSecond(tCharCountMap, subStringMap) && len(sol) > (right-left) {
			sol = s[left:right]
		} else {
			tVal, _ := tCharCountMap[s[left]]
			subVal, _ := subStringMap[s[left]]
			if subVal > tVal {
				subStringMap[s[left]] -= 1
				left++
			} else if right < len(s) {
				subStringMap[s[right]] += 1
				right++
			} else {
				break
			}
		}
	}

	return sol
}

func firstContainsSecond(tMap, sMap map[byte]int) bool {
	for key, val := range tMap {
		if sMap[key] < val {
			return false
		}
	}

	return true
}
