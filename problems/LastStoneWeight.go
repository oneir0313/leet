package problems

import "container/heap"

// https://leetcode.com/problems/last-stone-weight/
// 此題重點 學習golang heap interface

func (p Problem) LastStoneWeight(stones []int) int {
    pq := maxHeap(stones)
    heap.Init(&pq)
    for pq.Len() > 1 {
		x, y := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
        if x != y {
            heap.Push(&pq, x-y)
        }
    }
    return heap.Pop(&pq).(int)
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[:n - 1]
    return x
}