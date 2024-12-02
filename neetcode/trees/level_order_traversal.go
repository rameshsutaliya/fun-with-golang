package trees

import "fun-with-golang/base/queue"

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Constraints:
    The number of nodes in the tree is in the range [0, 2000].
    -1000 <= Node.val <= 1000
*/

type Queue struct {
}

func levelOrder(root *TreeNode) [][]int {
	sol := make([][]int, 0)
	if root == nil {
		return sol
	}
	q := queue.NewQueue[*TreeNode]()
	q.Enqueue(root)
	for !q.IsEmpty() {
		list := make([]int, 0)
		for count := q.Size(); count > 0; count-- {
			top, err := q.Dequeue()
			if err != nil {
				return sol
			}
			if top != nil {
				list = append(list, top.Val)
				q.Enqueue(top.Left)
				q.Enqueue(top.Right)
			}
		}
		if len(list) > 0 {
			sol = append(sol, list)
		}
	}
	
	return sol
}
