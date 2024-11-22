package binary_search

/*
You are given an m x n 2-D integer array matrix and an integer target.
    Each row in matrix is sorted in non-decreasing order.
    The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.
Can you write a solution that runs in O(log(m * n)) time?
Example 1:
Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 10
Output: true

Example 2:
Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 15
Output: false

Constraints:
    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 100
    -10000 <= matrix[i][j], target <= 10000
*/

func BinarySearch(board [][]int, target int) bool {
	left, right := 0, len(board)-1
	for left <= right {
		mid := left + (right-left)/2
		if board[mid][0] == target {
			return true
		} else if board[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 {
		return false
	}

	l, r := 0, len(board[0])-1

	for l <= r {
		mid := l + (r-l)/2
		if board[right][mid] == target {
			return true
		} else if board[right][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
