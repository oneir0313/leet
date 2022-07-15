package problems

import "sort"

func (p Problem) MaxWidthOfVerticalArea(points [][]int) int {
    res := points
    sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
    diff := 0
    for i := 1; i < len(res); i++ {
        calDiff := res[i][0] - res[i-1][0]
        if calDiff > diff {
            diff = calDiff
        }
    }
    return diff
}