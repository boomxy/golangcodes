package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Starting...")
	// 1. 创建 一个缓冲区大小为 10 的 channel
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			// 2. 往 channel 中写入数据
			// 3. 这里会阻塞，直到有一个 goroutine 从 channel 中读取数据
			ch <- i * 2
		}
		// 4. 关闭 channel
		close(ch)
	}()
	// 5. 从 channel 中读取数据
	// 6. 这里会阻塞，直到 channel 中有数据可读，最终直到 channel 被关闭
	for v := range ch {
		fmt.Println(v)
	}
	log.Println("Done")
	os.Exit(0)
	// log.Println("This will not be printed")
}
