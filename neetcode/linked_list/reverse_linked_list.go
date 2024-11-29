package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	temp := head
	var last *ListNode
	reverse := temp
	for temp.Next != nil {
		currTemp := temp
		temp = temp.Next
		reverse.Next = last
		reverse = temp
		last = currTemp
	}
	return reverse
}
