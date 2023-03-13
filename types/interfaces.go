package types

import "fmt"

func InterfaceMain() {
	// 1. 接口
	// 接口是一种抽象类型，是多个方法生命的集合
	// 在Go中，只有目标实现了接口要求的所有方法，就实现了这个接口
	var c Duck = &Cat{}
	c.Eat()

	// 将实例对象赋值给接口
	// 一定要传指针
	var d Duck = &Dog{}
	d.Eat()

	s := []Duck{
		&Cat{},
		&Dog{},
	}
	for _, n := range s {
		n.Eat()
	}
	// 输出
	// cat eat
	// dog eat

	var e Duck = &Cat{}
	// duck 可不传指针，但cat必须传
	var f Duck = e
	f.Eat()
	// 已经初始化的接口变量 c1 直接赋值给另一个接口变量 c2，要求 c2 的方法集是 c1 的方法集的子集。
	var c1 Duck1 = &Dog{}
	var c2 Duck = c1
	c2.Eat()

	// 2. 空接口
	// 没有方法，用interface {} 表示，所有类型都实现了空接口
	s1 := "Hello World"
	i := 50
	strt := struct {
		name string
	}{
		name: "xiasen",
	}
	test(s1)
	test(i)
	test(strt)

	// 3. 类型断言
	// x.(T) x是接口类型，T是断言类型
	// 如果有两个接收值，那么断言不会在失败时崩溃，而是会多返回一个布尔值，一般命名为 ok，来表示断言是否成功。
	var n1 interface{} = "hello"
	var n2 interface{} = 55
	assertFlag(n1)
	assertFlag(n2)
	// 4. 类型查询
	// x.(type)

	// 类型查询
	SearchType(50)         // Int: 50
	SearchType("zhangsan") // String: zhangsan
	SearchType(c)          // dog eat
	SearchType(50.1)       // Unknown type

}

// 定义接口，包含Eat方法
type Duck interface {
	Eat()
}

// 定义Cat结构体，并实现Eat方法
type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("cat eat")
}

// 定义Dog结构体，并实现Eat方法
type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("dog eat")
}

type Duck1 interface {
	Eat()
	Walk()
}

func (d *Dog) Walk() {
	fmt.Println("dog walk")
}

func test(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assertFlag(i interface{}) {
	if s, ok := i.(int); ok {
		fmt.Println(s)
	} else {
		fmt.Println("not an int")
	}
}

func SearchType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("String: %s\n", i.(string))
	case int:
		fmt.Printf("Int: %d\n", i.(int))
	case Duck:
		v.Eat()
	default:
		fmt.Printf("Unknown type\n")
	}
}
