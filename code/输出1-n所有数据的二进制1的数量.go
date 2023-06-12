package code

// 要求复杂度为n
// n=7，输出为：[1 1 2 1 2 2 3]
// 1：001	2：010	3：011
// 4：100	5：101	6：110	7：111

// 动态规划，最高有效位
// 当前值1的数量 = （该值-该值最大有效位）对应1的数量 + 1
// 如 dp[7] = dp[7-4]+1 = 3
// 等价于： 111 = （111-100）+1
func countBits(n int) []int {
	var dp = make([]int, n+1)
	var highBit int
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 { //判断是否是当前最高有效位
			highBit = i
		}
		dp[i] = dp[i-highBit] + 1
	}
	return dp[1:]
}
