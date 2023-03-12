package variable

import "fmt"

func ArraySlice() {
	// 1. 数组，长度固定，类型相同
	// 较少使用
	var a [3]int
	// 输出第一个元素
	fmt.Println(a[0])
	// 输出数组长度
	fmt.Println(len(a))

	// 数组字面量初始化
	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}
	fmt.Println(b, c[2])

	// ...表示长度
	d := [...]int{1, 2, 3, 4, 5}
	// %T the type of
	fmt.Printf("%T\n", d)

	e := [4]int{5, 2: 10}
	f := [...]int{2, 4: 6}
	fmt.Println(e, f)

	// 多维数组，仅第一个用...
	var g [4][2]int
	h := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	i := [4][2]int{1: {0: 20}, 3: {1: 40}}
	fmt.Println(g, h, i)

	// 数组长度一致可比较
	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{1, 3}

	fmt.Println(a1 == a2, a1 == a3, a2 == a3)

	// 数组遍历
	for i, v := range e {
		fmt.Println(i, v)
	}

	// 数组是值类型
	x := [2]int{10, 20}
	y := x

	fmt.Printf("x: %p, %v\n", &x, x) // x: 0xc00012e020, [10 20]
	fmt.Printf("y: %p, %v\n", &y, y) // y: 0xc00012e030, [10 20]
	test(x)
	fmt.Printf("x: %p, %v\n", &x, x) // x: 0xc00012e020, [10 20]

	// 2. 切片slice
	// 指针：指向slice可以访问到第一个元素（基于数组创建）
	// 长度：slice中元素个数（现有的）
	// 容量：slice起始元素到底层数组最后一个元素间的元素个数（最大的）

	// 2.1 基于数组创建
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := array[3:6] // [4, 5, 6], len = 3, cap = 6
	s2 := array[:5]  // [1, 2, 3, 4, 5], len =5, cap = 9
	s3 := array[4:]  // [5, 6, 7, 8], len = 4, cap = 5
	s4 := array[:]

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	// 2.2 基于make创建
	s5 := make([]int, 10)     // len 10, cap 10
	s6 := make([]int, 10, 15) // len 10, cap 15
	fmt.Printf("s5: %v, len: %d, cap: %d\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %v, len: %d, cap: %d\n", s6, len(s6), cap(s6))

	// 2.3 直接创建
	s7 := []int{12, 13} // len 2, cap 2
	fmt.Printf("s7: %v, len: %d, cap: %d\n", s7, len(s7), cap(s7))

	for i, n := range s1 {
		fmt.Println(i, n)
	}

	// 比较，不能比较两个slice
	// 但是可与nil比
	var s []int
	fmt.Println(len(s) == 0, s == nil) // true true
	s = nil
	fmt.Println(len(s) == 0, s == nil) // true true
	s = []int(nil)
	fmt.Println(len(s) == 0, s == nil) // true true
	s = []int{}
	fmt.Println(len(s) == 0, s == nil) // true false

	// append, 输出新的slice
	s4 = append(s4, 10)
	fmt.Printf("s4: %v\n", s4)
	s8 := append(s4, 10, 11)
	fmt.Printf("s8: %v\n", s8)

	// append s6 to s7
	s7 = append(s7, s6...)
	fmt.Printf("s7: %v\n", s7) // [12, 13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

	// copy：不能超过cap
	s9 := []int{1, 2, 3, 4, 5}
	s10 := []int{5, 4, 3}
	s11 := []int{6}

	copy(s9, s10)
	fmt.Printf("s9: %v\n", s9) // s8: [5 4 3 4 5]
	copy(s11, s10)
	fmt.Printf("s11: %v\n", s11) // s10: [5]

	// 引用为地址引用
	modify(s9)
	fmt.Println("main: ", s9) // main:  [30 4 3]

}

func test(a [2]int) {
	a[0] = 100
	fmt.Printf("a: %p, %v\n", &a, a)
}

func modify(a []int) {
	a[0] = 30
	fmt.Println("modify: ", a) // modify:  [30 4 3]
}
