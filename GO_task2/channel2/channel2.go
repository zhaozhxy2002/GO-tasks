/*  实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
    - 考察点 ：通道的缓冲机制。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10) //缓冲区大小为 10,延迟数据阻塞
	var wg sync.WaitGroup
	wg.Add(2)

	//生产者协程
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- i
		}
		close(ch)
	}()
	//消费者协程
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收：", num)
		}
	}()
	wg.Wait()
	fmt.Println("全部处理完毕！")
}
