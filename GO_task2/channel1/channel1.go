/*
 1. 编写一个程序，使用通道实现两个协程之间的通信。

一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
- 考察点 ：通道的基本使用、协程间通信。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2) // 2 个 goroutine

	//发送者
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch <- i
		}
		close(ch) // 告诉接收者没有数据了
	}()
	//接收者
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收：", num)
		}
	}()
	wg.Wait()
	fmt.Println("完成！")
}
