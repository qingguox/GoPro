package main

// package含义:
// 	Go 程序是通过 package 来组织的。
// 	只有 package 名称为 main 的源码文件可以包含 main 函数。
// 	一个可执行程序有且仅有一个 main 包。
// 	通过 import 关键字来导入其他非 main 包。
// 	import fmt2 "fmt" 起一个别名

import (
	. "fmt"

	"./myMath" // 自定义包, 注意mod的关联性
)

// 全局变量
var s = 1001
var totalDesc = "全局变量"

type myStruct struct {
	name string
	age  int64
}
type myInterface interface {
}

// 可见性规则 使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用。
// 首字母小写, private
func myFuncConvert(x int, y string) (string, int) {
	return y, x
}

// 首字母大写, public
func Swap(x int, y int) (int, int) {
	tmp := x
	x = y
	y = tmp
	return x, y
}

func main() {
	Print("hello ]nw")
	Print("sa ")

	// 1. 基本的语法
	// 	通过 const 关键字来进行常量的定义。
	const ads, str = 1.0, "nihao1"
	Println(ads)
	// 过在函数体外部使用 var 关键字来进行全局变量的声明和赋值。
	Println(s)

	// 2. 关键字 type
	// 通过 type 关键字来进行结构(struct)和接口(interface)的声明。
	mssda12 := myStruct{"2", 2}
	Println(mssda12.age)
	Println("my func")
	// 通过 func 关键字来进行函数的声明。
	k, j := myFuncConvert(int(mssda12.age), mssda12.name)

	// 3. 输入输出
	// Sprintf是格式化参数生成格式化的字符串, 返回具体的结果
	// Printf是格式化参数生成格式化的字符串, 并直接输出
	ko := Sprintf("%4d, %s\n", j, k)
	Println(ko)

	Printf("printf %4d, %s\n", j, k)
	// Println("%s", "你好") 结果: %s 你好 这种不行

	// swap函数
	var m, y int = 1, 2
	x1, x2 := Swap(m, y)
	Printf("[Swap] x1:%d, x2:%d origialX:%d Y:%d\n", x1, x2, m, y)

	// 引入外面的包
	// 会遇到找不到包的情况, 直接: go env -w GO111MODULE=auto
	Println(myMath.AddInt(1, 2))
}
