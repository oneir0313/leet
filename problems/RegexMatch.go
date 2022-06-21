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
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 2; i <= n; i += 2 {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc := s[i-1]
			pc := p[j-1]
			if sc == pc || pc == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pc == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if sc == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}
