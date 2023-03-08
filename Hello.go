// 00_hello.go

package main // 声明 main 包

import "fmt" // 导入内置 fmt 包

func main() { // main函数，程序执行入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!

	// 没有初始值，会赋默认零值
	var v1 int
	var v2 string
	var v3 bool
	var v4 [10]int // 数组
	var v5 []int   // 数组切片
	var v6 struct {
		e int
		f string
	}

	fmt.Println(v1, v2, v3, v4, v5, v6)
}
