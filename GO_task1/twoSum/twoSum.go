/* 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 创建一个map：键是数值，值是下标
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num // 想要的“补数”
		// 看看这个补数是否已经在map里
		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		// 否则，把当前数记录下来
		m[num] = i
	}

	// 如果没找到，可以返回空切片
	return []int{}
}

func main() {
	// 示例测试用例（按注释）
	tests := []struct {
		nums []int
		target int
	}{
		{[]int{2,7,11,15}, 9},    // 示例1 -> [0,1]
		{[]int{3,2,4}, 6},        // 额外测试 -> [1,2]
		{[]int{3,3}, 6},          // 重复元素 -> [0,1]
	}

	for _, tc := range tests {
		fmt.Printf("输入: nums=%v target=%d -> 输出: %v\n", tc.nums, tc.target, twoSum(append([]int(nil), tc.nums...), tc.target))
	}
}

