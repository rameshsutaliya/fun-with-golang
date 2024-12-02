package linked_list

/*
You are given the beginning of a linked list head, and an integer n.
Remove the nth node from the end of the list and return the beginning of the list.

Example 1:
Input: head = [1,2,3,4], n = 2
Output: [1,2,4]

Example 2:
Input: head = [5], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 2
Output: [2]

Constraints:
    The number of nodes in the list is sz.
    1 <= sz <= 30
    0 <= Node.val <= 100
    1 <= n <= sz
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	// [1,2,3,4], n = 2
	// first = 1, n=2
	// first =2, n=1
	// first = 3, n = 0
	// first = 4, n=-1
	for n >= 0 && first != nil {
		n--
		first = first.Next
	}
	if first == nil && n == 0 {
		return head.Next
	}
	// second = 1
	// first = 4, second = 2
	// first = nil, second = 3
	second := head
	for first != nil {
		second = second.Next
		first = first.Next
	}
	// [1,2,4,nil]
	temp := second.Next
	second.Next = temp.Next
	return head

}
