package main

import "fmt"

// isValidSudoku checks if a 9x9 board is a valid Sudoku puzzle.
// this is not my solution but is optimized my code with bing AI
func isValidSudoku2(board [][]byte) bool {
	rowMask := make([]int, 9)
	colMask := make([]int, 9)
	boxMask := make([]int, 9)
	// Iterate over the board and check the validity of each digit.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Get the current digit as an integer.
			digit := int(board[i][j] - '0')
			// If the digit is not between 1 and 9, skip it.
			if digit < 1 || digit > 9 {
				continue
			}
			// Create a bit mask for the digit, such as 1 << 2 for 2.
			mask := 1 << digit
			// Get the index of the sub-box that contains the digit, such as 3 * (i / 3) + (j / 3) for (i, j).
			boxIndex := 3*(i/3) + (j / 3)
			// Check if the digit is already present in the row, column, or sub-box using the bit mask and the bitwise AND operation.
			if rowMask[i]&mask != 0 || colMask[j]&mask != 0 || boxMask[boxIndex]&mask != 0 {
				// If the digit is repeated, return false.
				return false
			}
			// If the digit is not repeated, update the bit masks for the row, column, and sub-box using the bitwise OR operation.
			rowMask[i] |= mask
			colMask[j] |= mask
			boxMask[boxIndex] |= mask
		}
	}
	// Return true if no false condition is met.
	return true
}

// https://leetcode.com/problems/valid-sudoku/
func main() {
	fmt.Println(isValidSudoku2([][]byte{
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
