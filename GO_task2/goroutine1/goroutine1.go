/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，
另一个协程打印从2到10的偶数。
- 考察点 ： go 关键字的使用、协程的并发执行。
*/
package main

import (
	"fmt"
	"sync"
)

func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数是：", i)
	}
}
func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数是：", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintOdd(&wg)
	go PrintEven(&wg)
	wg.Wait()
	fmt.Println("协程打印完毕！")
}
