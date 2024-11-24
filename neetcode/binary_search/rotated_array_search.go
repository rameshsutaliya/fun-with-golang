package binary_search

/*
You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:
    [3,4,5,6,1,2] if it was rotated 4 times.
    [1,2,3,4,5,6] if it was rotated 6 times.
Given the rotated sorted array nums and an integer target, return the index of target within nums, or -1 if it is not present.
You may assume all elements in the sorted rotated array nums are unique,
A solution that runs in O(n) time is trivial, can you write an algorithm that runs in O(log n) time?

Example 1:
Input: nums = [3,4,5,6,1,2], target = 1
Output: 4

Example 2:
Input: nums = [3,5,6,0,1,2], target = 4
Output: -1

Constraints:
    1 <= nums.length <= 1000
    -1000 <= nums[i] <= 1000
    -1000 <= target <= 1000
*/

func RotatedArraySearch(arr []int, target int) int {
	rPoint := rotationPoint(arr)
	idx := binarySearch(arr, 0, rPoint-1, target)
	if idx != -1 {
		return idx
	} else {
		return binarySearch(arr, rPoint, len(arr)-1, target)
	}
}

func rotationPoint(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] < arr[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func binarySearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := left + (right - left)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
