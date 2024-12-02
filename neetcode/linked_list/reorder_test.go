package linked_list

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test:1",
			args: args{head: combinedList()},
			want: getReversedTest(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getReversedTest() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 7}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	return head
}
