package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func getContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取内容失败。")
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	return string(result)
}

func writeFile(filepath, content string) {
	// 创建文件
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	// 写入文件
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 获取response内容
	content := getContent("http://zmzlove.xyz/wordpress")
	// 创建文件
	writeFile("./blog01.html", content)

}
