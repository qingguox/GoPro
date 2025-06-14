package main

import (
	"errors"
	. "fmt"
)

func main() {

	Println("***Error Study****")

	// 1. error接口
	erro := errors.New("This is error!")
	Println("err: ", erro)
	// 2. 显示返回值
	err2 := printError("vistor!")
	Println(err2)
	// 3. 自定义错误

	cus := &CustomError{2001, "用户未登陆!"}
	Println(cus)

	// 4. errors.Is(err, ko)  errors.As(er, 2w)
	if errors.Is(findItem((1)), ErrNotFound) {
		Printf("ErrNotFound\n")
	}

	err3 := getError()
	var myError *CustomError
	if errors.As(err3, &myError) {
		Println("error:", myError)
	}
	// 4. panic 和 recover: 处理不可恢复的严重错误
	// 	Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。

	// panic:
	// 导致程序崩溃并输出堆栈信息。
	// 常用于程序无法继续运行的情况。
	// recover:
	// 捕获 panic，避免程序崩溃。
	/*
			 	1、panic 在没有用 recover 前以及在 recover 捕获那一级函数栈，panic 之后的代码均不会执行；一旦被 recover 捕获后，外层的函数栈代码恢复正常，所有代码均会得到执行；
		 		2、panic 后，不再执行后面的代码，立即按照逆序执行 defer，并逐级往外层函数栈扩散；defer 就类似 finally；
		 		3、利用 recover 捕获 panic 时，defer 需要再 panic 之前声明，否则由于 panic 之后的代码得不到执行，因此也无法 recover；
	*/

	Println("Start Program...")
	safeFunction()
	Println("End")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			Println("recover 捕获panic崩溃, r: [", r, "]")
		}
	}()
	panic("panic: 触发或者抛出崩溃!!")
}

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	return Errorf("database error: %w", ErrNotFound)
}

func printError(s string) error {
	return errors.New("err:" + s)
}

type CustomError struct {
	code int
	msg  string
}

func (err *CustomError) Error() string {
	return Sprintf("Custom Error! code:%d, msg:%s\n", err.code, err.msg)
}

func getError() error {
	return &CustomError{404, "Not Found"}
}
