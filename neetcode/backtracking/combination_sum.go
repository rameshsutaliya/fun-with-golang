package backtracking

func combinationSum(candidates []int, target int) [][]int {
	combinationSol := make([][]int, 0)
	setList, total := make([]int, 0), 0
	var CSBT func(int)
	CSBT = func(i int) {
		if i >= len(candidates) {
			return
		}
		if total == target {
			temp := make([]int, len(setList))
			copy(temp, setList)
			combinationSol = append(combinationSol, temp)
			return
		} else if total < target {
			total += candidates[i]
			setList = append(setList, candidates[i])
			CSBT(i + 1)
			total -= setList[len(setList)-1]
			setList = setList[:len(setList)-1]
			CSBT(i + 1)
		} else {
			return
		}
	}

	CSBT(0)
	return combinationSol
}
