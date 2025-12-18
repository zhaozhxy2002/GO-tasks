/*
只出现一次的数字
给你一个非空整数数组 nums ，除了某个元素只出现 一次 外，其余每个元素均出现 两次 。请你找出并返回那个只出现了一次的元素。
示例 1 ：
输入：nums = [2,2,1]
输出：1
*/
package main

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = m[num] + 1 //相当于在map表中，以num为索引的对应值从0变为1，代表当前读取数字的出现次数为1
	} //如果没出现过，m[num]默认为0，加1
	//再次遍历 map，找出出现次数为1的那个数
	for k, v := range m { //理解为k == num, v ==m[num]
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	tests := [][]int{
		{2, 2, 1},       // 示例1 -> 1
		{4, 1, 2, 1, 2}, // 额外测试 -> 4
		{17},            // 单元素 -> 17
	}

	for _, tc := range tests {
		fmt.Printf("输入: %v -> 输出: %d\n", tc, singleNumber(append([]int(nil), tc...)))
	}
}
