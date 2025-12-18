/*
加一
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。

示例 2：
输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。
*/
package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1 //默认从最后一位开始，每一位都加一
		if digits[i] == 10 {
			digits[i] = 0 //如果这位需要进位，则直接置0

		} else if digits[i] < 10 {
			return digits //如果这一位不需要进位，循环终止，i之前的数字不用再循环digits[i] = digits[i] + 1
		}
	}
	//如果循环结束了还没有return结果，则说明每一位都进位为0
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	// 按注释示例构建测试用例并打印结果
	tests := [][]int{
		{1, 2, 3},    // 示例1 -> [1,2,4]
		{9},          // 示例2 -> [1,0]
		{4, 3, 2, 1}, // 额外测试 -> [4,3,2,2]
		{9, 9, 9},    // 全位进位 -> [1,0,0,0]
	}

	for _, tc := range tests {
		fmt.Printf("输入: %v -> 输出: %v\n", tc, plusOne(append([]int(nil), tc...)))
	}
}
