package problems

import "sort"

// https://leetcode.com/problems/merge-intervals/

func (p Problem) MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		top := len(res) - 1
		if intervals[i][0] <= res[top][1] {
			max := 0
			if res[top][1] > intervals[i][1] {
				max = res[top][1]
			} else {
				max = intervals[i][1]
			}

			res[top] = []int{res[top][0], max}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
