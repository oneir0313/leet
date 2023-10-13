package problems

func (p Problem) RegexMatch(s string, pa string) bool {
	if len(pa) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (s[0] == pa[0] || pa[0] == '.')
	if len(pa) >= 2 && pa[1] == '*' {
		// isMatch(s, pa[2:]):               s = "aab" pa = "c*a*b"
		// firstMatch && isMatch(s[1:], pa): s = "aab" pa = "a*"
		return p.RegexMatch(s, pa[2:]) || (firstMatch && p.RegexMatch(s[1:], pa))
	}
	// firstMatch && isMatch(s[1:], pa[1:]): s = "aab" pa = "aab" or s = "b" pa = "b"
	return firstMatch && p.RegexMatch(s[1:], pa[1:])
}

// DP
func IsMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	if m > 0 && n > 0 {
		dp[1][1] = s[0] == p[0] || p[0] == '.'
	}
	for j := 2; j <= m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i <= n; i += 1 {
		for j := 2; j <= m; j += 1 {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}
	return dp[n][m]
}
