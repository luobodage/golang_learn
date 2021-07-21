package main

import (
	real2 "first/retriever/real"
	"fmt"
	"io"
	"os"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://zmzlove.xyz/wordpress")
}

func main() {
	var r Retriever
	r = real2.Retriever{}
	var content string = download(r)
	f, err := os.Create("blog.html")
	if err != nil {
		panic(err)
	}
	n, err := io.WriteString(f, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("写入%d字节", n)
	}

}
