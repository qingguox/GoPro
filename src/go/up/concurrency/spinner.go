package main

import (
	. "fmt"
	"time"
)

func main() {
	// go
	Println("go routinue test")
	// go spinner(100 * time.Millisecond)
	const n = 45
	fiN := fib(n)
	Printf("\rFibonacci(%d) = %d\n", n, fiN)

	// \r 光标会到下一行第一位
	// 主函数返回时，所有的goroutine都会被直接打断，程序退出， 但是有例外
}

func spinner(delay time.Duration) {
	var count int = 0
	for {
		for _, v := range `-\|/` {
			count++
			Printf("\r%c, %d", v, count)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
