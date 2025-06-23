package main

import (
	"bufio"
	. "fmt"
	"regexp"
	"strings"
)

func main() {
	// 统计输入的行数, 和 输入的单词数
	// input := os.Stdin
	var inputStr = "hello world\n second __line\n hwnd line\n"
	input := strings.NewReader(inputStr)
	scanner := bufio.NewScanner(input)

	// 行数
	// scanner.Split(bufio.ScanLines)
	// var sum = 0
	// var words
	// for scanner.Scan() {
	// 	Printf("%#v\n", scanner.Text())
	// 	Printf("%s\n", scanner.Text())
	// 	sum++

	// }
	// Println("lines:", sum)

	// 单词数
	scanner.Split(bufio.ScanWords)
	var words = 0
	for scanner.Scan() {
		Printf("%s\n", scanner.Text())
		words++
	}
	Println("words:", words)

	wordV2 := CalcWords(inputStr)
	Printf("wordV2:%d\n", wordV2)

	liss := strings.Fields(inputStr)
	Println("strings.Fields:", liss)
	wordCount := make(map[string]int)
	for _, v := range liss {
		wordCount[v]++
	}
	Println(wordCount)
	var mapKeys []string
	for key := range wordCount {
		mapKeys = append(mapKeys, key)
	}
	Println(mapKeys)

	var ko string = " hello"
	for index, v := range ko {
		// 注意v是一个rune类型, 是一个int32类型的unicode码, 需要string转为一个具体值
		Printf("index:%d, v:%c , vType:%T v:%v, string:%s\n", index, v, v, v, string(v))
	}
}

func CalcWords(s string) int {
	// 计算单词数
	v2 := strings.Replace(s, "\n", "", -1)
	Println("v2:", v2)
	list := strings.Split(v2, " ")

	//前后字符串 并且 w是字符串 | 下划线 | 数字, + 是1或者无数个
	pattern := `^\w+$`
	regex2 := regexp.MustCompile(pattern)
	for _, v := range list {
		Println(v)
		var matched bool = regex2.MatchString(v)
		if matched {
			Println("matched: ", v)
		}
	}

	Println(list)
	// "sadf afd 2dasdf 2dead"
	// i 是最左侧字符, j = ' ' | j == len 中断统计
	return len(list)
}
