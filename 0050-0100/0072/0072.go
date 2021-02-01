package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	// 初始化边界情况
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	// 动态规划
	for i := range word1 {
		for j := range word2 {
			tmp := min(dp[i][j+1], dp[i+1][j]) + 1
			if word1[i] == word2[j] {
				tmp = min(dp[i][j], tmp)
			} else {
				tmp = min(dp[i][j]+1, tmp)
			}
			dp[i+1][j+1] = tmp
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
