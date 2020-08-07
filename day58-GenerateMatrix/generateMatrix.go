package main

import "fmt"

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/
func main() {
	fmt.Println(generateMatrix(3))
}
func generateMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	size := n * n
	a, b := 0, n-1 //左右边界
	c, d := 0, n-1 //上下边界

	idx := 1
	for idx <= size {
		//向右走
		for i := a; i <= b; i++ {
			result[c][i] = idx
			idx++
		}
		c++ //上边界收缩
		//向下走
		for i := c; i <= d; i++ {
			result[i][b] = idx
			idx++
		}
		b-- //右边界收缩
		//向左走
		for i := b; i >= a; i-- {
			result[d][i] = idx
			idx++
		}
		d-- //。。。
		//向上走
		for i := d; i >= c; i-- {
			result[i][a] = idx
			idx++
		}
		a++ //。。。
	}
	return result
}
