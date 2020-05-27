package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
*/

func main() {
	fmt.Print("Get input by Reader:")
	s := getInput()
	re, _ := strconv.Atoi(s)
	Result := reverse(re)
	fmt.Println(Result)

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

func reverse(x int) int {
	var y int = 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > 2147483647 || y < -2147483648 {
		return 0
	}

	return y
}
