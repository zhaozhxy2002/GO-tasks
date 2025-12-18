/*  题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
    - 考察点 ：协程原理、并发任务调度。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func TasksSchedule(tasks []func()) {
	var wg sync.WaitGroup
	//遍历任务列表，每一个任务列表都需要启动goroutine，任务开始和任务结束在这里标记，并每一个都记录时间
	for i, task := range tasks {
		wg.Add(1)
		//开始遍历，遍历i+1则加一个任务
		//这个协程输入遍历的i值和闭包函数t，用来简易指代任务的具体实现函数
		go func(i int, t func()) {
			defer wg.Done()
			start := time.Now()
			fmt.Printf("===第%d个任务开始执行===\n", i+1)
			t() //执行任务
			duration := time.Since(start)
			fmt.Printf("===第%d个任务结束,用时%v===\n", i+1, duration)
		}(i, task) //把具体的值传入
	}
	// Wait()用来等待所有需要等待的goroutine完成
	wg.Wait()
}

// Sleep 模拟每个任务不同的工作量/耗时
func task1() {
	time.Sleep(2 * time.Second)
	fmt.Println("任务1：数据清洗完成！")
}

func task2() {
	time.Sleep(1 * time.Second)
	fmt.Println("任务2：图像处理完成！")
}

func task3() {
	time.Sleep(3 * time.Second)
	fmt.Println("任务3：日志分析完成！")
}

func main() {
	//定义一个函数切片，来表示t []func()多个不同但同类的函数
	Tasks := []func(){
		task1,
		task2,
		task3,
	}
	TasksSchedule(Tasks)
}
