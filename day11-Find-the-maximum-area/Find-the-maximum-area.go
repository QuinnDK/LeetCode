package main

import "fmt"

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

*/

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(s)
	result1 := maxArea(a)
	fmt.Println(result)
	fmt.Println(result1)
}

func maxArea(height []int) int {
	max := 0
	p1, p2 := 0, len(height)-1
	for p1 <= p2 {
		if height[p1] <= height[p2] {
			tempMax := height[p1] * (p2 - p1)
			max = Max(tempMax, max)
			p1++
		} else {
			tempMax := height[p2] * (p2 - p1)
			max = Max(tempMax, max)
			p2--
		}
	}
	return max
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
