package trees

/*
Solution: Use the level order traversal and keep the last element of the level for each level.
*/

func rightSideView(root *TreeNode) []int {
	rightView := make([]int, 0)
	if root == nil {
		return rightView
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		temp := make([]*TreeNode, 0)
		for i := 0; i < len(list); i++ {
			if list[i].Left != nil {
				temp = append(temp, list[i].Left)
			}
			if list[i].Right != nil {
				temp = append(temp, list[i].Right)
			}
		}
		rightView = append(rightView, list[len(list)-1].Val)
		list = temp
	}

	return rightView
}
