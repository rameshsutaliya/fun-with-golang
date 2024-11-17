package array_hashing

import (
	"strings"
	"unicode"
)

func IsValidPalindrome(str string) bool {
	str = strings.ToLower(str)
	left, right := 0, len(str)-1
	for left < right {
		if !(unicode.IsLetter(rune(str[left])) || unicode.IsNumber(rune(str[left]))) {
			left += 1
		} else if !(unicode.IsLetter(rune(str[right])) || unicode.IsNumber(rune(str[right]))) {
			right -= 1
		} else if str[left] == str[right] {
			left += 1
			right -= 1
		} else {
			return false
		}
	}

	return true
}
