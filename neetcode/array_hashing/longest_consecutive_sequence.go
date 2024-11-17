package array_hashing

func GetSequenceCount(arr []int) int {
	numMap := make(map[int]bool)
	for _, num := range arr {
		numMap[num] = true
	}

	maxVal, localMax := 0, 0
	for _, val := range arr {
		isNotVisited, ok := numMap[val]
		if ok && isNotVisited {
			localMax = 1 + getLowerCount(val, numMap) + getUpperCount(val, numMap)
		}

		if localMax > maxVal {
			maxVal = localMax
		}
	}

	return maxVal
}

func getLowerCount(val int, nMap map[int]bool) int {
	// Recursive function call stack issue with online platform due to memory stack.
	count := 0
	val -= 1
	_, ok := nMap[val]
	for ok {
		count += 1
		val -= 1
		_, ok = nMap[val]
		if ok {
			nMap[val] = false
		} else {
			return count
		}
	}

	return count
	// Commenting the recursive implementation
	/*
		if ok {
			nMap[val-1] = false
			return 1 + getLowerCount(val-1, nMap)
		} else {
			return 0
		}
	*/
}

func getUpperCount(val int, nMap map[int]bool) int {
	// Recursive function call stack issue with online platform due to memory stack.
	count := 0
	val += 1
	_, ok := nMap[val]
	for ok {
		count += 1
		val += 1
		_, ok = nMap[val]
		if ok {
			nMap[val] = false
		} else {
			return count
		}
	}

	return count
	/*
		_, ok := nMap[val+1]
		if ok {
			nMap[val+1] = false
			return 1 + getUpperCount(val+1, nMap)
		} else {
			return 0
		}
	*/
}

/*
Improved Version can be looked as:
map[int]int
where map[num]= 1+map[num-1]+map[num+1]
sol will be the max value of this map.
*/
