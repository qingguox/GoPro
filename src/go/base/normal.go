package main

import (
	. "fmt"
	. "strings"
	"unsafe"
)

/*
1. go基础数据结构: bool, int float32 float64 复数, 字符串(一串固定长度的字符连接起来的字符序列, 单字节拼接起来的 UTF-8 Unicode文本)
	1.2 其他类型: 指针, 数组, struct, Channel, 函数func, 切片, 接口interface, Map类型

2. 变量 var in type or var in1, in2 type

*/

var totalX = 10

func prinX(x int) {
	var srt string = ""
	for i := 1; i <= x; i++ {
		srt = srt + "*"
	}
	Println(srt)
	Println("")
}

func main() {

	prinX(totalX)
	dataAndStruct()

	prinX(totalX)

	variableFunc()
	prinX(totalX)

	operate()
	prinX(totalX)
}

func operate() {
	// Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

	var a int = 10
LOOP: // 循环
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP // 直接跳转到 LOOP语句
		}
		Printf("a的值为 : %d\n", a)
		a++
	}

	numbers := [6]int{1, 2, 4, 55, 5, 5}
	for i, x := range numbers {
		Printf("第%d位 x 的值 = %d \n", i, x)
	}
	// for _, x := range numbers {   _ 代表不想用index
	// 	Printf("第%d位 x 的值 = %d", x)
	// }
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[3] = 3.0

	for key, value := range map1 {
		Printf("key:%d, value:%f\n", key, value)
	}
	// 读取key
	// for key := range map1
	// 读取value
	// for _, value := range map1

	// 数组
	// var arrayName [size]dataType
	var balance [10]float32
	balance[0] = 1

	var balance2 = [6]int{1, 2, 4, 2, 45}
	balance3 := [...]int{2, 3, 4}       // 数组大小不确定
	balance4 := [...]int{1: 2, 3, 4: 7} // 直接初始化 index:对应的值
	printArray(balance2[:])
	printArray(balance3[:])
	printArray(balance4[:])

	// 多为数组
	balance5 := [][]int{} // x行 y列
	b1 := []int{1, 2, 5}
	b2 := []int{2, 5, 5}
	balance5 = append(balance5, b1)
	balance5 = append(balance5, b2)

	Println(balance5)

	balance6 := [...][5]int{
		{0, 2, 4},
		{2, 4, 5, 6},
		{2}}
	Println(balance6)
}

func printArray(x []int) {
	var size = len(x)
	for i := 0; i < size; i++ {
		Printf("%d", x[i])
	}
	Println()
}

func variableFunc() {
	ax := "hello"
	Println(unsafe.Sizeof(ax)) // 16 字符串在go语言中是个结构, 包含指向底层数组的指针和长度, 每部分都是8字节, 故为16字节

	// 定义常量的时候, 如果不提供初始值, 则表示将使用上行的表达式. b = 1 c=1s
	const (
		a = 1
		b
		c
	)

	// iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。
	// i = 0, j=1, x = 2 编译器会从0开始 调用一次递增1
	const (
		i = iota
		j
		x
	)
	const xx = iota // 0
}

func dataAndStruct() {

	// 布尔型, 直接定义并赋值
	var b bool = true
	var b3 = false // 忽略类型的声明
	Println(b)
	Println(b3)
	// 先定义, 后赋值
	var b2 bool
	// b2 := true 这种是不用定义
	b2 = false
	Println(b2)

	// int类型: 有两种 uint int uint8 uint16 uint32 uint64 无符号 + 有符号 一共 8种
	// float类型: float32 + float64 complex64 complex128
	// 其他数字类型: byte(类似uint8), rune(类似int32), uint(32或者64), int(与uint一样大小), uintptr(无符号整型, 用于存放一个指针)

	var string2 = "ni ha o1 "
	Println(string2)
	string3 := Replace(string2, "ha", "replace", -1)
	Println(string3)
}
