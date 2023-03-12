package flow

import "fmt"

func Flow() {
	// 1. if-else
	// 不需（），但需要{}
	// if 后可添加变量初始化语句，使用；分隔
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// 2. switch
	// 可添加变量初始化语句，使用；分割
	// 使用fallthrough，会紧跟着执行下一个，不会判断case语句
	// 支持default
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fmt.Print(" hhhhh ")
		fallthrough
	case 3:
		fmt.Println("three") // three
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("None")
	}

	// 3. for
	// 无需（），但需{}
	// 可添加变量初始化语句，使用；分割
	// 支持break，continue

	i = 1
	// 只有条件
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 有变量初始化和条件
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 遍历数组
	a := [...]int{10, 20, 30, 40}
	for i := range a {
		fmt.Println(i)
	}
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 遍历切片
	s := []string{"a", "b", "c"}
	for i := range s {
		fmt.Println(i)
	}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// 遍历字典
	m := map[string]int{"a": 10, "b": 20, "c": 30}
	for k := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
