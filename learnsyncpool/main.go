package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// 1. 创建一个对象池，使用&sync.Pool{}创建一个对象池
var pool = &sync.Pool{
	New: func() any {
		log.Println("create a new object")
		b := make([]byte, 1024)
		return &b
	},
}

func main() {

	// 一个使用sync.Pool的例子
	log.Println("main start")
	fmt.Println("这是一个使用sync.Pool的例子")

	// 2. 获取一个对象，如果对象不存在，会调用New函数创建一个
	obj := pool.Get().(*[]byte) // 2.1 类型断言.(*[]byte) 为 *[]byte 指针类型

	// 3. 确保使用完对象后，将对象放回对象池
	defer pool.Put(obj)

	// 4. 清空对象池
	*obj = (*obj)[:0]

	// 5. 使用对象
	*obj = append(*obj, "hello world"...)
	fmt.Println("对象池内容:", string(*obj))


	// obj = append(*obj, "hello world"...) // "hello world"... 是一个切片字面量，表示一个字符串切片
	// fmt.Println("对象池内容:", string(obj))
	log.Println("main end")

	os.Exit(0)
}
