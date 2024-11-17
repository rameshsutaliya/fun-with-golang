package array_hashing

/*
func IsValidSudoku(board [][]byte) bool {
	// convert the byte board to 2-d int sudoku metrics
	// check if the row doesn't contain the duplicate.
	// check the column doesn't contain the duplicate number.
	// check the 3*3 square for 1-9 digits.
	sudoku := make([][]int, len(board))

	for i := 0; i < 9; i++ {
		if isValidRow(sudoku, i) && isValidColumn(sudoku, i) && isValidSquare(sudoku, i) {
			continue
		} else {
			return false
		}
	}

	return true
}

func isValidRow(sudoku [][]int, row int) bool {
	nums := make([]int, 9)
	for i := 0; i < 9; i++ {
		if sudoku[row][i] < 0 && sudoku[row][i] > 9 {
			return false
		} else if nums[sudoku[row][i]-1] != 0 {
			return false
		} else {
			nums[sudoku[row][i]-1] = sudoku[row][i]
		}
	}

	return true
}

func isValidColumn(sudoku [][]int, column int) bool {
	nums := make([]int, 9)
	for i := 0; i < 9; i++ {
		if sudoku[i][column] < 0 && sudoku[i][column] > 9 {
			return false
		} else if nums[sudoku[i][column]-1] != 0 {
			return false
		} else {
			nums[sudoku[i][column]-1] = sudoku[i][column]
		}
	}

	return true
}

func isValidSquare(sudoku [][]int, square int) bool {
	i, j := (square/3)*3, (square*3)%9
	nums := make([]int, 9)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if sudoku[i+row][j+col] < 0 && sudoku[i+row][j+col] > 9 {
				return false
			} else if nums[sudoku[i+row][j+col]-1] != 0 {
				return false
			} else {
				nums[sudoku[i+row][j+col]-1] = sudoku[i+row][j+col]
			}
		}
	}
	return true
}
*/
