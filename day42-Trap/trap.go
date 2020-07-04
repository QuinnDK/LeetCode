package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
func main() {
	heig := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(heig)
	fmt.Println(res)
}
func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//设置左边最高柱子
				leftMax = height[left]
			} else {
				//右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				//设置右边最高柱子
				rightMax = height[right]
			} else {
				//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
