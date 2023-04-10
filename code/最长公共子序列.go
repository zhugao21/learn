package code

func longestCommonSubsequence(text1 string, text2 string) int {
	x := len(text1) + 1
	y := len(text2) + 1

	var dp [][]int
	for j := 0; j < y; j++ {
		dp = append(dp, make([]int, x))
	}

	for j := 1; j < y; j++ {
		char1 := text2[j]
		for i := 1; i < x; i++ {
			char2 := text1[i]
			if char1 == char2 {
				dp[j][i] = dp[j-1][i-1] + 1
			} else {
				if dp[j-1][i] > dp[j][i-1] {
					dp[j][i] = dp[j-1][i]
				} else {
					dp[j][i] = dp[j][i-1]
				}
			}
		}
	}
	return dp[len(text2)][len(text1)]
}
