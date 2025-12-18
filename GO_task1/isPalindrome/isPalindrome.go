/*回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

示例 1：
输入：x = 121
输出：true
*/

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// ...existing code...

func main() {
	// 按注释示例构建测试用例并打印结果
	tests := []int{
		121,  // 示例：回文 -> true
		-121, // 负数不是回文 -> false
		10,   // 不是回文 -> false
		1221, // 回文 -> true
		0,    // 回文 -> true
	}

	for _, tc := range tests {
		fmt.Printf("输入: %d -> 输出: %v\n", tc, isPalindrome(tc))
	}
}
