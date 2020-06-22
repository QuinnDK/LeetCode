package main

import (
	"fmt"
	"sort"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/

func main() {
	s := []int{158476531}
	nextPermutation(s)
}

func nextPermutation(nums []int) {
	// 第一步从右往左查找最大的升序队列
	i := len(nums) - 2 // 倒数第二个数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		// 从右往左一直是升序
		sort.Ints(nums)
	} else {
		j := len(nums) - 1
		// 从右往左找第一个比i索引值大的数字索引坐标
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		//交换他两位置
		nums[i], nums[j] = nums[j], nums[i]
		// 重新排序
		sort.Ints(nums[i+1:])

	}
	fmt.Println(nums)
}
