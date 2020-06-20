package main

import "fmt"

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2


示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3

*/
func main() {
	dividend := 10
	divisor := 3
	result := divide(dividend, divisor)
	fmt.Println(result)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	maxInt := 1 << 31

	if dividend == maxInt-1 && divisor == -1 {
		return maxInt - 1
	}
	isMinus := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)

	var (
		ds = divisor
		dd = dividend
	)

	if ds < 0 {
		ds = -ds
	}

	if dd < 0 {
		dd = -dd
	}

	var result int
	for i := 31; i >= 0; i-- {
		it := uint(i)
		if dd>>it >= ds {
			result += 1 << it
			dd -= ds << it
		}
	}

	if isMinus {
		result = -result
	}

	if result > maxInt-1 {
		result = maxInt - 1
	}
	return result
}
