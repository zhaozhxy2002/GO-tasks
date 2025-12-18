/*合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间.

示例 3：
输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
*/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//先把每个区间的第一个数进行排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	}) //把intervals中几个无序的区间一一对比，根据每个区间[0]第一个数的从小到大排列

	//排序后，建立result数组来储存结果数组
	result := make([][]int, 0, len(intervals))
	// 把第一个区间放入结果
	result = append(result, []int{intervals[0][0], intervals[0][1]})

	//用i遍历区间，对比前两个区间是否有重叠，
	//通过贪心算法，如果有重叠，合并到intervals[0]
	//如果不重叠，直接添加第二个区间到result中

	for i := 1; i < len(intervals); i++ {
		current := intervals[i] // 直接引用 result 最后一个区间
		lastIdx := len(result) - 1
		last := result[lastIdx] //retult的最后一个区间，用来和current作对比
		// 有重叠（当前起点 <= 上一个区间的终点）
		if current[0] <= last[1] {
			// 合并：更新为右端点的最大值
			if current[1] > last[1] {
				result[lastIdx][1] = current[1]
			}
			// 如果 current[1] <= last[1]，保持原来的 last[1]
		} else {
			// 无重叠，直接追加
			result = append(result, []int{current[0], current[1]})
		}
	}
	return result
}
func main() {
	intervals := [][]int{
		{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6},
	}
	fmt.Println("输入：", intervals)
	merged := merge(intervals)
	fmt.Println("合并后：", merged) // 期望 [[1,3],[4,7]]
}
