/*
2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
  - 考察点 ：指针运算、切片操作。
*/
package main

import (
	"fmt"
)

func multiply(Ptr *[]int) {
	for i := range *Ptr {
		(*Ptr)[i] *= 2
	}

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("原始切片：%v\n", nums)
	multiply(&nums)
	fmt.Printf("处理后切片：%v\n", nums)

}
