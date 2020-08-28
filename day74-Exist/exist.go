package main

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
*/
func exist(board [][]byte, word string) bool {
	n := len(board)    // 矩阵长度
	m := len(board[0]) // 矩阵宽度

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrace79(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrace79(board [][]byte, word string, i, j, k int) bool {
	// 找到了，不用继续往下找了
	if k == len(word) {
		return true
	}
	// 如果越界了，就说明没有啦
	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	tmp := board[i][j]
	// 重置它是为了回溯往回找的时候避免重复使用，干脆，如果找到它是对的，就直接把它置为 空
	// 等结束了之后再重置回来
	board[i][j] = byte(' ')

	// 开始上、下、左、右探测
	// 向左
	if backtrace79(board, word, i-1, j, k+1) {
		return true
	}
	// 向右
	if backtrace79(board, word, i+1, j, k+1) {
		return true
	}
	// 向上
	if backtrace79(board, word, i, j-1, k+1) {
		return true
	}
	// 向下
	if backtrace79(board, word, i, j+1, k+1) {
		return true
	}
	board[i][j] = tmp
	return false

}
