package main

import "fmt"

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
*/

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	row := len(matrix)
	col := len(matrix[0])
	up := 0            //上界
	down := row - 1    //下界
	left := 0          //左界
	right := col - 1   // 右界
	total := row * col //总数
	for {
		// 从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		if len(res) == total {
			break
		}
		up++ // 上边界+1 向下移动一位
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		if len(res) == total {
			break
		}
		right--
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		if len(res) == total {
			break
		}
		down--
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		if len(res) == total {
			break
		}
		left++
	}
	return res
}

func main() {
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(num))
}
