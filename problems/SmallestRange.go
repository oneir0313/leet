package problems

import (
	"math"
	"sort"
)

type pair struct{
    num int
    group int
}

func (p Problem) SmallestRange(nums [][]int) []int {
    nGroup := len(nums)
    data := make([]pair, 0, 1024)
    for i := 0; i < nGroup; i++ {
        for j := 0; j < len(nums[i]); j++ {
            data = append(data, pair{nums[i][j], i})
        }
    }
    sort.Slice(data, func(i, j int) bool {return data[i].num < data[j].num})
    m := make(map[int]int)
    res := []int{}
    minRange := math.MaxInt64
    p1 := 0
    p2 := 0
    for p1 < len(data) && p2 < len(data) {
        m[data[p2].group]++
        p2++
        for len(m) == nGroup {
            if data[p2-1].num - data[p1].num < minRange {
                minRange = data[p2-1].num - data[p1].num
                res = []int{data[p1].num, data[p2-1].num}
            }
            m[data[p1].group]--
            if m[data[p1].group] == 0 {
                delete(m, data[p1].group)
            }
            p1++
        }
    }
    return res
}