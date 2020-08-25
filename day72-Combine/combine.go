package main

import "fmt"

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
示例:
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{} // 处理边界
	}
	var res [][]int        // 定义结果集合
	var elem []int         // 定义当前组合
	nums := make([]int, n) // 构造[1,..,n]
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	combineRec(nums, 0, k, elem, &res)
	return res
}
func combineRec(nums []int, start, k int, elem []int, result *[][]int) {
	if len(elem) == k { // 当前组合元素个数等于k时
		tmp := make([]int, k)
		copy(tmp, elem) // 将当前组合加入到结果集合
		*result = append(*result, tmp)
	}
	// 从下标start开始遍历数组
	for i := start; i < len(nums); i++ {
		// 取出一个数字加入到当前组合
		elem = append(elem, nums[i])
		// 递归调用自己，不能重复使用数字，下一轮递归从i+1开始
		combineRec(nums, i+1, k, elem, result)
		// 递归结束后把当前加入的数字移除，回溯
		elem = elem[:len(elem)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
