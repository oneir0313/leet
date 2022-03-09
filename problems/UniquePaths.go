package problems

func (p Problem) UniquePaths(m int, n int) int {
	left := 1
	up := make([]int, m)

	for i := range up {
		up[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			up[j] = up[j] + left
			left = up[j]
		}
		left = 1
	}

	return up[m-1]
}

func UniquePathsSolutionII(m int, n int) uint64 {
	if m >= n {
		return factorialTo(uint64(m+n-2), uint64(m-1)) / factorial(uint64(n-1))
	} else {
		return factorialTo(uint64(m+n-2), uint64(n-1)) / factorial(uint64(m-1))
	}
}

func factorial(n uint64) uint64 {
	var res uint64
	if n > 0 {
		res = uint64(n * factorial(n-1))
		return res
	}
	return 1
}

func factorialTo(n uint64, m uint64) uint64 {
	if n == m {
		return 1
	}
	var res uint64
	if n > 0 {
		res = uint64(n * factorialTo(n-1, m))
		return res
	}
	return 1
}
