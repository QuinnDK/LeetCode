package main

/*
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := getInput()
	result := romanToInt(s)
	fmt.Println(result)

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

func romanToInt(s string) int {
	romans := [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	var start int
	var v int
	for i := len(romans) - 1; i >= 0; {
		charLen := len(romans[i])
		end := start + charLen
		if end > len(s) || s[start:end] != romans[i] {
			i--
			continue
		}

		start += charLen
		v += nums[i]
	}
	return v
}
