package main

import (
	"bytes"
	. "fmt"
	"io"
	"net/http"
	"strings"

	// "log"
	"os"
)

var httpPrefix = "http://"

func main() {
	Println(os.Args[:])
	// go build fetch.go
	// ./fetch gopl.io
	// output: [./fetch gopl.io]
	// 网页信息

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if resp.StatusCode != 200 {
			Fprintf(os.Stderr, "fetch respStatusCode:%d !=200 \n", resp.StatusCode)
			os.Exit(1)
		}
		var bytes bytes.Buffer
		wd := io.ReadWriter(&bytes)
		wBytes, err := io.Copy(wd, resp.Body)
		resp.Body.Close()
		if err != nil {
			Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
		Printf("%7d\n", wBytes)

		lens := len(bytes.Bytes())
		Println(lens)
		Printf("%s\n", bytes.String())
		// var data = make([]byte, lens)
		// wd.Read(data)
		// Printf("%s\n", string(data))
	}
}
