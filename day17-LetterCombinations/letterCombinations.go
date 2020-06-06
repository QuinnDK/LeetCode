package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。


示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

*/

func main() {
	s := getInput()
	result := letterCombinations(s)
	fmt.Print(result)

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

var table []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}

	for i := 0; i < len(digits); i++ {
		t := table[digits[i]-'0']
		temp := []string{}
		for j := 0; j < len(t); j++ {
			for z := 0; z < len(result); z++ {
				temp = append(temp, result[z]+string([]byte{t[j]}))
			}
		}
		result = temp
	}

	return result
}

/**
  这题使用广度优先搜索比较简单

*/

//
//var num1string map[byte][]string = map[byte][]string{
//	'2':{"a","b","c"},
//	'3':{"d","e","f"},
//	'4':{"g","h","i"},
//	'5':{"j","k","l"},
//	'6':{"m","n","o"},
//	'7':{"p","q","r","s"},
//	'8':{"t","u","v"},
//	'9':{"w","x","y","z"},
//
//}
//
//func letterCombinations(digits string)[]string  {
//	res:=[]string{}
//	if digits==""{
//		return res
//	}
//	res=append(res,"")
//	for i := 0; i < len(digits); i++ {
//		nums := num1string[digits[i]]
//		pre := res
//		res := []string{}
//		for _, num := range nums {
//			for _, v := range pre {
//				res = append(res, v+num)
//			}
//		}
//	}
//	return res
//}
