package main

import (
	. "fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	Println("***** ClockClient*****")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
