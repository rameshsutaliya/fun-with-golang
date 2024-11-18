package stack

import "fun-with-golang/base/stack"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"
Output: true

Constraints:
    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.

*/

/*
Approach: Use the stack and push the all the open type parentheses and pop when close type parentheses comes and check, is it match? than continue else return false.
*/
func IsValidParentheses(str string) bool {
	charMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	s := stack.NewStack[rune]()
	for _, char := range str {
		if char == '(' || char == '{' || char == '[' {
			s.Push(char)
		} else if s.IsEmpty() {
			return false
		} else {
			topVal, err := s.Pop()
			closeChar, _ := charMap[topVal]
			if err != nil {
				return false
			} else if closeChar != char {
				return false
			}
		}
	}

	return s.IsEmpty()
}
