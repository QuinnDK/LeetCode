package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"

*/
func main() {
	num1 := getInput()
	num2 := getInput()
	res := multiply(num1, num2)
	fmt.Printf(res)

}
func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}
func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	pos := 0
	res := make([]byte, l1+l2+1) // 可能向前进位，多保留一位
	for i := range res {         // 初始化res
		res[i] = '0'
	}
	for i := l1 - 1; i >= 0; i-- { // 需要用一个数乘以另一个数的每一位，因此两层遍历
		for j := l2 - 1; j >= 0; j-- {
			temp := (num1[i] - '0') * (num2[j] - '0')
			pos = i + j + 2
			res[pos-1] += (temp + res[pos] - '0') / 10 // 进位
			res[pos] = (temp+res[pos]-'0')%10 + '0'
		}
	}
	ret := strings.TrimLeft(string(res), "0") // 移除左侧多余的0
	if ret == "" {
		return "0"
	}
	return ret
}
