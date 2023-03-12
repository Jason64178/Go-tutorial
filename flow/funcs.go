package flow

import "fmt"

func Funcs() {
	// 关键词func，函数名，参数列表，返回列表，函数体
	fmt.Printf("%T\n", add) // func(int, int) int
	fmt.Printf("%T\n", sub) // func(int, int) int

	fmt.Println(funcSum(1, 2, 3))
	s := []int{1, 2, 3, 4}
	fmt.Println(funcSum(s...))
	fmt.Println(funcSum1(s))
	fmt.Println(swap(1, 2)) // 2 1
	// 值传递，而非引用传递，go里面没有引用传递
	// 可传指针
	x, _ := swap(1, 2)
	fmt.Println(x) // 2

	fmt.Println(funcSum2(add, 1, 2))
	f := warp("add")
	fmt.Println(f(2, 4))

	// 直接调用
	// 类似于lambda
	fmt.Println(func(a, b int) int { return a + b }(4, 5))

}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return z
}

// 支持不定参数
// 传slice时，需要...将slice展开
func funcSum(args ...int) (ret int) {
	for _, v := range args {
		ret += v
	}
	// 直接写return即可
	return
}

func funcSum1(l []int) (ret int) {
	for _, v := range l {
		ret += v
	}
	// 直接写return即可
	return
}

// 可以有多个返回值
func swap(x, y int) (int, int) {
	return y, x
}

// 匿名函数
// 作为参数
func funcSum2(f func(int, int) int, x, y int) int {
	return f(x, y)
}

// 作为返回值
func warp(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a, b int) int {
			return a - b

		}
	default:
		return nil
	}
}
