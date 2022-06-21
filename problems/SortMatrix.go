package problems

import "sort"

func (p Problem) SortMatrix(matrix [][]int) []int {
	res := make([]int, len(matrix))
	sum := make([][]int, len(matrix))
	for i, x := range matrix {
		for xi, y := range x {
			if y == 0 || xi == len(x)-1 {
				sum[i] = []int{i, xi}
				break
			}
		}
	}

	sort.Slice(sum, func(i, j int) bool {
		return sum[i][1] < sum[j][1]
	})
	for i, v := range sum {
		res[i] = v[0]
	}
	return res
}
