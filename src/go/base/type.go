package main

import (
	. "fmt"
	"strconv"
)

func main() {
	Println("***类型转换****")
	// type(expression)
	var a int = 10
	var b float64 = float64(a)
	Println("b:=", b)

	var sum int = 9
	count := 2
	var mean float32

	mean = float32(sum) / float32(count)
	Println("mean :=", mean)

	// 1. 字符串类型转换
	// 1.1字符串转为整数
	var str string = "10"
	var num int
	num, err := strconv.Atoi(str)
	if err != nil {
		Println("字符串转为整数失败!! num:", num, "err:=", err)
	} else {
		Println("字符串转为整数成功!! num:", num)
	}
	// 1.2 整数转为字符串
	num = 124
	str = strconv.Itoa(num)
	Println("整数转为字符串!! str:", str)

	// 1.3 字符串转为浮点数
	str = "3.14145"
	num2, _ := strconv.ParseFloat(str, 64)
	Println("字符串转为浮点数: num", num2)

	// 1.4 浮点数转为字符串
	num3 := 3.14
	str2 := strconv.FormatFloat(num3, 'f', 2, 64)
	Println("浮点数转为字符串, ", str2)

	// 1.5 类型断言 value.(type)
	var i interface{} = "hello world!"
	strk, ok := i.(string)
	if ok {
		Printf("%s is a string\n", strk)
	} else {
		Printf("conversion fail!\n")
	}

	// 1.6 类型转换 T(value)
	// 只有指针在方法中才能修改对象的属性值
	var w Writer = &StringWriter{"init"}
	sw := w.(*StringWriter) // 这一步是必须的

	Println(sw.str)

	sw.Write([]byte("ko"))
	Println("sw change :", sw.str)

	ko, _ := w.Write([]byte("w"))
	Println("ko", ko)
	// 必须强转
	// Println("w :", w.str)

	// 空接口类型
	printValue(4)
	printValue("s1")
	printValue(1.22)

	//
	PrintValue("hello")
	PrintValue(12.11)
	PrintValue(223)
	PrintValue([]int{1, 2, 2})

	var ko2 interface{}
	Println("ko2 == nil : ", ko2 == nil)
}
func PrintValue(val interface{}) {
	Printf("value: %v, type: %T \n", val, val)
}

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str = string(data)
	return len(data), nil
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		Println("Integer: v", v)
	case string:
		Println("String: v", v)
	default:
		Println("Unknown type ")
	}
}
