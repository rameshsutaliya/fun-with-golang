package stack

import "math"

/*
Design a stack class that supports the push, pop, top, and getMin operations.

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

Each function should run in O(1)O(1) time.
Example 1:
Input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]
Output: [null,null,null,null,0,null,2,1]
Explanation:
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1
Constraints:
    -2^31 <= val <= 2^31 - 1.
    pop, top and getMin will always be called on non-empty stacks
*/

type MinStack struct {
	min   int
	items []int
}

func Constructor() MinStack {
	return MinStack{
		min:   math.MaxInt64,
		items: []int{},
	}
}

func (s *MinStack) Push(val int) {
	if len(s.items) == 0 {
		s.items = append(s.items, 0)
		s.min = val
	} else {
		s.items = append(s.items, val-s.min)
		if val < s.min {
			s.min = val
		}
	}
}

func (s *MinStack) Pop() {
	if len(s.items) == 0 {
		return
	} else {
		top := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		if top < 0 {
			s.min = s.min - top
		}
	}
}

func (s *MinStack) Top() int {
	top := s.items[len(s.items)-1]
	if top > 0 {
		return top + s.min
	}
	return s.min
}

func (s *MinStack) GetMin() int {
	return s.min
}
