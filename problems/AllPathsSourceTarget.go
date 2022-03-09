package problems

func (p Problem) AllPathsSourceTarget(graph [][]int) [][]int {
	var paths [][]int
	n := len(graph)

	var backtrack func(int, []int)
	backtrack = func(node int, curPath []int) {
		if node == n-1 {
			paths = append(paths, curPath)
			return
		}
		lenCurPath := len(curPath)
		for _, child := range graph[node] {
			tmp := make([]int, lenCurPath, lenCurPath+1)
			_ = copy(tmp, curPath) // create a new empty slice in which we copy the current path
			tmp = append(tmp, child)
			backtrack(child, tmp) // backtrack on a new instance of the current path - does not 'contaminate' the curPath variable.
		}
	}
	backtrack(0, []int{0})
	return paths
}
