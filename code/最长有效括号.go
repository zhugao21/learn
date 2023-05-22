package code

//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
func longestValidParentheses(s string) int {
	// 思路：l,r判断"("与")"的数量。
	// 当两者相同时，情况为()()，说明此时已经是成对出现了，长度为2*l
	// 当l<r时，情况为())，需要重置l，r
	// 防止一直出现l<r，情况为(()，需要反向遍历一遍
	var left, right, max int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if max < 2*left {
				max = 2 * left
			}
		}
		if left < right {
			left = 0
			right = 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if max < 2*left {
				max = 2 * left
			}
		}

		if left > right {
			left = 0
			right = 0
		}
	}
	return max
}
