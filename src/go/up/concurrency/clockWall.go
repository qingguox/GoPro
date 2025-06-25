package main

import (
	. "fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	Println("***** ClockWall *****")
	argList := os.Args[1:]
	var nameAndAddressMap map[string]string = make(map[string]string)
	for _, v := range argList {
		ko := strings.Split(v, "=")
		name := ko[0]
		address := ko[1]
		nameAndAddressMap[name] = address
	}
	Println("nameAndAddressMap: ", nameAndAddressMap)
	for k, v := range nameAndAddressMap {
		Printf("name:%s, address:%s\n", k, v)
		go client(k, v)
	}
	time.Sleep(1 * time.Minute)
}

func client(name, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	Printf("name:%s ", name)
	mustCopy(os.Stdout, conn)
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
