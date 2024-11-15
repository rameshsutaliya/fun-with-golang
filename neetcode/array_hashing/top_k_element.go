package array_hashing

import "sort"

/*
Given an integer array nums and an integer k, return the k most frequent elements within the array.
The test cases are generated such that the answer is always unique.
You may return the output in any order.
Example 1:
Input: nums = [1,2,2,3,3,3], k = 2
Output: [2,3]
Example 2:
Input: nums = [7,7], k = 1
Output: [7]
Constraints:
    1 <= nums.length <= 10^4.
    -1000 <= nums[i] <= 1000
    1 <= k <= number of distinct elements in nums.

*/

// Approach for this problem:
// Approach 1. Create a heap where keep the count as key.
// Approach 2. Make a map and short the 2-D array.

func KTopElement(arr []int, k int) []int {
	countMap := make(map[int]int)
	// counting the frequency of array elements.
	for _, num := range arr {
		count, ok := countMap[num]
		if !ok {
			count = 0
		}
		count += 1
		countMap[num] = count
	}

	// Creating 2-D elements list
	freqArr := make([][]int, 0)
	for key, value := range countMap {
		freqArr = append(freqArr, []int{value, key})
	}

	// sorting the elements based on the frequency.
	sort.Slice(freqArr, func(i, j int) bool {
		return freqArr[i][0] > freqArr[j][0]
	})

	// Getting the top K elements from the sorted array list.
	sol := make([]int, 0)
	for i := 0; i < k; i++ {
		sol = append(sol, freqArr[i][1])
	}
	return sol
}
