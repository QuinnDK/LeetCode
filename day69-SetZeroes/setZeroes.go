package main

import "fmt"

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。
请使用原地算法。
示例 1:
输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
*/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rows, cols := make(map[int]bool), make(map[int]bool)
	for i := 0; i < m; i++ {
		if rows[i] { // 该行被标记置为0
			continue
		}
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i], cols[j] = true, true // 将所在行、列标记
			}
		}
	}
	for r, _ := range rows { // 将标记的每一行置0
		for j := 0; j < n; j++ {
			matrix[r][j] = 0
		}
	}
	for col, _ := range cols { // 将标记的每一列置位0
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
	}
	return
}

func main() {
	num := [][]int{{1, 2, 1, 1}, {1, 0, 2, 1}, {1, 2, 3, 1}}
	setZeroes(num)
	fmt.Println(num)
}
