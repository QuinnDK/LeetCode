package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	fmt.Println(permute2(2))
}
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
func permute2(num int) int {
	nums := make([]int, num-1)
	for i := 0; i < num; i++ {
		nums[i] = i
	}
	if len(nums) == 1 {
		return 0
	}

	res := [][]int{}
	result := 0
	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for k, s := range sub {
			res = append(res, append(s, num))
			result = k
		}
	}
	return result
}

//回溯
func permute1(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res, 0)
	return res
}

func dfs(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		*res = append(*res, dump(nums))
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
