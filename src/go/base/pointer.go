package main

import (
	. "fmt"
)

func main() {

	Println("****Pointer*****\n")
	// 1. 指针基础知识, 和 java没大的区别
	var a int = 20
	var ip *int

	ip = &a

	Printf("a 变量的地址为: %x \n", &a)
	Printf("ip 变量存储的指针地址为: %x \n", ip)

	Printf("ip 变量的值为: %d\n", *ip)

	// 空指针 prt == nil
	var prt *int
	Printf("prt 值为: %x\n", prt) // 0

	var arr = [3]int{1, 3, 4}
	var pt2 [3]*int       // 指针数组
	var pt *[3]int = &arr // 数组指针
	for k, _ := range arr {
		pt2[k] = &arr[k]
	}
	for i := 0; i < 3; i++ {
		Printf("%d, ppt2:%d \n", (*pt)[i], *pt2[i])
	}

	// 2. 指向指针指针
	var ptr4 *int
	var ptr5 **int

	a = 400
	ptr4 = &a
	ptr5 = &ptr4
	Printf("变量a = %d\n", a)
	Printf("指针变量 *ptr4 = %d\n", *ptr4)
	Printf("指向指针的指针变量 **ptr5 = %d\n", **ptr5)

	// 指针传递到函数中
	var x int = 1
	var y int = 2

	Printf("交换之前 x:%d, y:%d\n", x, y)
	SwapXy(&x, &y)
	Printf("交换之后 x:%d, y:%d\n", x, y)
}

func SwapXy(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
