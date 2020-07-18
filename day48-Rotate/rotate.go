package main

import "fmt"

/*
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。
说明：
你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
示例 1:
给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	var tmp, i, j int
	// mirror by diagonal
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	// reverse the element order in each row
	for i = 0; i < n; i++ {
		for j = 0; j < n-j-1; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = tmp
		}
	}
}
func main() {
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(num)
	fmt.Println(num)
}
