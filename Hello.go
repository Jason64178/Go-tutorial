// 00_hello.go

package main // 声明 main 包

import (
	"desktop-tutorial/flow"
	"desktop-tutorial/types"
	"desktop-tutorial/variable"
	"fmt" // 导入内置 fmt 包
)

func main() { // main函数，程序执行入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
	variable.Variables()
	variable.VariableTypes()
	variable.ArraySlice()
	variable.MapStruct()
	flow.Flow()
	flow.Funcs()
	flow.Errors()
	types.MethodMain()
	types.InterfaceMain()
}
