package types

import "fmt"

func MethodMain() {
	// 1. 声明
	// 方法在定义时，会在func和各方法名之间增加一个参数
	// 此参数为接收者
	p := Person{name: "xiasen"}
	fmt.Println(p.String())

	// 2. 值传递与引用传递，
	// 默认为值传递，但是也可以定义传指针
	// 也可定义引用传递（使用指针）

	// 值引用
	p.modify()
	fmt.Println(p.String()) // still xiasen
	// 传指针
	(&p).modify()
	fmt.Println(p.String()) // still xiasen

	p.modify2()
	fmt.Println(p.String()) // change to hhhhh

	// p.Add 可以赋值给一个方法变量，相当于一个函数
	// 只需要提供实参而不需要接收者
	p1 := Point{x: 1, y: 2}
	p2 := Point{3, 4}
	f := p1.Add
	fmt.Println(f(p2))
	fmt.Println(p1)

	// 方法表达式
	// 提供接收者和形参
	// 类型.方法
	f1 := Point.Add
	fmt.Println(f1(p1, p2))

	p1.Sum(p2)
	fmt.Println(p1)
}

type Person struct {
	name string
}

func (p Person) String() string {
	return "person name is " + p.name
}

func (p Person) modify() {
	p.name = "Jason"
}

func (p *Person) modify2() {
	p.name = "hhhh"
}

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p *Point) Sum(q Point) {
	p.x += q.x
	p.y += q.y
}
