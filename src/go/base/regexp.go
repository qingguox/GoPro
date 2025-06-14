package main

import (
	. "fmt"
	"regexp"
)

func main() {
	Println("***正则表达式regexp***")

	// 正则表达式的基本语法
	// 以下是一些常用的正则表达式语法：

	// .：匹配任意单个字符（除了换行符）。
	// *：匹配前面的字符 0 次或多次。
	// +：匹配前面的字符 1 次或多次。
	// ?：匹配前面的字符 0 次或 1 次。
	// \d：匹配数字字符（等价于 [0-9]）。
	// \w：匹配字母、数字或下划线（等价于 [a-zA-Z0-9_]）。
	// \s：匹配空白字符（包括空格、制表符、换行符等）。
	// []：匹配括号内的任意一个字符（例如 [abc] 匹配 a、b 或 c）。
	// ^：匹配字符串的开头。
	// $：匹配字符串的结尾。

	// Go 语言的标准库提供了 regexp 包，用于处理正则表达式。以下是 regexp 包中常用的函数和方法：
	// 1. Compile 和 MustCompile
	// 用于编译正则表达式。Compile 返回一个 *Regexp 对象和一个错误，而 MustCompile 在编译失败时会直接 panic。

	// 2. MatchString
	// 检查字符串是否匹配正则表达式。

	pattern := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(pattern)

	str := "hello123"
	if regex.MatchString(str) {
		Println("字符串匹配表达式")
	} else {
		Println("字符串不匹配表达式")
	}

	// 3. FindString 和 FindAllString
	// 用于查找匹配的字符串。FindString 返回第一个匹配项，FindAllString 返回所有匹配项。
	pattern = `\d+`
	regex = regexp.MustCompile(pattern)

	str = "我有 3 个苹果 和 5 个 香蕉"
	matches := regex.FindAllString(str, -1)
	Println("找到的数字: ", matches)

	// 4. ReplaceAllString
	// 用于替换匹配的字符串。
	pattern = `\s+` // 匹配空白字符, +代表 1个或者多个
	regex = regexp.MustCompile(pattern)

	str = "hello   World"
	result := regex.ReplaceAllString(str, " ")
	Println("result :", result)

	// 5. Split
	// 根据正则表达式分割字符串。
	pattern = `,`
	regex = regexp.MustCompile(pattern)

	str = "apple,banana,orange,"
	parts := regex.Split(str, -1)
	Println("分隔符, 分割后切片:", parts)
}
