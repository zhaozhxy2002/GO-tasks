/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ { //从第一个字符串的第i个字符开始遍历，同时i一定是小于第一个字符串的长度的
		prefix := strs[0][i]             //先默认第一个字符串的第一个字符为prefix
		for j := 1; j < len(strs); j++ { //j用来在i的每一个循环下，第j个字符串和第一个字符串作对比
			//默认循环内都是相等，等有特殊情况时返回循环到的当前i的字符串
			//特殊情况就是：当第一个字符串长度大于第二个的时候，循环停止即可
			//或者，第j个字符串的第i个字符和最开始存的prefix不一致了，则循环停止，直接返回当前i所在前缀即可
			if i >= len(strs[j]) || strs[j][i] != prefix { //注意一定要i大于等于strs[j]，
				return strs[0][0:i] // 如果strs =["ab", "a"]，则最后返回的是ab，越界

			}
		}
	}
	return strs[0] //如果以上特殊情况都没法发生，则直接返回第一个字符串即可
}

func main() {
	tests := [][]string{
		{"flower", "flow", "flight"},            // 示例：输出 "fl"
		{"dog", "racecar", "car"},               // 示例：输出 ""（无公共前缀）
		{"interview", "interrupt", "integrate"}, // 示例：输出 "int"
		{"a"},                                   // 单元素，输出 "a"
		{},                                      // 空切片，输出 ""
	}

	for _, tc := range tests {
		fmt.Printf("输入: %v -> 输出: %q\n", tc, longestCommonPrefix(tc))
	}
}
