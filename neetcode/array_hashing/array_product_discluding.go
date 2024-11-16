package array_hashing

import "errors"

/*
Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
Each product is guaranteed to fit in a 32-bit integer.
Follow-up: Could you solve it in O(n) time without using the division operation?
Example 1:
Input: nums = [1,2,4,6]
Output: [48,24,12,8]
Example 2:
Input: nums = [-1,0,1,2,3]
Output: [0,-6,0,0,0]
Constraints:
    2 <= nums.length <= 1000
    -20 <= nums[i] <= 20
*/

func ProductWithoutIncludingSelf(arr []int) ([]int, error) {
	arrLen := len(arr)
	if arrLen < 2 {
		return nil, errors.New("invalid Input, Input array length should be >= 2")
	}
	leftProducts := make([]int, arrLen)
	rightProducts := make([]int, arrLen)
	leftProducts[0] = 1
	rightProducts[0] = 1

	sol := make([]int, arrLen)
	for idx := 1; idx < arrLen; idx++ {
		leftProducts[idx] = leftProducts[idx-1] * arr[idx-1]
		rightProducts[idx] = rightProducts[idx-1] * arr[arrLen-idx]
	}

	for i := 0; i < arrLen; i++ {
		sol[i] = leftProducts[i] * rightProducts[arrLen-i-1]
	}
	return sol, nil
}
