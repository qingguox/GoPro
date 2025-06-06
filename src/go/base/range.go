package main

import (
	. "fmt"
)

func main() {
	Println("*****range study*****")

	// range 应用于 map, 数组, slice, 字符串, channel等
	// 1. map
	mapExp := make(map[int64]float32)
	mapExp[1] = 1.2
	mapExp[2] = 1.9
	Println(mapExp)

	for key, value := range mapExp {
		Printf("mapExp key:%d, value:%f\n", key, value)
	}
	for key := range mapExp {
		Printf("mapExp key:%d, \n", key)
	}
	for _, value := range mapExp {
		Printf("mapExp value:%.2f, \n", value)
	}

	// 2. 数组 i =0 开始
	var pow = [...]int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		Printf("list 2**i:%d = v:%d\n", i, v)
	}

	// 3. 字符串 i=0开始, 中文采用的是UTF-8编码, 用for i :=0; i< len(string1); i++ { %x, s[i]}
	var string1 = "hello"
	for i, v := range string1 {
		Printf("index: %d, chat :%c\n", i, v)
	}

	// 4. 通道 先进先出
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for v := range ch {
		Printf("ch v:%d\n", v)
	}
	// 5. 涉及指针多注意, v 是个单独的地址
	nums := [3]int{5, 6, 7}
	for k, v := range nums {
		Println("原地址:", &nums[k], "\t value地址:", &v)
		Printf("原地址: %x \t value地址: %x\n", &nums[k], &v)
	}

}
