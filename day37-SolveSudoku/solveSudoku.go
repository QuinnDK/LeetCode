package main

/*
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。

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
	solveSudoku(nums)

}
func solveSudoku(board [][]byte) {
	backtrace(board, 0, 0)
}

func backtrace(board [][]byte, row int, col int) bool {
	if row == 9 { //遍历完成，找到答案
		return true
	}
	if col == 9 { //换行
		return backtrace(board, row+1, 0)
	}
	if board[row][col] != '.' { //跳过已填有数字的格子
		return backtrace(board, row, col+1)
	}
	for i := '1'; i <= '9'; i++ { //遍历1--9
		if isOK(board, row, col, byte(i)) { //判断行列单元格是否存在当前数字
			board[row][col] = byte(i)         //放置数字
			if backtrace(board, row, col+1) { //放置下一个格子
				return true
			}
			board[row][col] = '.' //若当前放置的格子无法使下一个格子放置成功，则清除当前操作
		}
	}
	return false
}

func isOK(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[row][i] == c {
			return false
		}
		// 判断列是否存在重复
		if board[i][col] == c {
			return false
		}
		// 判断单元格是否存在重复
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
			return false
		}
	}
	return true
}
