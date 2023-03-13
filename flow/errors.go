package flow

import (
	"fmt"
)

func Errors() {
	// error是一种内置的类型，返回值是string
	// var e error
	// fmt.Println(e.Error())
	n, error := echo(10)
	if error != nil {
		fmt.Println("error: " + error.Error())
	} else {
		fmt.Println(n)
	}

	// 1. defer
	// 延迟函数调用，defer后会跟一个函数，但是不会立即执行
	// 而是等包含它的程序返回时（return，goroutine panic等）
	// 才会执行
	fmt.Println(triple(2))
	fmt.Println(example())
	onetwothree()

	// 2. Panic
	// panic(error) 直接抛出error信息及位置
	// defer 要做panic之上才执行
	// defer func() {
	// 	fmt.Println("panic and defer")
	// }()
	// _, err := os.Create("/tmp/file")
	// if err != nil {
	// 	panic(err)
	// }

	// 3. recover, 需要在defer函数中执行
	// pass the panic call to recover
	// 指的是上层函数恢复，省略本层函数
	G()
	f()
	fmt.Println("Returned normally from f.")

}

func echo(x int) (int, error) {
	return x, nil
}

// 自定义error类型
// 先定义一个type
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()

	return x + x
}

// defer在return后进行
// 但是最后会return result而不是return double(x)
// 在函数最后一刻执行
func triple(x int) (result int) {
	defer func() {
		result += x
	}()

	return double(x)
}

func example() int {
	for i := 0; i < 9; i++ {
		if i == 6 {
			return i
		}
		// defer 将被执行6次，每个循环一次
		defer func() {
			fmt.Println("end")
		}()
	}
	return 10
}

// defer按照倒叙执行
func onetwothree() {
	defer func() {
		fmt.Println("first")
	}()

	defer func() {
		fmt.Println("second")
	}()

	fmt.Println("done")
}

func G() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
