package main

import "fmt"

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/
func sortColors(nums []int) {
	var cnts [3]int
	for _, v := range nums {
		cnts[v]++
	}
	j := 0
	for i, cnt := range cnts {
		for ; j < len(nums); j++ {
			nums[j] = i
			if cnt == 0 {
				break
			}
			cnt--
		}
	}
}
func sortColors1(nums []int) {
	if len(nums) == 0 {
		return
	}
	//1 初始化三个指针
	p0 := 0
	p2 := len(nums) - 1
	cur := 0
	//2 cur开始遍历
	for cur <= p2 {
		if nums[cur] == 0 { //3 cur位置与p0位置交换
			nums[cur] = nums[p0]
			nums[p0] = 0
			p0++
		} else if nums[cur] == 1 { //4 cur加1
			cur++
		} else if nums[cur] == 2 { //5 cur与p2位置交换
			nums[cur] = nums[p2]
			nums[p2] = 2
			p2--
		}
		//6 调整cur指针位置
		if cur < p0 {
			cur = p0
		}
	}
}
func main() {
	num := []int{2, 0, 2, 1, 1, 0}
	sortColors1(num)
	fmt.Println(num)
}
