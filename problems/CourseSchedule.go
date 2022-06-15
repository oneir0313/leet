package problems

func (p Problem) CourseSchedule(numCourses int, prerequisites [][]int) bool {
	// BFS uses the indegrees of each node. We will first try to find a node with 0 indegree. 
	// If we fail to do so, there must be a cycle in the graph and we return false.
	// Otherwise we set its indegree to be -1 to prevent from visiting it again and reduce the indegrees of its neighbors by 1. 
	// This process will be repeated for n (number of nodes) times.
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		degree[prerequisite[0]] += 1
	}
	
	bfs := make([]int, 0) // 將indegree為0的放入bfs 表示此課程能不需前置課程
	for course, d := range degree {
		if d == 0 {
			bfs = append(bfs, course)
		}
	}
	
	for i := 0; i < len(bfs); i ++{
		course := bfs[i]
		// 將indegree為0的node鄰居之計算-1 如此鄰居也變為indgree 0 表示此課程前置已完成能不需前置課程 將此放入bfs
		for _, j := range graph[course] {
			degree[j] -= 1
			if degree[j] == 0 {
				bfs = append(bfs, j)
			}
		}
	}
	
	// 判斷能完成之課程等於總數課程數量相同回傳結果
	return len(bfs) == numCourses
}

func  (p Problem) CourseScheduleByDFS(numCourses int, prerequisites [][]int) bool {
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
    }
    
    state := make([]int, numCourses)
    
    for i:=0; i<numCourses; i++ {
        if !DFS(graph, i, state) {
            return false
        }
    }
    
    return true
}

func DFS(g map[int][]int, cur int, state []int) bool {
    if state[cur] == 1 {
        return true
    }
    
    if state[cur] == 2 {
        return false
    }
    
    state[cur] = 2
    for _, neighbor := range g[cur] {
        if !DFS(g, neighbor, state) { //發現迴圈 因為state 2還正在訪問中
            return false
        }
    }
    state[cur] = 1
    
    return true
}