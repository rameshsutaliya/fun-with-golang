package trees

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test:1",
			args: args{root: getBinaryTree()},
			want: testInvertedBT(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getBinaryTree() *TreeNode {
	head := &TreeNode{Val: 2}
	head.Left = &TreeNode{Val: 1}
	head.Right = &TreeNode{Val: 3}
	return head
}

func testInvertedBT() *TreeNode {
	head := &TreeNode{Val: 2}
	head.Left = &TreeNode{Val: 3}
	head.Right = &TreeNode{Val: 1}
	return head
}
