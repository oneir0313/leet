package problems

import "sort"

// Input: mat =
// [[1,1,0,0,0],           // 2
//  [1,1,1,1,0],           // 4
//  [1,0,0,0,0],           // 1
//  [1,1,0,0,0],           // 2
//  [1,1,1,1,1]]           // 5

// 依照有1的數量從小到大把row的index印出來
// row[] of mat 只要遇到0後面則都是0
// Output: [2,0,3,1,4]


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
