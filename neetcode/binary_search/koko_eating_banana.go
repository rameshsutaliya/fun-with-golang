package binary_search

/*
You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.
You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.
Return the minimum integer k such that you can eat all the bananas within h hours.

Example 1:
Input: piles = [1,4,3,2], h = 9
Output: 2
Explanation: With an eating rate of 2, you can eat the bananas in 6 hours. With an eating rate of 1, you would need 10 hours to eat all the bananas (which exceeds h=9), thus the minimum eating rate is 2.

Example 2:
Input: piles = [25,10,23,4], h = 4
Output: 25

Constraints:
    1 <= piles.length <= 1,000
    piles.length <= h <= 1,000,000
    1 <= piles[i] <= 1,000,000,000
*/

func MinimumEatingRate(piles []int, hours int) int {
	// Approach 1:
	/*
		rate start with 1, and hours reduce by piles[i]/rate and iff piles got iterated till end than return rate else
		start with rate+1 and from the 0 idx again.

		Time: O(m*n)
	*/
	// Approach 2: as we can notice in approach 1, rate is already sorted than we can approach towards binary search.

	left, right := 1, 1_000_000_000
	for left <= right {
		mid := left + (right-left)/2
		total := 0
		for _, pile := range piles {
			total += pile / mid
			if pile%mid != 0 {
				total += 1
			}
			if total > hours {
				left = mid + 1
				break
			}
		}
		if total <= hours {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
