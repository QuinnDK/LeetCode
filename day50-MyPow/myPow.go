package main

import "fmt"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
*/
func main() {
	ans := myPow(2.00000, 10)
	fmt.Println(ans)
}
func myPow(x float64, n int) float64 {
	if n >= 0 {
		ans := 1.00
		xx := x
		for n > 0 {
			if n%2 == 1 {
				ans *= xx
			}
			xx *= xx
			_ = n >> 1
		}
		return ans
	}
	return 1 / myPow(x, -n)
}
