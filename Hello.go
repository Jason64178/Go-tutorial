// 00_hello.go

package main // 声明 main 包

import "fmt" // 导入内置 fmt 包

func variables() {
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

	// 声明单个变量
	var a = "initial"
	var d = true

	fmt.Println(a, d)

	// 声明多个变量
	var b, c int = 1, 2
	var (
		b1, c1 int
		b2, c2 = 3, 4
	)
	fmt.Println(b, c, b1, c1, b2, c2)

	// 短变量声明
	f := "short"
	g, h := 5, "alwaysday1"
	// 声明多个变量时，至少有一个新变量
	var i int
	// i != 100 // 报错，no new variables on left side of :=
	i, j := 100, 101 //有新值，不报错
	fmt.Println(f, g, h, i, j)

}

func main() { // main函数，程序执行入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!

	variables()
}
