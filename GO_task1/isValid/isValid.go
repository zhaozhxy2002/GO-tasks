/* 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：输入：s = "()"     输出：true
示例 2：输入：s = "()[]{}" 输出：tru
示例 3：输入：s = "(]"     输出：false
示例 4：输入：s = "([])"   输出：true
示例 5：输入：s = "([)]"   输出：false
*/

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	StringMap := map[rune]rune{
		')': '(', //rune格式用单引号
		']': '[',
		'}': '{',
	} //定义map,通过右括号作为key，来查找相应的左括号

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char) //如果是左括号，直接压入栈顶
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false //如果此时右括号之前没有出现过左括号，则一定为false
			}
			top := stack[len(stack)-1] //弹出栈顶的左括号

			if StringMap[char] != top {
				return false //对比弹出的左括号，是否和右括号char的查找结果相同
			}
			stack = stack[0 : len(stack)-1] //如果不相同，直接false，但不管相同或不同，stack都需要弹出栈顶的左括号

		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([])",
		"([)]",
	}

	for _, tc := range tests {
		fmt.Printf("输入: %q -> 输出: %v\n", tc, isValid(tc))
	}
}
