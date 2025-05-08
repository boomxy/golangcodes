package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	log.Println("Starting the application...")
	// 学习 Go 语言的 WaitGroup

	// 1. 创建一个 WaitGroup
	var wg sync.WaitGroup

	// 2. 计数器加 2， 表示要等待两个 goroutine 完成。
	wg.Add(2)

	// 3. 启动两个 goroutine
	go func() {
		defer wg.Done() // 计数器 -1
		log.Println("Hello, World!")
	}()
	go func() {
		defer wg.Done() // 计数器 -1
		log.Println("Hello, Go!")
	}()
	// 4. 等待所有 goroutine 完成
	wg.Wait()
	// 等待所有goroutine完成，后边没有其他代码的时候可以不用 defer wg.Wait()
	// 如果我们把Wait()放在defer语句中，那么Wait()会在main()函数返回之前执行。
	//defer wg.Wait()
	log.Println("All goroutines finished.")
	// 5. 退出程序
	os.Exit(0)
	// 6. 结束
}
