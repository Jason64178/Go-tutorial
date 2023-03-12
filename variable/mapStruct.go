package variable

import "fmt"

func MapStruct() {
	// 1. map, map[k]V
	// 字面量方式创建
	var m = map[string]int{"a": 1, "b": 2}
	fmt.Println(m)

	// 使用make创建
	m1 := make(map[string]int)
	fmt.Println(m1)

	// 可初始化字典长度，提升执行效率
	m2 := make(map[string]int, 10)
	fmt.Println(m2)

	// 零值是nil，对nil赋值报错
	var m3 map[string]int
	fmt.Println(m3 == nil, len(m3) == 0)

	// nil 赋值报错
	// m3["a"] = 1
	// fmt.Println(m3)    // panic: assignment to entry in nil map

	// 赋值
	m1["a"] = 1
	m1["d"] = 1
	fmt.Println(m1)
	fmt.Println(m1["a"], m1["d"])
	fmt.Println(m1["e"]) // 返回0

	// 删除
	delete(m1, "a")
	delete(m1, "f") // 不报错
	fmt.Println(m1)

	// 长度
	fmt.Println(len(m1))

	// 判断是否存在
	if value, ok := m1["d"]; ok {
		fmt.Println(value)
	}

	// 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 引用类型

	// 2. 结构体
	// 使用type来定义
	type user struct {
		name string
		age  int
	}

	u1 := user{"xiasen", 19}
	fmt.Println(u1)

	u2 := user{
		name: "xiasen",
		age:  19,
	}
	fmt.Println(u2)

	fmt.Println(u2.name, u2.age)
	u2.name = "Jason"
	fmt.Println(u2.name, u2.age)

	// 可比较
	fmt.Println(u1 == u2)

	// 嵌套结构体
	type admin struct {
		u       user
		isAdmin bool
	}

	a := admin{
		u:       u1,
		isAdmin: true,
	}
	fmt.Println(a)
	fmt.Println(a.u.name)
	fmt.Println(a.u.age)
	fmt.Println(a.isAdmin)

	// 匿名成员，不指定名称，只指定类型
	type admin1 struct {
		user
		isAdmin bool
	}

	a2 := admin1{
		user:    u2,
		isAdmin: true,
	}
	a2.age = 20
	a2.isAdmin = false

	fmt.Println(a2)
	fmt.Println(a2.name)
	fmt.Println(a2.age)
	fmt.Println(a2.isAdmin)
}
