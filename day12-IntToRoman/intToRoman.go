package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := intToRoman(n)
	fmt.Println(result)
}

func intToRoman(num int) string {
	m := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD",
		100: "C", 90: "XC", 50: "L", 40: "XL",
		10: "X", 9: "IX", 5: "V", 4: "IV",
		1: "I"}
	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for _, v := range arr {
		t := num / v
		num = num - t*v

		for t > 0 { // 可以整除就拼接
			res += m[v]
			t--
		}
		if num == 0 {
			break
		}
	}
	return res
}
