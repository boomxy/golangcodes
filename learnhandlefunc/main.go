package main

import (
	"fmt"
	"log"
	"os"
)

/* 阅读 http/server.go 源码 
	了解 handlerFunc 的实现
*/

// 创建一个函数接受者，用于 log 打印
type logReciver func(string)

// logReciver 实现了 log 接口
func (l logReciver) LogMe(s string) {
	l(s)
}

func MyPrintln(s string) {
	fmt.Println("fmt:>>>", s)
	log.Println(s)
}

func main() {
	// 创建一个 logReciver
	lr := logReciver(MyPrintln)

	// 调用 logReciver 的 LogMe 方法
	lr.LogMe("Hi, you")

	// 调用 logReciver 的 LogMe 方法
	lr.LogMe("Hello, World!")

	os.Exit(0)
}
