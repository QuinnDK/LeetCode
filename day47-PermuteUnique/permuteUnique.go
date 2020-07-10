package main

import "fmt"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

*/
func main() {
	nums := []int{1, 1, 3}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res, 0)
	return res
}

func dfs(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		*res = append(*res, dump(nums))
	}

	m := map[int]int{}
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		m[nums[i]] = 1
	}
}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
