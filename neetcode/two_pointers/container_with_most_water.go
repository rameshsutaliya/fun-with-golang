package two_pointers

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

Constraints:
    n == height.length
    2 <= n <= 105
    0 <= height[i] <= 104

*/

func ContainerVolume(arr []int) int {
	left, right := 0, len(arr)-1
	maxV, localV := 0, 0
	for left < right {
		localV = getMin(arr[left], arr[right]) * (right - left)
		if arr[left] == arr[right] {
			left += 1
			right -= 1
		} else if arr[left] < arr[right] {
			left += 1
		} else {
			right -= 1
		}

		if localV > maxV {
			maxV = localV
		}
	}

	return maxV
}

func getMin(val1, val2 int) int {
	if val1 < val2 {
		return val1
	} else {
		return val2
	}
}
