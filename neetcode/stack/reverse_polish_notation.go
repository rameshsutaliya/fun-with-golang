package stack

import (
	"fun-with-golang/base/stack"
	"strconv"
)

/*
You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.
Return the integer that represents the evaluation of the expression.
    The operands may be integers or the results of other operations.
    The operators include '+', '-', '*', and '/'.
    Assume that division between integers always truncates toward zero.
Example 1:
Input: tokens = ["1","2","+","3","*","4","-"]
Output: 5
Explanation: ((1 + 2) * 3) - 4 = 5
Constraints:
    1 <= tokens.length <= 1000.
    tokens[i] is "+", "-", "*", or "/", or a string representing an integer in the range [-100, 100].
*/

func EvalRPN(tokens []string) int {
	s := stack.NewStack[int]()
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			s.Push(val)
		} else {
			right, err := s.Pop()
			if err != nil {
				return -1
			}
			left, err := s.Pop()
			if err != nil {
				return -1
			}

			switch token {
			case "+":
				s.Push(left + right)
			case "-":
				s.Push(left - right)
			case "/":
				s.Push(left / right)
			case "*":
				s.Push(left * right)
			}
		}
	}

	sol, err := s.Pop()
	if err != nil {
		return -1
	}
	return sol
}
