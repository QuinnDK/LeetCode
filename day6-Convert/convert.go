package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numRows int
	var s string

	//fmt.Scanf("%s",&s)     //Scanf 只会获得一次输入就退出
	//fmt.Println(s)
	//fmt.Scanf("%d",&numRows)
	//fmt.Println(numRows)

	fmt.Print("Get input by Reader:")
	s = getInput()
	fmt.Println(s)
	fmt.Scanf("%d", &numRows)
	fmt.Println(numRows)

	//fmt.Print("Get input by Scanner:")
	//str2 := getInputByScanner()
	//fmt.Println(str2)

	//fmt.Print("Get input by fmt:")
	//str3 := getInputByFmt()
	//fmt.Println(str3)

	//最原始的测试
	//s:="LEETCODEISHIRING"
	//numRows:=3
	result := convert(s, numRows)
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

//以下两种方法也可以实现多次输入
func getInputByScanner() string {
	var str string
	//使用os.Stdin开启输入流
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		str = in.Text()
	} else {
		str = "Find input error"
	}
	return str
}

func getInputByFmt() string {
	//定义变量
	var str string
	//Scanf函数读取输入到变量中 两个返回值
	//分别为读取到的长度 失败信息
	length, err := fmt.Scanf("%s", &str) //注意使用%s读取输入字符串只能读取到空白符之前
	if err != nil {
		return err.Error()
	}
	fmt.Println("Read length is ", length)
	return str
}

func convert(s string, numRows int) string {
	// 不满足，提前返回
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	// 保存最终结果
	var res string
	// 保存每次的结果
	var tmp = make([]string, numRows)
	// 初始位置
	curPos := 0
	// 用来标记是否该转向了
	shouldTurn := -1
	// 遍历每个元素
	for _, val := range s {
		// 添加进tmp里面，位置为curPos
		tmp[curPos] += string(val)
		// 如果走到头或者尾，就该转向了
		// 因为就在numRows的长度里面左右震荡
		if curPos == 0 || curPos == numRows-1 {
			// 转向
			shouldTurn = -shouldTurn
		}
		// curPos 用来标记tmp里面我们该去哪个位置
		curPos += shouldTurn
	}
	// 这时tmp里面已经保存了数据了，我们只需要转换一下输出格式
	for _, val := range tmp {
		res += val
	}
	// 最后的输出
	return res
}
