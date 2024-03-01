package solutions

func areDuplicates(nums []byte) bool {
	var seen []byte
	for _, num := range nums {
		if num == '.' {
			continue
		}

		for _, s := range seen {
			if num == s {
				return true
			}
		}
		seen = append(seen, num)
	}
	return false
}

func areRowsValid(board [][]byte) bool {
	for _, row := range board {
		if areDuplicates(row) {
			return false
		}
	}
	return true
}

func areColumnsValid(board [][]byte) bool {
	// We know that there are always 9 rows and 9 cols
	for colIndex := 0; colIndex < 9; colIndex++ {
		col := []byte{
			board[0][colIndex],
			board[1][colIndex],
			board[2][colIndex],
			board[3][colIndex],
			board[4][colIndex],
			board[5][colIndex],
			board[6][colIndex],
			board[7][colIndex],
			board[8][colIndex],
		}
		if areDuplicates(col) {
			return false
		}
	}
	return true
}

func areSquaresValid(board [][]byte) bool {
	// We know that there are always 9 squares
	for rowIndex := 0; rowIndex < 9; rowIndex += 3 {
		for columnIndex := 0; columnIndex < 9; columnIndex += 3 {
			square := []byte{
				board[rowIndex][columnIndex],
				board[rowIndex][columnIndex+1],
				board[rowIndex][columnIndex+2],
				board[rowIndex+1][columnIndex],
				board[rowIndex+1][columnIndex+1],
				board[rowIndex+1][columnIndex+2],
				board[rowIndex+2][columnIndex],
				board[rowIndex+2][columnIndex+1],
				board[rowIndex+2][columnIndex+2],
			}
			if areDuplicates(square) {
				return false
			}
		}
	}
	return true
}

// isValidSudoku is the thirty-sixth LeetCode problem:
// - https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	rowsValid := areRowsValid(board)
	columnsValid := areColumnsValid(board)
	squaresValid := areSquaresValid(board)

	return rowsValid && columnsValid && squaresValid
}

var IsValidSudoku = isValidSudoku /* For testing */
