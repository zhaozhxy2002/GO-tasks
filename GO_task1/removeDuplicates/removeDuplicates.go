/*
	删除有序数组的重复项

给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，
使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
判题标准:
系统会用下列代码来测试你的题解:
int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案
int k = removeDuplicates(nums); // 调用
assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

如果所有断言都通过，那么您的题解将被 通过。

示例：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4,_,_,_,_,_]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
package main

import "fmt"

func removeDuplicates(nums []int) int {
	//快指针遍历数组，慢指针从nums[0]记录目前遍历到的最后一个不重复的数
	//由于不重复的元素只出现一次，所以slow只需要在发现下一个不重复数的时候，slow++即可
	//通过fast是否和slow相等来判断是否发现了下一个不重复的数字
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		//如果相等，说明不是重复元素，只需要fast往下走
		if nums[fast] != nums[slow] {
			slow++ //如果不相等，说明发现了下一个不重复元素，把这个数赋给slow++后的位置
			nums[slow] = nums[fast]
		}
	}
	return slow + 1 //最后返回的是nums的唯一元素的数量，也就是slow以及之前的所有数
}

func main() {
	// 测试用例1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("长度: %d, 数组: %v\n", k1, nums1[:k1]) // 长度: 2, 数组: [1 2]

	// 测试用例2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("长度: %d, 数组: %v\n", k2, nums2[:k2]) // 长度: 5, 数组: [0 1 2 3 4]

	// 测试用例3：空数组
	nums3 := []int{}
	k3 := removeDuplicates(nums3)
	fmt.Printf("长度: %d, 数组: %v\n", k3, nums3) // 长度: 0, 数组: []
}
