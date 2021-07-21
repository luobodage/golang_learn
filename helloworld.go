package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
)

func swap(a, b int) (int, int) {
	return b, a
}

func arsonist() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fileInputAndOutputOperations()
}

func fileInputAndOutputOperations() {
	const filename = "abd.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args "+"(%d ,%d)\n", opName, a, b)
	return op(a,b)

}

func pow(a,b int) int  {
	return int(math.Pow(float64(a),float64(b)))
}

func main() {
	var a = "hello "
	fmt.Println(a)
	b := "bytedance"
	fmt.Println(a + b)
	c := "world!"
	println(a + c)
	var e, d = swap(1, 2)
	fmt.Println(e, d)
	arsonist()
	printFile("abd.txt")
	println(apply(pow, 1, 2))

}
