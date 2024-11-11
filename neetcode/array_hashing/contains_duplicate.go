package array_hashing

/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.
*/

func ContainsDuplicate(arr []int) bool {
	hashSet := make(map[int]bool)
	for _, num := range arr {
		_, exists := hashSet[num]
		if exists {
			return true
		} else {
			hashSet[num] = true
		}
	}
	return false
}
