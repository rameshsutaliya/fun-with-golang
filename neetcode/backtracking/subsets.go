package backtracking

/*
Given an array nums of unique integers, return all possible subsets of nums.
The solution set must not contain duplicate subsets. You may return the solution in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [7]
Output: [[],[7]]

Constraints:
    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
*/

func subsets(nums []int) [][]int {
	subs := make([][]int, 0)
	var bt func([]int, int)
	bt = func(ints []int, idx int) {
		if idx == len(nums) {
			subs = append(subs, ints)
			return
		}
		ints = append(ints, nums[idx])
		bt(ints, idx+1)
		ints = ints[:len(ints)-1]
		bt(ints, idx+1)
	}
	bt([]int{}, 0)

	return subs
}
