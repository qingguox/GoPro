package main

import (
	. "fmt"
)

func main() {
	Println("*** slice 切片, 动态数组, 类似java中ArrayList")
	// slice可变数组, 可以追加元素, 追加后 cap 和 len都变了
	var slice1 []int

	// PrintSlice(slice1)
	if slice1 == nil {
		Println("slice1 是空的")
	}

	slice1 = make([]int, 3, 4)
	slice1[0] = 11
	slice1[1] = 22
	PrintSlice(slice1)

	slice2 := []int{1, 2, 4}
	PrintSlice(slice2)

	// 容量是k slice[i:j]  下标i ~ 下标j-1
	// len: j - i
	// cap: k - i  实际上此时 cap是从 原slice中的i下标开始的arr
	s2 := slice2[1:]
	PrintSlice(s2)

	s1 := s2
	Println("PrintSliceV1****")
	Println(s1)
	PrintSliceV1(s1)
	Println(s1)

	Println("PrintSliceV2****")
	Println(s2)
	PrintSliceV2(&s2)
	Println(s2)

	// todo 切片 append 扩容原理
	// 当slice append 一些新元素时
	// 1. cap = 旧cap + 新增len
	// 2. 然后 GOROOT/src/runtime/slice.go源码是这样的:
	// newcap := old.cap
	// doublecap := newcap + newcap
	// if cap > doublecap {   // cap > 2倍的 oldCap
	//     newcap = cap
	// } else {
	//     if old.len < 1024 {
	//         newcap = doublecap
	//     } else {
	//         // Check 0 < newcap to detect overflow
	//         // and prevent an infinite loop.
	//         for 0 < newcap && newcap < cap {   // 1/4增长
	//             newcap += newcap / 4
	//         }
	//         // Set newcap to the requested cap when
	//         // the newcap calculation overflowed.
	//         if newcap <= 0 {    // 溢出后, 回归为cap
	//             newcap = cap
	//         }
	//     }
	// }
}

func PrintSlice(x []int) {
	Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func PrintSliceV1(x []int) {
	// 直接修改会生效
	x[1] = 100
	x = append(x, 7) // 添加一个元素, 函数外不生效
	Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func PrintSliceV2(x *[]int) {
	(*x)[1] = 20
	*x = append(*x, 11) // 添加一个元素, 函数外也生效
	Printf("len=%d cap=%d slice=%v\n", len(*x), cap(*x), x)
}
