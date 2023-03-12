package variable

import "fmt"

func Variables() {
	// 1. 使用var声明变量
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

	// 2. 短变量声明
	f := "short"
	g, h := 5, "alwaysday1"
	// 声明多个变量时，至少有一个新变量
	var i int
	// i := 100 // 报错，no new variables on left side of :=
	i, j := 100, 101 //有新值，不报错
	fmt.Println(f, g, h, i, j)

	// 3. 使用new声明指针
	k := 6
	l := &k
	fmt.Println(*l) // 输出6
	*l = 7
	fmt.Println(k) // 输出7
	var p = new(int)
	fmt.Println(*p) // 输出0，默认值
	*p = 8
	fmt.Println(*p) // 输出8

	// 4. 使用=赋值
	m, n := 1, 2
	m, n = n, m
	fmt.Println(m, n) // 2, 1

	// 5. 常量
	const Pi float64 = 3.14159265358979323846
	const nConst = 5000000000
	const dConst = 3e20 / nConst
	const zero = 0.0
	fmt.Println(Pi, nConst, dConst, zero)
	const ac, bc, cc = 3, 4, "foo"
	const (
		size int64 = 1024
		eof  int   = -1
	)
	fmt.Println(ac, bc, cc, size, eof)

	// 6.iota 从0开始，逐项加1
	const (
		a0 = iota // 0
		a1 = iota // 1
		a2 = iota // 2
	)
	fmt.Println(a0, a1, a2)

	const (
		d0 = iota // 0
		d1        // 1
		d2        // 2
	)
	fmt.Println(d0, d1, d2)

	// 重置为0
	const x = iota // 0
	fmt.Println(x)

	// 枚举
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

}
