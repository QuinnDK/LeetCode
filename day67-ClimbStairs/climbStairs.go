package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
*/
func climbStairs(n int) int {
	var result = []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}
	return result[n]
}
func main() {
	fmt.Println(climbStairs(6))
}
