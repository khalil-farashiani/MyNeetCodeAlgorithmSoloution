package main

import "fmt"

// first implementation for sudoku that seems not very good time complexity (2 nested loop)
func isValidSudoku(board [][]byte) bool {
	var firstMapBoxRepeated = make(map[byte]bool, 9)
	var secondMapBoxRepeated = make(map[byte]bool, 9)
	var thirdMapBoxRepeated = make(map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		var mapRowRepeated = make(map[byte]bool, 9)
		var mapColRepeated = make(map[byte]bool, 9)
		if i%3 == 0 {
			firstMapBoxRepeated = make(map[byte]bool, 9)
			secondMapBoxRepeated = make(map[byte]bool, 9)
			thirdMapBoxRepeated = make(map[byte]bool, 9)
		}
		for j := 0; j < 9; j++ {
			if mapRowRepeated[board[i][j]] == true && board[i][j] != 46 {
				return false
			}
			mapRowRepeated[board[i][j]] = true
			if mapColRepeated[board[j][i]] == true && board[j][i] != 46 {
				return false
			}
			mapColRepeated[board[j][i]] = true

			if j < 3 && firstMapBoxRepeated[board[i][j]] == true && board[i][j] != 46 {
				return false
			}
			if j > 2 && j < 6 && secondMapBoxRepeated[board[i][j]] == true && board[i][j] != 46 {
				return false
			}
			if j > 5 && thirdMapBoxRepeated[board[i][j]] == true && board[i][j] != 46 {
				return false
			}
			if j < 3 {
				firstMapBoxRepeated[board[i][j]] = true
			}
			if j > 2 && j < 6 {
				secondMapBoxRepeated[board[i][j]] = true
			}
			if j > 5 {
				thirdMapBoxRepeated[board[i][j]] = true
			}
		}
	}
	return true
}

// https://leetcode.com/problems/valid-sudoku/
func main() {
	fmt.Println(isValidSudoku([][]byte{
		//0    1    2    3    4    5    6    7    8
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, // 0
		{'3', '.', '.', '.', '4', '6', '5', '.', '.'}, // 1
		{'.', '.', '5', '.', '2', '.', '.', '.', '.'}, // 2
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, // 3
		{'.', '3', '.', '.', '.', '.', '.', '.', '1'}, // 4
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, // 5
		{'.', '.', '.', '.', '.', '.', '6', '.', '.'}, // 6
		{'.', '.', '.', '.', '.', '5', '.', '6', '.'}, // 7
		{'.', '.', '.', '9', '.', '.', '.', '.', '.'}, // 8
	},
	))
}
