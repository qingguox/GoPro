package main

import (
	. "fmt"
)

func main() {
	// map是使用
	map1 := make(map[string]int)
	map1["Google"] = 1
	map1["Facebook"] = 2

	Println("len", len(map1))

	var searchKey = "faceBook"
	searchKey = "Facebook"
	value, ok := map1[searchKey]
	if ok {
		Println("searchKey: ", searchKey, "existed! value:", value)
	} else {
		Println("searchKey: ", searchKey, "not existed!")
	}

	// 不存在value2 = 0 默认值, 但是可以用value, ok  ok==true存在 false不存在
	value2 := map1["11"]
	Println(value2)

	delete(map1, "Facebook")
	Println(map1)
}
