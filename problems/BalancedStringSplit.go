package problems

// https://leetcode.com/problems/split-a-string-in-balanced-strings/
func (p Problem) BalancedStringSplit(s string) int {
    var Ldepth, Rdepth, res int
	for i := range s {
		switch s[i] {
		case 'L':
            Ldepth++
			if Rdepth > 0 && Ldepth ==  Rdepth {
				Rdepth, Ldepth = 0, 0
                res++
            }
		case 'R':
            Rdepth++
			if Ldepth > 0 && Ldepth ==  Rdepth {
				Rdepth, Ldepth = 0, 0
                res++
            }
		}
	}
	return res
}