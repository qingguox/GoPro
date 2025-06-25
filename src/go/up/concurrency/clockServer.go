package main

import (
	"errors"
	"flag"
	. "fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// ./clockServer -port=8000
	// go run clockServer.go -port=8000

	args := os.Args
	serverSequence := time.Now().UnixMilli()
	portPtr := flag.Int("port", 0, "端口")
	// 解析命令行参数
	flag.Parse()

	Println("portPtr:", *portPtr)
	port := *portPtr
	if port <= 0 {
		log.Print(errors.New("port <= 0 args:" + strings.Join(args, " ")))
		os.Exit(1)
	}

	Printf("****ClockServer sequence:%d Port:%d %T***", serverSequence, port, port)

	// 编写一个时钟服务器
	address := "localhost:" + strconv.Itoa(port)
	Println("adress:", address)
	result, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 阻塞获取一个客户端连接
		conn, err := result.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 这里并行处理, 连接无上限, 系统资源是上限
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {

		_, err := io.WriteString(c, c.RemoteAddr().String()+"\t"+time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
