package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("main start")
	// 学习 select
	// 1. select 是 Go 语言中的一个关键字，用于处理多个通道的发送和接收操作。
	// 2. select 会阻塞直到其中一个通道可以进行操作。
	// 3. select 可以用于实现超时机制。
	// 4. select 可以用于实现多路复用。
	// 5. select 可以用于实现异步编程。
	// 6. select 可以用于实现并发编程。
	// 7. select 可以用于实现协程间的通信。
	// 8. select 可以用于实现协程间的同步。
	// 9. select 可以用于实现协程间的互斥。
	// 10. select 可以用于实现协程间的共享资源。
	// 11. select 可以用于实现协程间的消息传递。
	// 12. select 可以用于实现协程间的事件驱动编程。
	// 13. select 可以用于实现协程间的回调函数。
	// 14. select 可以用于实现协程间的信号量。
	// 15. select 可以用于实现协程间的条件变量。
	// 16. select 可以用于实现协程间的读写锁。
	// 17. select 可以用于实现协程间的读写通道。

	numch := make(chan int, 1)
	// 1. 创建一个通道，容量为 1
	for i := 0; i < 10; i++ {
		select {
		case num := <-numch: // 接收数据, 从 numch 中接收数据
			log.Println("--numch recv:", num)
		case numch <- i: // 发送数据， 将 i 发送到 numch 中
			log.Println("+++numch send:", i)
		}
	}

	ch := make(chan int, 3)
	quit := make(chan int)
	go func() {
		for i := 0; i < 12; i++ {
			ch <- i + 63
		}
		quit <- 0
	}()

	for {
		select {
		case v := <-ch:
			log.Println("main recv:", v)
		case <-quit:
			log.Println("main quit")
			// 退出协程
			// 关闭通道
			close(ch)
			close(quit)
			// 退出程序
			os.Exit(0)
		}
	}
}
