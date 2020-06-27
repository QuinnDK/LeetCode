package main

import "fmt"

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

*/
func main() {
	nums := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	result := isValidSudoku(nums)
	fmt.Println(result)
}
func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index_box := (i/3)*3 + j/3
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if sbox[index_box][num] == 1 {
					return false
				} else {
					sbox[index_box][num] = 1
				}
			}
		}
	}
	return true
}
