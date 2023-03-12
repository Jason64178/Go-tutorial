package variable

import "fmt"

func VariableTypes() {
	// 1. 整数
	// 要求算术或逻辑运算时，类型一致
	var a int = 10
	var b int32 = 9
	// fmt.Println(a + b) 报错
	// 转换
	fmt.Println(a + int(b))

	// 2. 数值运算
	// % 取模，符号与被除数一致
	fmt.Println(5 % 3)
	fmt.Println(-5 % -3)
	// 除法
	fmt.Println(5 / 3)     // 1
	fmt.Println(5.0 / 3.0) // 1.66666666666666667

	// 3. 比较运算
	// >, >=, <, <=, ==, !=
	// 不同类型不可比较
	var i int32
	var j int64

	// if i == j 报错
	if i == 1 || j == 2 {
		fmt.Println("equal.")
	}

	// 4. 位运算
	// & | ^ &^ << >>

	// 5. 浮点数，float32以及float64
	// 默认为float64

	f := 10.0
	fmt.Println(f)

	// 当对浮点数进行比较运算时，不能直接使用 == 和 !=，结果会不稳定。应该使用 math 标准库。

	// 6.复数, complex64 和 complex128
	// complex(2,19) 构造复数
	// real(x) 实部 imag(x) 虚部
	var x complex64 = 3 + 5i
	var y complex128 = complex(4, 10)

	fmt.Println(real(x), imag(x)) // 输出 3 5
	fmt.Println(real(y), imag(y)) // 输出 4 10

	// 7. boolean, 不能与整型相互转换
	// if 后必须加bool型

	ok := true
	fmt.Println(ok)

	// 8. 字符串 string
	// 不能改变其中的值 immutable
	s1 := "hello"
	s2 := "world"

	s := `row1\r\n
	row2`
	fmt.Println(s1, s2, s)
	s3 := "row1\r\nrow2"
	fmt.Println(s3)
	// ``使用保留原始字符串

	// 拼接字符串
	s4 := s1 + s2
	fmt.Println(s4)
	fmt.Println(len(s4))

	// 字符串切片
	fmt.Println(s4[2:4])
	fmt.Println(s4[:4])
	fmt.Println(s4[2:])
	fmt.Println(s4[:])

	// 遍历字符串
	s5 := "hello 世界"

	for i := 0; i < len(s5); i++ {
		fmt.Println(i, s5[i])
	}

	// 输出
	// 0 104
	// 1 101
	// 2 108
	// 3 108
	// 4 111
	// 5 32
	// 6 228
	// 7 184
	// 8 150
	// 9 231
	// 10 149
	// 11 140

	for i, v := range s5 {
		fmt.Println(i, v)
	}

	// 输出
	// 0 104
	// 1 101
	// 2 108
	// 3 108
	// 4 111
	// 5 32
	// 6 19990
	// 9 30028
}
